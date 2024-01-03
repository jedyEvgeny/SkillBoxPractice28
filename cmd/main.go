package main

import (
	"fmt"

	sto "28/pkg/storage"
)

func main() {
	studentsStorage := sto.NewStorage()
	studentsStorage.Get()
	fmt.Println("--------------------")
	fmt.Println("Имена всех студентов:")
	fmt.Println("--------------------")
	studentsStorage.Print()
}
