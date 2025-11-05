package main

import (
	"fmt"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func main() {

	fmt.Println(generatePassword(12))

}

func generatePassword(maxLength int) string {
	res := make([]rune, maxLength)

	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(res)
}
