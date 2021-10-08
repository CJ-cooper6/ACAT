package dao

import (
	"ACAT/errmsg"
	"ACAT/model"
)

//评价相关

func SaveAppraise(student_num int,interview *model.Interview)int{
	err:=model.Db.Model(interview).Where("student_num=?",student_num).Updates(map[string]interface{}{
		"appraise": interview.Appraise,
		"evaluator": interview.Evaluator,
		"score": interview.Score}).Error
	if err!=nil{
		return errmsg.Error
	}
	return errmsg.Success

}
