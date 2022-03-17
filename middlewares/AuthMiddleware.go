package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"plant/common"
	"plant/model"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"Code": 401, "Msg": "1"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"Code": 401, "Msg": "2"})
			c.Abort()
			return
		}

		userId := claims.Id
		fmt.Println(userId)
		db := common.InitDB()
		var user model.User
		db.First(&user, userId)

		if user.Id == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"Code": 401, "Msg": "3"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
