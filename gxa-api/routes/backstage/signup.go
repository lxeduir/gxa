package backstage

import (
	"gindemo01/public"
	"gindemo01/public/sql"
	"gindemo01/struct"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type signUpAdminR struct {
	msg   int    //操作结果
	state int    //账号状态
	uid   string //账号id
	token string //识别码
}

func signUpAdmin(c *gin.Context) {
	var u _struct.Admininfo
	var U signUpAdminR
	U.msg = 1
	U.state = 0
	U.token = "err"
	U.uid = "err"
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		u.Uid = strconv.FormatInt((time.Now().Unix()-660000000)*100+int64(r.Intn(128)), 10)
		u.Passwd = public.MD5(u.Passwd + u.Uid)
		u.State = 111
		u.RoleId = 111
		u.Token = SetTokenAdmininfo(u, time.Hour)

	}
	if public.VerifyEmailFormat(u.Email) {
		if sql.AdmininfoAdd(u) == 1 {
			U.state = 111
			U.token = u.Token
			U.uid = u.Uid
		}

	} else {
		U.msg = 2
	}

	c.JSON(200, gin.H{
		"msg":   U.msg,
		"state": U.state,
		"uid":   U.uid,
		"token": U.token,
	})
}
