package backstage

import (
	"gindemo01/public/sql"
	"gindemo01/struct"
	"github.com/gin-gonic/gin"
)

type metas struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}
type super struct {
	RoleId    int     `json:"role_id"`
	Path      string  `json:"path"`
	Component string  `json:"component"`
	Name      string  `json:"name"`
	Meta      metas   `json:"meta"`
	Children  []super `json:"children,omitempty"`
}

func path(c *gin.Context) {
	cla, ok1 := c.Get("cla")
	Cla := cla.(Claimadmins)
	if ok1 != true {
		c.JSON(200, gin.H{"msg": "uid不能为空"})
		return
	} else {
		paths := sql.AdminRoutFind("rout_id LIKE ?", "%")
		var rpath []_struct.AdminRout
		for _, v := range paths {
			if Permissionvalidation(Cla.Mps, "rout", v.RoutId, 1) {
				rpath = append(rpath, v)
			}
		}
		rrpath := pathmake(rpath)
		c.JSON(200, gin.H{
			"uid":  Cla.UserId,
			"list": rrpath,
		})
	}
	return
}

func pathmake(path []_struct.AdminRout) interface{} {
	var rpath []super
	for _, v := range path {
		if v.Super == "/" {
			var m metas
			m.Title = v.Title
			m.Icon = v.Icon
			var s super
			s.Path = v.Path
			s.Component = v.Component
			s.Name = v.Name
			s.Meta = m
			s.RoleId = v.RoutId
			rpath = append(rpath, s)
		}
	}
	for i, v := range rpath {
		var rpaths []super
		for _, vv := range path {
			if vv.Super == v.Path {
				var m metas
				m.Title = vv.Title
				m.Icon = vv.Icon
				var s super
				s.Path = vv.Path
				s.Component = vv.Component
				s.Name = vv.Name
				s.Meta = m
				s.RoleId = vv.RoutId
				rpaths = append(rpaths, s)
			}
		}
		rpath[i].Children = rpaths
	}
	return rpath
}
