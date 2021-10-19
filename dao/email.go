package dao

import (
	"ACAT/model"
)

func Checknameemail(student_num int)(string,string,int) {
	var s model.Interview
	model.Db.Where("student_num = ?", student_num).First(&s)
	return s.Student_email,s.Student_Name,s.Choice
}