package student

import (
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func New(name string, age, grade int) *Student {
	return &Student{Name: name, Age: age, Grade: grade}
}

func Validate(ageStr, gradeStr string) (int, int, string) {
	var msg string
	age, ageErr := strconv.Atoi(ageStr)
	grade, gradeErr := strconv.Atoi(gradeStr)
	if ageErr != nil && gradeErr != nil {
		msg = "Ошибка в возрасте и классе"
	}
	if ageErr != nil {
		msg = "Ошибка в возрасте"
	}
	if gradeErr != nil {
		msg = "Ошибка в классе"
	}
	return age, grade, msg
}
