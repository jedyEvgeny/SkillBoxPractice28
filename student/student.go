package student

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade int
}

func (s *Student) String() string {
	strOut := fmt.Sprintf("Имя студента: %v, возраст студента: %v, курс студента: %v.", s.Name, s.Age, s.Grade)
	return strOut
}
