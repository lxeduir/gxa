package backstage

import (
	"encoding/json"
	"gindemo01/public/sql"
	"gindemo01/struct"
	"github.com/gin-gonic/gin"
)

type permission struct {
	Type string `json:"type"`
	Id   string `json:"type_id"`
	P    string `json:"permission"`
}
type rAdminPermission struct {
	PermissionId   int
	Permissions    string
	PermissionJson listAdminPermission
	CreateBy       string
	CreateTime     string
	UpdateBy       string
	UpdateTime     string
}
type listAdminPermission struct {
	List []permission `json:"list"`
}
type wAdminPermission struct {
	PermissionId   int                 `json:"PermissionId"`
	Permissions    string              `json:"Permissions"`
	PermissionJson listAdminPermission `json:"PermissionJson"`
}

func GetPermission(context *gin.Context) {
	P := sql.AdminPermissionFind("permission_id LIKE ?", "%")
	var R []rAdminPermission
	for _, v := range P {
		b := []byte(v.PermissionJson)
		var m listAdminPermission
		err := json.Unmarshal(b, &m)
		if err != nil {
			context.JSON(201, gin.H{
				"msg": err.Error(),
			})
			return
		}
		var r rAdminPermission
		r.PermissionId = v.PermissionId
		r.CreateBy = v.CreateBy
		r.CreateTime = v.CreatedAt.String()
		r.UpdateBy = v.UpdateBy
		r.UpdateTime = v.UpdatedAt.String()
		r.PermissionJson = m
		r.Permissions = v.Permissions
		R = append(R, r)
	}
	context.JSON(200, gin.H{
		"list":       R,
		"type_id":    "0表示所有id",
		"type":       "0表示所有类型",
		"permission": "0表示不可见，1表示可见，2表示可编辑，3表示可删除",
	})
}
func DelPermission(context *gin.Context) {

}
func PutPermission(context *gin.Context) {

}
func PostPermission(context *gin.Context) {
	var u wAdminPermission
	if err := context.ShouldBind(&u); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
	} else {
		var s _struct.AdminPermission
		s.PermissionId = u.PermissionId
		s.Permissions = u.Permissions
		s.CreateBy = "等待实现"
		s.UpdateBy = "等待实现"
		var js string = `{"list":[`
		for i, v := range u.PermissionJson.List {
			js = js + `{"type":"` + v.Type + `","type_id":"` + v.Id + `","permission":"` + v.P + `"}`
			if i != len(u.PermissionJson.List)-1 {
				js = js + ","
			}
		}
		js = js + `]}`
		s.PermissionJson = js
		msg := sql.AdminPermissionAdd(s)
		context.JSON(200, gin.H{
			"msg": msg,
		})
	}
}
