package dao

import (
	"ACAT/errmsg"
	"ACAT/model"
	"gorm.io/gorm"
)

func Getmainpage(pageSize int,pageNum int)([]model.Interview,int,int64){		//查询所有用户面试状态
	var list []model.Interview
	var total int64
	offset:= (pageNum - 1) * pageSize
	if pageNum==-1&&pageSize==-1 {
		offset=-1
	}
	err:=model.Db.Limit(pageSize).Offset(offset).Find(&list).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,errmsg.Error,0
	}
	return list,errmsg.Success,total
}


func Getappoint(role int,pageSize int,pageNum int)([]model.Interview,int,int64){	//查询特定用户面试状态
	var list []model.Interview
	var total int64
	offset:= (pageNum - 1) * pageSize
	if pageNum==-1&&pageSize==-1 {
		offset=-1
	}
	err:=model.Db.Limit(pageSize).Offset(offset).Where("choice=?",role).Find(&list).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,errmsg.Error,0
	}
	return list,errmsg.Success,total
}

func Getadminpage(pageSize int,pageNum int)([]model.Admin,int,int64){		//查询所有管理员
	var admins []model.Admin
	var total int64
	offset:= (pageNum - 1) * pageSize
	if pageNum==-1&&pageSize==-1 {
		offset=-1
	}
	err:=model.Db.Limit(pageSize).Offset(offset).Find(&admins).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,errmsg.Error,0
	}
	return admins,errmsg.Success,total
}