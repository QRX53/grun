package core

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var accounts [10]Account

type Account struct {
	IdNum       int64
	HolderName  string
	Balance     float32
	HoldsStocks bool
	AccountType string
}

func Welcome() {
	fmt.Println("Welcome to GRUN, an example CLI based banking application")
	fmt.Println("Would you like to open an acount today? Y/N: ")

	inputManager := bufio.NewReader(os.Stdin)
	openAccOrNot, _ := inputManager.ReadString('\n')
	openAccOrNot = strings.Trim(openAccOrNot, "\n")

	fmt.Println(openAccOrNot)

	if strings.Contains(openAccOrNot, "Y") {
		openAccount(*inputManager)
	} else {
		bank(*inputManager)
	}

}

func bank(inputManager bufio.Reader) {

	fmt.Println("Welcome to GRUN's banking system.")

	for {
		fmt.Println("What would you like to do?\n1. Open Account.\n2. Transfer Money\n3. Check Balance\n4. Buy Stocks")

		numChosen, _ := inputManager.ReadString('\n')
		numChosen = strings.Trim(numChosen, "\n")

		switch numChosen {
		case "1":
			openAccount(inputManager)
		case "2":
			transferMoney(inputManager)
		case "3":
			fmt.Println("You currently have $", accounts[0].Balance, "in your ", accounts[0].AccountType, "account.")
		case "4":
			buyStocks(inputManager)
		}
	}
}

func buyStock(stock string) {
	fmt.Println("You have chosen to buy stock in " + stock + ".")
	accounts[0].HoldsStocks = true
}

func buyStocks(inputManager bufio.Reader) {
	fmt.Println("Welcome to GRUN's stock buying menu. What would you like to buy? \nPlease note, your balance will stay the same, as your stocks count towards your balance.")
	fmt.Println("1. Apple\n2. Facebook\n3.Microsoft")

	numChosen, _ := inputManager.ReadString('\n')
	numChosen = strings.Trim(numChosen, "\n")

	switch numChosen {
	case "1":
		buyStock("Apple")
	case "2":
		buyStock("Facebook")
	case "3":
		buyStock("Microsoft")
	default:
		fmt.Println("You chose wrong. " + numChosen + " was not an option.")
	}

}

func transferMoney(inputManager bufio.Reader) {
	fmt.Println("How much money would you like to transfer? ")

	amt, _ := inputManager.ReadString('\n')
	amt = strings.Trim(amt, "\n")

	value, err := strconv.ParseFloat(amt, 32)

	if err != nil {
		fmt.Println("The amount you entered: \"" + amt + "\" is not valid.")
	}

	if float32(value) > accounts[0].Balance {
		fmt.Println("You cannot transfer $"+amt+" because you only have $", accounts[0].Balance, "available.")
		bank(inputManager)
	}

	fmt.Println("Who would you like to transfer the money to? (6 digit id number): ")
	who, _ := inputManager.ReadString('\n')
	who = strings.Trim(who, "\n")

	accounts[0].Balance -= float32(value)

	fmt.Println("Great job, you successfully transferred $"+amt+" to "+who+". You now have", accounts[0].Balance)
}

func openAccount(inputManager bufio.Reader) {
	fmt.Println("Welcome to the GRUN account creator. What type of account would you like to open today? (ie: Checkings, Savings): ")
	accType, _ := inputManager.ReadString('\n')
	accType = strings.Trim(accType, "\n")

	fmt.Println("Great. Now what is your first, and last name? ")
	flName, _ := inputManager.ReadString('\n')
	flName = strings.Trim(flName, "\n")

	fmt.Println("Great to meet you", flName+".")

	newAccount := Account{
		IdNum:       time.Now().UnixNano(),
		HolderName:  flName,
		Balance:     0.0,
		HoldsStocks: false,
		AccountType: accType,
	}

	fmt.Println("Would you like to deposit money into your new", accType, "today? Y/N: ")
	depoMoney, _ := inputManager.ReadString('\n')
	depoMoney = strings.Trim(depoMoney, "\n")

	if strings.Contains(depoMoney, "Y") {
		fmt.Println("How much money would you like to deposit today? ")
		amt, _ := inputManager.ReadString('\n')
		amt = strings.Trim(amt, "\n")

		value, err := strconv.ParseFloat(amt, 32)

		if err != nil {
			fmt.Println("The amount you entered: \"" + amt + "\" is not valid.")
		}

		newAccount.Balance += float32(value)
		fmt.Println("You have successfuly deposited $", value, "into your", accType, "account. You now have: $", newAccount.Balance)
	}

	accounts[0] = newAccount
	bank(inputManager)
}
