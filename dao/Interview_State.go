package dao

import (
	"ACAT/errmsg"
	"ACAT/model"
)

//面试状态相关
func ChangeInterview_State(student_num int,Interview_State int)int{
	var interview model.Interview
	var student model.Student
	err:=model.Db.Model(&interview).Where("student_num=?",student_num).Updates(map[string]interface{}{	//修改interview表中的面试状态
		"interview_state":Interview_State}).Error

	err=model.Db.Model(&student).Where("student_num=?",student_num).Updates(map[string]interface{}{		//修改student表中的面试状态
		"interview_state":Interview_State}).Error

	if err!=nil{
		return errmsg.Error
	}
	return errmsg.Success

}