package sql

import (
	"gindemo01/struct"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func UserinfoAdd(U _struct.Userinfo) int {
	u := UserinfoFind("email = ?", U.Email)
	if len(u) > 0 {
		return 3
	} else {
		db := GetDB()
		db.SingularTable(true)
		//fmt.Println(db.NewRecord(&U))
		db.Create(&U) //将上边定义的用户数据写入到数据库user表中
		//fmt.Println(db.NewRecord(&U))
	}
	return 1
}
func AdmininfoAdd(U _struct.Admininfo) int {
	u := AdmininfoFind("email = ?", U.Email)
	if len(u) > 0 {
		return 3
	} else {
		db := GetDB()
		db.SingularTable(true)
		//fmt.Println(db.NewRecord(&U))
		db.Create(&U) //将上边定义的用户数据写入到数据库user表中
		//fmt.Println(db.NewRecord(&U))
	}
	return 1
}
func UserImgAdd(U _struct.UserImg) int {
	u := UserinfoFind("uid = ?", U.Uid)
	if len(u) == 0 {
		return 3
	} else {
		db := GetDB()
		db.SingularTable(true)
		//fmt.Println(db.NewRecord(&U))
		db.Create(&U) //将上边定义的用户数据写入到数据库user表中
		//fmt.Println(db.NewRecord(&U))
	}
	return 1
}
func UserRedisAdd(U _struct.UserRedis) int {
	db := GetDB()
	db.SingularTable(true)
	//fmt.Println(db.NewRecord(&U))
	db.Create(&U) //将上边定义的用户数据写入到数据库user表中
	//fmt.Println(db.NewRecord(&U))
	return 1
}
func AdminRoleAdd(U _struct.AdminRole) error {
	db := GetDB()
	db.SingularTable(true)
	//fmt.Println(db.NewRecord(&U))
	err := db.Create(&U).Error //将上边定义的用户数据写入到数据库user表中
	//fmt.Println(db.NewRecord(&U))
	return err
}
func AdminPermissionAdd(U _struct.AdminPermission) int {
	db := GetDB()
	db.SingularTable(true)
	db.Create(&U) //将上边定义的用户数据写入到数据库user表中
	return 1
}
func AffairsAdd(U _struct.Affairs) error {
	db := GetDB()
	db.SingularTable(true)
	err := db.Create(&U).Error //将上边定义的用户数据写入到数据库user表中
	return err
}
func UserIdentityAdd(U _struct.UserIdentity) error {
	db := GetDB()
	db.SingularTable(true)
	err := db.Create(&U).Error //将上边定义的用户数据写入到数据库user表中
	return err
}
