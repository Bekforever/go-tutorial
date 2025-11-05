package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
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

func newAccount(login, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	return &account{
		url:   urlString,
		login: login,
	}, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func main() {
	login := promptData("Введите логин")
	// password := promptData("Введите пароль")
	url := promptData("Введите Url")

	myAccount, err := newAccount(login, url)
	if err != nil {
		fmt.Println("Неверный формат URL")
		return
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
