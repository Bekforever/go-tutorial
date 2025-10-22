package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"Google":  "https://google.com",
		"Yandex":  "https://yandex.ru",
		"Youtube": "https://youtube.com",
	}
	currentAction := 0

	showMenu()

	fmt.Print("Выберите действие: ")
	fmt.Scan(&currentAction)

	switch currentAction {
	case 1:
		showBookmarks(m)

	case 2:
		key, value := addBookmark()
		m[key] = value
		showBookmarks(m)
	case 3:
		selectedBookmark := deleteBookmark(m)
		delete(m, selectedBookmark)

		showBookmarks(m)
	default:
		break
	}
}

func showMenu() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Меню")
	fmt.Println("1: Посмотреть закладки")
	fmt.Println("2: Добавить закладку")
	fmt.Println("3: Удалить закладку")
	fmt.Println("4: Выход")
}

func showBookmarks(bookmarks map[string]string) {
	fmt.Println("")

	for key, value := range bookmarks {
		fmt.Println(key, " - ", value)
	}

	fmt.Println("")
}

func addBookmark() (string, string) {
	key, value := "", ""

	fmt.Print("Введите название для закладки: ")
	fmt.Scan(&key)
	fmt.Print("Введите значение: ")
	fmt.Scan(&value)

	return key, value
}

func deleteBookmark(bookmarks map[string]string) string {
	selectedBookmark := ""

	showBookmarks(bookmarks)

	fmt.Print("Введите ключ закладки которую хотите удалить: ")
	fmt.Scan(&selectedBookmark)

	return selectedBookmark
}
