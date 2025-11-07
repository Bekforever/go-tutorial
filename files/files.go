package files

import (
	"fmt"
	"os"
)

func ReadFile() {

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
