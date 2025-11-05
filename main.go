package main

import (
	"fmt"
	"math/rand"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputPassword() {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(maxLength int) {
	res := make([]rune, maxLength)

	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	acc.password = string(res)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func main() {
	login := promptData("Введите логин")
	// password := promptData("Введите пароль")
	url := promptData("Введите Url")

	myAccount := account{
		url:      url,
		login:    login,
	}
	myAccount.generatePassword(12)
	myAccount.outputPassword()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
