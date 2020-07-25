package api

import (
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/stormentt/dumbstored/db"
)

func ChangePassword(c *gin.Context) {
	user_id := c.GetInt("user_id")
	b64NewPassword := c.GetHeader("NewPassword")
	if len(b64NewPassword) == 0 {
		c.String(400, "no NewPassword header provided")
		return
	}

	newPassword, err := base64.StdEncoding.DecodeString(b64NewPassword)
	if err != nil {
		c.String(400, "NewPassword header is invalid, must be base64 encoded")
		return
	}

	err = db.ChangeUserPassword(user_id, string(newPassword))
	if err != nil {
		c.String(500, "internal server error")
		return
	}

	c.String(200, "password changed")
}
