package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mlchain/mlchain-sandbox/internal/static"
)

func Auth() gin.HandlerFunc {
	config := static.GetMlchainSandboxGlobalConfigurations()
	return func(c *gin.Context) {
		if config.App.Key != c.GetHeader("X-Api-Key") {
			c.AbortWithStatus(401)
			return
		}
	}
}
