package main

import (
	"fmt"
	"log"

	"github.com/usrname/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("kim")
	account.Deposit(10)
	fmt.Println(account.Balance())

	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account.Balance())
}
