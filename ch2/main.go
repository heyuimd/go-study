package main

import (
	"ch2/accounts"
	"ch2/dict"
	"fmt"
)

func main() {
	account := accounts.NewAccount("me")
	account.Deposit(100)
	err := account.Withdraw(200)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)

	// -----------------------------------------
	dictionary := dict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)

	err = dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dictionary)
}
