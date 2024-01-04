package app

import (
	"fmt"

	in "28/pkg/input"
	str "28/pkg/storage"
	stu "28/pkg/student"
)

func Run(studentStorage str.Storage) {
	for {
		inputStr, end := in.Get()
		if end {
			break
		}
		data, msg := in.ParseInput(inputStr)
		if msg != "" {
			fmt.Println(msg)
			continue
		}
		nameStr := data[0]
		ageNum, gradeNum, msg := stu.Validate(data[1], data[2])
		if msg != "" {
			fmt.Println(msg)
			continue
		}
		studentStorage.Put(nameStr, ageNum, gradeNum)
	}
}
