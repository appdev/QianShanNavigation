package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goNav/util"
	"log"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过http header中的token解析来认证
		token := c.Request.Header.Get("token")
		if token == "" {
			ResponseError(c,401,errors.New("无权限访问"))
			c.Abort()
			return
		}
		log.Print("get token: ", token)
		// 初始化一个JWT对象实例，并根据结构体方法来解析token
		// 解析token中包含的相关信息(有效载荷)
		claims, err := util.ParseToken(token)
		if err != nil {
			// token过期
			if err == TokenExpired {
				ResponseError(c,4001,errors.New("token授权已过期"))
				c.Abort()
				return
			}
			// 其他错误
			ResponseError(c,-1,err)
			c.Abort()
			return
		} // 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set(util.USERNAME, claims.Username)
		c.Set(util.USERID, claims.ID)
		c.Next()
	}
}
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "bgbiao.top" // 签名信息应该设置成动态从库中获取
)
