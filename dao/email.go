package dao

import (
	"ACAT/model"
)

func Checkemail(student_num int)string {
	var student model.Student
	model.Db.Where("student_num = ?", student_num).First(&student)
	return student.Student_email
}