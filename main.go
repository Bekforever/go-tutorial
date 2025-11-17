package main

import (
	"demo/app/account"
	"demo/app/files"
	"fmt"
)

func main() {
	createAccount()
}

func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите Url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или логина")
		return
	}
	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать файл в json")
	}
	files.WriteFile(file, "data.json")
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
