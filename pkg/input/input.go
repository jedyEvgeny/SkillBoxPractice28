package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Get() (string, bool) {
	fmt.Print("Введите данные студента (имя возраст класс) или Enter для завершения: ")
	scanner := bufio.NewScanner(os.Stdin)
	var endProgram bool
	if err := scanner.Err(); !scanner.Scan() || err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения: %v\n", err)
	}
	input := scanner.Text()
	if input == "" {
		endProgram = true
	}
	return input, endProgram
}

func ParseInput(input string) (data []string, msg string) {
	data = strings.Split(input, " ")
	if len(data) != 3 {
		msg = "Неверный формат ввода"
	}
	return data, msg
}
