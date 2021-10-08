package model

import (
	"gorm.io/gorm"
)

type Admin struct { //管理员信息表
	gorm.Model
	Admin_Num      string `gorm:"type:varchar(20);not null" json:"admin_num"`      //	学号
	Admin_Password string `gorm:"type:varchar(20);not null" json:"admin_password"` //密码
	Admin_Name     string `gorm:"type:varchar(20);not null" json:"admin_name"`     //姓名
	Role 		   int `gorm:"type:int(10);DEFAULT:0" json:"role"`				   //角色
}

