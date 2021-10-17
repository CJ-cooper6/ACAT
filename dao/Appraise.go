package dao

import (
	"ACAT/errmsg"
	"ACAT/model"
)

//评价相关

func SaveAppraise(student_num int,interview *model.Interview)int{	//评价
	err:=model.Db.Model(interview).Where("student_num=?",student_num).Updates(map[string]interface{}{
		"appraise": interview.Appraise,
		"evaluator": interview.Evaluator,
		"score": interview.Score}).Error
	if err!=nil{
		return errmsg.Error
	}
	return errmsg.Success

}

func Getappraise(student_num int)(*model.Interview,int){
	var view model.Interview
	err:=model.Db.Model(&view).Where("student_num=?",student_num).Find(&view).Error
	if err!=nil{
		return nil,errmsg.Error
	}
	return &view,errmsg.Success
}
