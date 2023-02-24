package _struct

import "time"

type GvaModel struct {
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	DeletedAt *string   `json:"-"` // 删除时间
}

type Userinfo struct {
	Uid         string `json:"uid,omitempty"`
	Email       string `json:"email,omitempty"`
	Username    string `json:"username,omitempty"`
	Passwd      string `json:"passwd,omitempty"`
	Permissions string `json:"permissions,omitempty"`
	Userstatus  int    `json:"userstatus,omitempty"`
	Signtime    string `json:"signtime,omitempty"`
	GvaModel
}
type Admininfo struct {
	Uid      string
	Email    string
	Username string
	Passwd   string
	Token    string
	State    int
	RoleId   int
	GvaModel
}
type UserRedis struct {
	Uid            string
	Data           string
	Type           string
	ExpirationTime string
	GvaModel
}
type UserLog struct {
	Uid  string
	Src  string
	data string
	GvaModel
}
type AdminRout struct {
	RoutId     int
	Name       string
	Title      string
	Icon       string
	Path       string
	Component  string
	Super      string
	Permission int
	GvaModel
}
type AdminRole struct {
	RoleId         string
	Name           string
	Description    string
	CreateBy       string
	UpdateBy       string
	PermissionJson string
	Orders         int
	GvaModel
}
type AdminPermission struct {
	PermissionId   int
	Permissions    string
	PermissionJson string
	CreateBy       string
	UpdateBy       string
	GvaModel
}
type Affairs struct {
	AffairsId   string
	Uid         string
	AffairsType string
	AffairsData string
	State       string
	DisposeTime string
	CreateBy    string
	UpdateBy    string
	GvaModel
}
type UserIdentity struct {
	Uid       string
	Name      string
	Sex       int
	Ethnic    string
	Polstatus string
	Origin    string
	Value     string
	Docunum   string
	Birthdate string
	Phonenum  string
	GvaModel
}

type UserImg struct {
	id   int
	Uid  string
	Name string
	File string
	GvaModel
}
