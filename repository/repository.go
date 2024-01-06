package repository

import (
	std "28/student"
	"errors"
)

type MemoryRepository struct {
	studentsKeyName map[string]*std.Student
}

func NewMemoryRepository() *MemoryRepository {
	NewRepository := &MemoryRepository{
		studentsKeyName: make(map[string]*std.Student),
	}
	return NewRepository
}

func (m *MemoryRepository) Get() ([]*std.Student, error) {
	var students []*std.Student
	for _, value := range m.studentsKeyName {
		students = append(students, value)
	}
	return students, nil
}

func (m *MemoryRepository) Put(student *std.Student) (err error) { //Получаем указатель на структуру Студент
	_, ok := m.studentsKeyName[student.Name]
	if !ok {
		m.studentsKeyName[student.Name] = student
	}
	if ok {
		err = errors.New("студент с таким именем зачислен ранее")
	}
	return err
}
