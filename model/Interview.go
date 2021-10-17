package model


type Interview struct { //学生面试信息表
	Student_Num     string ` gorm:"primaryKey" json:"student_num"`     //学号
	Student_Name    string `gorm:"type:varchar(20);not null" json:"student_name"`    //姓名
	Interview_State int    `gorm:"type:int(10);DEFAULT:0" json:"interview_state"` //面试状态
	Choice 			int `gorm:"type:int(10);DEFAULT:0" json:"choice"`//选择方向
	Sign              int    `gorm:"type:int(10);DEFAULT:0"  json:"sign"`              //签到状态
	Appraise        string `gorm:"type:varchar(200)" json:"appraise"`        //评价
	Evaluator       string `gorm:"type:varchar(10)" json:"evaluator"`       //评价人
	Score           int    `gorm:"type:int(10)" json:"score"`           //分数
	Student_email     string ` gorm:"type:varchar(20);not null" json:"student_email"`     //邮箱
	Student_telephone string ` gorm:"type:varchar(20);not null" json:"student_telephone"` //电话号码
}


