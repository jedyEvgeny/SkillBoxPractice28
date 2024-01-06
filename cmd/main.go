package main

import (
	"28/app"
	rep "28/repository"
)

func main() {
	repository := rep.NewMemoryRepository()
	app := app.New(repository)
	app.Make()
}
