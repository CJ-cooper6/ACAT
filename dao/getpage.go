package dao

import (
	"ACAT/errmsg"
	"ACAT/model"
	"gorm.io/gorm"
)

func Getmainpage(pageSize int,pageNum int,interview_state int)([]model.Interview,int,int64){		//查询所有用户返回符合特定状态的
	var list []model.Interview
	var total int64
	offset:= (pageNum - 1) * pageSize
	if pageNum==-1&&pageSize==-1 {
		offset=-1
	}
	err:=model.Db.Limit(pageSize).Offset(offset).Where("interview_state=?",interview_state).Find(&list).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,errmsg.Error,0
	}
	return list,errmsg.Success,total
}


func Getappoint(role int,pageSize int,pageNum int,interview_state int)([]model.Interview,int,int64){	//查询特定方向的用户返回符合特定状态的
	var list []model.Interview
	var total int64
	offset:= (pageNum - 1) * pageSize
	if pageNum==-1&&pageSize==-1 {
		offset=-1
	}
	err:=model.Db.Limit(pageSize).Offset(offset).Where("choice=? AND interview_state=?",role,interview_state).Find(&list).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,errmsg.Error,0
	}
	return list,errmsg.Success,total
}

func Getadminpage(pageSize int,pageNum int)([]model.Admins,int,int64){		//查询所有管理员
	var admins []model.Admins
	var total int64
	offset:= (pageNum - 1) * pageSize
	if pageNum==-1&&pageSize==-1 {
		offset=-1
	}
	err:=model.Db.Model(&model.Admin{}).Limit(pageSize).Offset(offset).Find(&admins).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,errmsg.Error,0
	}
	return admins,errmsg.Success,total
}