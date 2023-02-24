package sql

import (
	"gindemo01/struct"
)

func DelUserinfo(uid string) int {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user _struct.Userinfo
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&_struct.Userinfo{})
	db.First(&user)
	db.Debug().Where("Uid = ?", uid).Delete(&_struct.Userinfo{})
	return 1
}
func DelAdmininfo(uid string) int {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user _struct.Admininfo
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&_struct.Admininfo{})
	db.First(&user)
	db.Debug().Where("Uid = ?", uid).Delete(&_struct.Admininfo{})
	return 1
}
func DelUserRedis(U _struct.UserRedis) int {
	db := GetDB()
	db.SingularTable(true)
	var user _struct.UserRedis
	db.AutoMigrate(&_struct.UserRedis{})
	db.First(&user)
	db.Debug().Where("Uid = ? AND Data = ?", U.Uid, U.Data).Delete(&_struct.UserRedis{})
	return 1
}
func DelAdminRole(U _struct.AdminRole) int {
	db := GetDB()
	db.SingularTable(true)
	var user _struct.AdminRole
	db.AutoMigrate(&_struct.AdminRole{})
	db.First(&user)
	db.Debug().Where("role_id = ?", U.RoleId).Delete(&_struct.AdminRole{})
	return 1
}
func DelAffaris(uuid string) error {
	db := GetDB()
	db.SingularTable(true)
	var user _struct.Affairs
	db.AutoMigrate(&_struct.Affairs{})
	db.First(&user)
	err := db.Debug().Where("affairs_id = ?", uuid).Delete(&_struct.Affairs{}).Error
	return err
}
func DelUserIdentity(U _struct.UserIdentity) int {
	db := GetDB()
	db.SingularTable(true)
	var user _struct.UserIdentity
	db.AutoMigrate(&_struct.UserIdentity{})
	db.First(&user)
	db.Debug().Where("uid = ?", U.Uid).Delete(&_struct.UserIdentity{})
	return 1
}
