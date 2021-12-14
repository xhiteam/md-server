package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/model/request"
)

func GetClaims(c *gin.Context) (*request.CustomClaims,error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.MD_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims,err
}

// 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl,err:= GetClaims(c);err!=nil{
			return 0
		}else{
			return cl.ID
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.ID
	}

}


// 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *request.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl,err:= GetClaims(c);err!=nil{
			return nil
		}else{
			return cl
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse
	}
}
