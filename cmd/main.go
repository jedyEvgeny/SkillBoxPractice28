package main

import (
	"28/pkg/app"
	str "28/pkg/storage"
)

func main() {
	studentStorage := str.New()
	app.Run(studentStorage)
	studentStorage.PrintAllStudents()
}
