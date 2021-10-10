package model

import "gorm.io/gorm"


type Interview struct { //学生面试信息表
	gorm.Model
	Student_Num     string `gorm:"type:varchar(20);not null" json:"student_num"`     //学号
	Student_Name    string `gorm:"type:varchar(20);not null" json:"student_name"`    //姓名
	Interview_State int    `gorm:"type:int(10);DEFAULT:0" json:"interview_state"` //面试状态
	Choice 			int `gorm:"type:int(10);DEFAULT:0" json:"choice"`//选择方向
	Appraise        string `gorm:"type:varchar(200)" json:"appraise"`        //评价
	Evaluator       string `gorm:"type:varchar(10)" json:"evaluator"`       //评价人
	Score           int    `gorm:"type:int(10)" json:"score"`           //分数
}


