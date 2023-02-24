package sql

import (
	"gindemo01/struct"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func UserinfoFind(query string, args string) []_struct.Userinfo {

	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&_struct.Userinfo{})
	var user []_struct.Userinfo
	db.Debug().Where(query, args).Find(&user)
	return user
}
func UserinfoFirst(field string, content string) (_struct.Userinfo, error) {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&_struct.Userinfo{})
	var user _struct.Userinfo
	err := db.Debug().Where(field+" = ?", content).First(&user).Error
	return user, err
}
func AdmininfoFind(query string, args string) []_struct.Admininfo {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&_struct.Admininfo{})
	var user []_struct.Admininfo
	db.Debug().Where(query, args).Find(&user)
	return user
}
func AdmininfoFirst(field string, content string) _struct.Admininfo {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&_struct.Admininfo{})
	var user []_struct.Admininfo
	db.Debug().Where(field+" = ?", content).First(&user)
	return user[0]
}
func UserRedisFind(field string, content string) []_struct.UserRedis {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&_struct.UserRedis{})
	var user []_struct.UserRedis
	db.Debug().Where(field+" = ?", content).Find(&user)
	return user
}
func AdminRoutFind(query string, args string) []_struct.AdminRout {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&_struct.AdminRout{})
	var user []_struct.AdminRout
	db.Debug().Where(query, args).Find(&user)
	return user
}
func AdminRoleFind(query string, args string) ([]_struct.AdminRole, error) {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&_struct.AdminRole{})
	var user []_struct.AdminRole
	err := db.Debug().Where(query, args).Find(&user).Error
	return user, err
}
func AdminPermissionFind(query string, args string) []_struct.AdminPermission {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&_struct.AdminPermission{})
	var user []_struct.AdminPermission
	db.Debug().Where(query, args).Find(&user)
	return user
}
func AffairsFind(query string, args string) ([]_struct.Affairs, error) {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&_struct.Affairs{})
	var user []_struct.Affairs
	err := db.Debug().Where(query, args).Find(&user).Error
	return user, err
}

func UserIdentityFind(query string, args string) []_struct.UserIdentity {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&_struct.UserIdentity{})
	var user []_struct.UserIdentity
	db.Debug().Where(query, args).Find(&user)
	return user
}
func UserIdentityFirst(field string, content string) _struct.UserIdentity {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&_struct.UserIdentity{})
	var user []_struct.UserIdentity
	db.Debug().Where(field+" = ?", content).First(&user)
	return user[0]
}
