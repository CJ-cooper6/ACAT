package model

import (
	"gorm.io/gorm"
)
type Student struct { //学生信息表
	gorm.Model
	Student_Num       string ` gorm:"type:varchar(20);not null" json:"student_num"`                                  //学号
	Student_Class     string ` gorm:"type:varchar(20);not null" json:"student_class"`     //	班级
	Student_Name      string ` gorm:"type:varchar(15);not null" json:"student_name"`      //姓名
	Interview_State   int 	 ` gorm:"type:int(10);DEFAULT:0" json:"interview_state"`   //面试状态
	Choice            int    ` gorm:"type:int(10);DEFAULT:0" json:"choice"`            //选择方向
	Sign              int    `gorm:"type:int(10);DEFAULT:0"  json:"sign"`              //签到状态
	Student_email     string ` gorm:"type:varchar(20);not null" json:"student_email"`     //邮箱
	Student_telephone string ` gorm:"type:varchar(20);not null" json:"student_telephone"` //电话号码

}