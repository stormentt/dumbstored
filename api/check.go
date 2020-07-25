package api

import (
	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	user_id := c.GetInt("user_id")

	c.String(200, "hello %d", user_id)
}
