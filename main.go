package main

import (
	"demo/app/account"
	"demo/app/files"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var users = make(map[string]*account.Account)

func main() {
	for {
		variant := showMenu()

		switch variant {
		case 1:
			createAccount()
			showMap()
		case 2:
			findAccount()
			showMap()
		case 3:
			deleteAccount()
			showMap()
		case 4:
			fmt.Println("До встречи!")
			data, err := showMap()
			if err != nil {
				fmt.Println("error")
				os.Exit(1)
			}
			files.WriteFile(data, "data.json")
			os.Exit(0)
		}
	}
}

func showMenu() int {
	var selected int
	fmt.Println("Меню:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Println("")
	fmt.Print("Выберите действие: ")
	fmt.Scan(&selected)
	return selected
}

func showMap() ([]byte, error) {
	data, err := json.Marshal(users)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func createAccount() error {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите Url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или логина")
		return errors.New("unknown error")
	}
	users[myAccount.Login] = myAccount
	return nil
}

func findAccount() {
	var searchingLogin string
	fmt.Print("Введите логин который вы хотите найти: ")
	fmt.Scan(&searchingLogin)
	if users[searchingLogin] != nil {
		fmt.Println("Аккаунт найден:")
		fmt.Println(users[searchingLogin])
		return
	}
	fmt.Println("Аккаунт не найден")
}

func deleteAccount() {
	var accountToDelete string
	fmt.Print("Введите логин который вы хотите удалить: ")
	fmt.Scan(&accountToDelete)
	if users[accountToDelete] != nil {
		delete(users, accountToDelete)
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
