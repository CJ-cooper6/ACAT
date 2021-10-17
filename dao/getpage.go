package dao

import (
	"ACAT/errmsg"
	"ACAT/model"
	"gorm.io/gorm"
)

func GetInterview(role int,pageSize int,pageNum int,interview_state int)([]model.Interview,int,int64){	//根据传入角色返回符合特定状态的
	var list []model.Interview
	var total int64

	offset:= (pageNum - 1) * pageSize
	if pageNum==-1&&pageSize==-1 {
		offset=-1
	}
	switch role{
	case 1:
		model.Db.Limit(pageSize).Offset(offset).Where("interview_state=? and sign =1",interview_state).Find(&list).Offset(-1).Limit(-1).Count(&total)	//查询所有用户返回符合特定状态的
	default:
		model.Db.Limit(pageSize).Offset(offset).Where("choice=? and interview_state=? and sign =1",role,interview_state).Find(&list).Offset(-1).Limit(-1).Count(&total)
	}

	return list,errmsg.Success,total
}



func Getadminpage(pageSize int,pageNum int,admin_Num string)([]model.Admins,int,int64){		//查询所有管理员
	var admins []model.Admins
	var total int64
	offset:= (pageNum - 1) * pageSize
	if pageNum==-1&&pageSize==-1 {
		offset=-1
	}
	err:=model.Db.Model(&model.Admin{}).Limit(pageSize).Offset(offset).Not("admin_num=?",admin_Num).Find(&admins).Offset(-1).Limit(-1).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,errmsg.Error,0
	}
	return admins,errmsg.Success,total
}