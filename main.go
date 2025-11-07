package main

import (
	"demo/app/account"
	"demo/app/files"
	"fmt"
)

func main() {
	files.WriteFile("Hello world", "test")
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите Url")

	// myAccount, err := newAccount(login, password, url)
	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или логина")
		return
	}
	myAccount.OutputPassword()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
