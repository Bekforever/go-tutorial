package account

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.login)
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(maxLength int) {
	res := make([]rune, maxLength)

	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	acc.password = string(res)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			url:      urlString,
			password: password,
			login:    login,
		},
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}
