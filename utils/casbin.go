package utils

import gormadapter "github.com/casbin/gorm-adapter/v3"

var Carbines = []gormadapter.CasbinRule{		//要存在数据库的策略
	{Ptype: "p",V0: "1",V1: "/api/admin/superadmin/getadminpage",V2: "change"},
	{Ptype: "p",V0: "1",V1: "/api/admin/superadmin/changepower",V2: "change"},
	{Ptype:"g",V0:"admin123",V1:"1"},
}

