package sql

import (
	"gindemo01/struct"
	"time"
)

func ReviseUserinfo(U _struct.Userinfo) error {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user _struct.Userinfo
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&_struct.Userinfo{})
	db.First(&user)
	err := db.Debug().Model(&user).Where("uid = ?", U.Uid).Updates(map[string]interface{}{"Username": U.Username, "Permissions": U.Permissions, "Passwd": U.Passwd, "Userstatus": U.Userstatus}).Error
	return err
}
func ReviseAdmininfo(U _struct.Admininfo) int {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user _struct.Admininfo
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&_struct.Admininfo{})
	db.First(&user)
	db.Debug().Model(&user).Where("uid = ?", U.Uid).Updates(map[string]interface{}{"Username": U.Username, "State": U.State, "Passwd": U.Passwd, "Token": U.Token})
	return 1
}
func ReviseAdminrole(U _struct.AdminRole) int {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user _struct.AdminRole
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&_struct.AdminRole{})
	db.First(&user)
	db.Debug().Model(&user).Where("role_id = ?", U.RoleId).Updates(map[string]interface{}{
		"name":            U.Name,
		"description":     U.Description,
		"update_time":     time.Now().String()[0:19],
		"update_by":       U.UpdateBy,
		"permission_json": U.PermissionJson,
		"orders":          U.Orders,
	})
	return 1
}
func ReviseAdminpermission(U _struct.AdminPermission) int {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user _struct.AdminPermission
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&_struct.AdminPermission{})
	db.First(&user)
	db.Debug().Model(&user).Where("permission_id = ?", U.PermissionId).Updates(map[string]interface{}{})
	return 1
}
func ReviseAffairs(U _struct.Affairs) error {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user _struct.AdminPermission
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&_struct.AdminPermission{})
	db.First(&user)
	err := db.Debug().Model(&user).Where("affairs_id = ?", U.AffairsId).Updates(map[string]interface{}{
		"affairs_data": U.AffairsData,
		"state":        U.State,
		"update_by":    U.UpdateBy,
		"dispose_time": U.DisposeTime,
	}).Error
	return err
}
func ReviseUserIdentity(U _struct.UserIdentity) error {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user _struct.UserIdentity
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&_struct.UserIdentity{})
	err := db.Debug().Model(&user).Where("uid = ?", U.Uid).Updates(map[string]interface{}{
		"name":      U.Name,
		"sex":       U.Sex,
		"ethnic":    U.Ethnic,
		"polstatus": U.Polstatus,
		"origin":    U.Origin,
		"value":     U.Value,
		"docunum":   U.Docunum,
		"birthdate": U.Birthdate,
		"phonenum":  U.Phonenum,
	}).Error
	return err
}
func ReviseUserPasswd(email string, passwd string) error {
	var u _struct.Userinfo
	err := GetDB().Model(&u).Where("email = ?", email).Updates(map[string]interface{}{
		"passwd": passwd,
	}).Error
	return err
}
