package middlewares

import (
	"github.com/gin-gonic/gin"
)

func PushUpdate(c *gin.Context) {
	c.String(200, "push_update")
}
