package storage

import (
	"fmt"

	std "28/pkg/student"
)

type Storage map[string]*std.Student

func New() Storage {
	return make(Storage)
}

func (s Storage) Put(name string, age int, grade int) {
	stud := std.New(name, age, grade)
	s[name] = stud
}

func (s Storage) PrintAllStudents() {
	fmt.Println("--------------------")
	fmt.Println("Список студентов:")
	fmt.Println("--------------------")
	for _, value := range s {
		fmt.Println(value.Name, value.Age, value.Grade)
	}
}
