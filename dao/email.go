package dao

import (
	"ACAT/model"
)

func Checkemail(student_num int)string {
	var s model.Interview
	model.Db.Where("student_num = ?", student_num).First(&s)
	return s.Student_email
}