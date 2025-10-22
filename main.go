package main

import "fmt"

func main() {
	m := map[string]string{
		"Bekforever": "https://bekforever.com",
	}

	fmt.Println(m)
	fmt.Println(m["Bekforever"])
	m["Google"] = "https://google.com"

	fmt.Println(m)
	delete(m, "Google")
	fmt.Println(m)
}
