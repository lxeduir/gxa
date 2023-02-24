package upload

import (
	"fmt"
	"gindemo01/public"
	"gindemo01/public/sql"
	"gindemo01/struct"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

func Img(c *gin.Context) {
	uid, err1 := c.GetQuery("uid")
	name, err2 := c.GetQuery("name")
	file, err := c.FormFile("file")
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".png" && fileExt != ".img" && fileExt != ".jpeg" && fileExt != ".jpg" && fileExt != ".7z" {
		c.JSON(http.StatusOK, gin.H{"uploading": "done", "message": "上传文件格式错误"})
		return
	}
	U := sql.UserinfoFind("uid = ?", uid)
	if len(U) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "用户不存在",
		})
		return
	}
	if !err1 || !err2 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err1,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 上传文件到指定的目录
	fileName := uid + "_" + name + "_" + strconv.FormatInt(time.Now().Unix(), 10) + path.Ext(file.Filename)
	err3 := c.SaveUploadedFile(file, "./Assets/"+fileName)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err3.Error(),
		})
		return
	}
	var u _struct.UserImg
	u.Uid = uid
	u.File = "./Assets/" + fileName
	u.Name = name
	msg := sql.UserImgAdd(u)
	upload, err := public.Upload(u.File, fileName, map[string]string{
		"x:name": name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' upload success.", file.Filename),
		"msg":     msg,
		"qiniu":   upload,
	})
}

func GetImg(c *gin.Context) {
	uuids := c.Param("uuids")
	fmt.Println(uuids)
}
func PostImg(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".png" && fileExt != ".img" && fileExt != ".jpeg" && fileExt != ".jpg" {
		c.JSON(http.StatusOK, gin.H{"uploading": "done", "message": "上传文件格式错误"})
		return
	}
	// 上传文件到指定的目录使用uuid作为文件名字
	fileName := strings.Replace(uuid.NewString(), "-", "", -1) + path.Ext(file.Filename)
	err = c.SaveUploadedFile(file, "./Assets/"+fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	upload, err := public.Upload("./Assets/"+fileName, fileName, map[string]string{
		"x:name": fileName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' upload success.", file.Filename),
		"src":     "https://img.lxlxlxlx.xyz/gxa/" + upload.Key,
	})
	return

}
