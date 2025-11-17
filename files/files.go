package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(file))
}

func WriteFile(content, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}

	_, error := file.WriteString(content)
	defer file.Close()

	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println("Запись успешна")

}
