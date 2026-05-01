package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/waltertaya/cash-flow-forecast-backend/internals/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("auth_token")

		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		claims, err := helpers.ValidateJWT(cookie)

		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
