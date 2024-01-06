package app

import (
	"28/student"
	std "28/student"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type DataStore interface {
	Put(*std.Student) error       //Передаём указатель на экземпляр структуры Студент
	Get() ([]*std.Student, error) //Получаем массив из указателей на экземпляры структур Студент
}

type App struct {
	cache DataStore
}

var russianLettersRegexp = regexp.MustCompile("^[А-Яа-я]+$")

func New(c DataStore) *App { //На вход подаётся неизвестный объект с методами Put и Get
	newApp := &App{ //Который преобразуется в структуру с методами Put и Get
		cache: c,
	}
	return newApp //На выходе указатель на созданную структуру
}

func (a *App) Make() { //аналог main в немейновом пакете - основная функция
	a.inputingStudent()
	a.printStudents()
}

func (a *App) inputingStudent() {

	for {
		fmt.Print("Введите данные студента (имя возраст класс) или нажмите клавишу Enter для завершения: ")
		scanner := bufio.NewScanner(os.Stdin)
		if err := scanner.Err(); !scanner.Scan() || err != nil {
			fmt.Println("Ошибка чтения: ", err)
		}
		str := scanner.Text()
		if str == "" {
			break
		}
		student, msgErr := validate(str)
		if msgErr != "" {
			fmt.Println(msgErr)
			continue
		}
		err := a.cache.Put(student)
		if err != nil {
			fmt.Println("Не удалось зачислить студента на курс: ", err)
			continue
		}
	}
}

func (a *App) printStudents() {
	students, err := a.cache.Get()
	if err != nil {
		fmt.Println("Ошибка при получении данных: ", err)
		return
	}
	fmt.Println("----------------")
	fmt.Println("Список студентов:")
	for _, v := range students {
		fmt.Printf("\t- %v\n", v)
	}
}

func validate(line string) (*std.Student, string) {
	var msg string
	arr := strings.Split(line, " ")
	if len(arr) != 3 {
		msg = "Неверный формат строки"
		return nil, msg
	}
	name, ok := createName(arr[0])
	if !ok {
		msg = "Имя должно состоять из букв кириллицы"
		return nil, msg
	}
	age, errAge := strconv.Atoi(arr[1])
	grade, errGrade := strconv.Atoi(arr[2])
	msg = createMessageErrAgeAndGrade(errAge, errGrade)
	student := createStudent(name, age, grade)
	return student, msg
}

func createName(str string) (string, bool) {
	var ok bool
	if russianLettersRegexp.MatchString(str) {
		ok = true
	}
	return str, ok
}

func createMessageErrAgeAndGrade(errAge, errGrade error) (msg string) {
	if errAge != nil {
		msg = "Неверный формат возраста"
	}
	if errGrade != nil {
		msg = "Неверный формат класса"
	}
	if errAge != nil && errGrade != nil {
		msg = "Неверный формат возраста и класса"
	}
	return msg
}

func createStudent(name string, age, grade int) *student.Student {
	student := &std.Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	return student
}
