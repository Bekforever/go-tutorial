package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.Login)
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal((acc))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (acc *Account) generatePassword(maxLength int) {
	res := make([]rune, maxLength)

	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       urlString,
		Password:  password,
		Login:     login,
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}
