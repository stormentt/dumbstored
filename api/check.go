package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stormentt/dumbstored/auth"
	"github.com/stormentt/dumbstored/db"
)

func Check(c *gin.Context) {
	username, password, ok := auth.DecodeHeader(c.GetHeader("Authorization"))
	if !ok {
		c.String(400, "bad Authorization header format")
		return
	}

	user, ok, err := db.GetUserByName(username)
	if err != nil {
		// todo: generic error
		c.String(500, "%s", err)
		return
	}

	if !ok {
		// todo: generic error
		//c.String(401, "bad username or password")
		auth.DummyCheckPassword()
		c.String(401, "bad username")
		return
	}

	ok = auth.CheckPassword(password, user.Password, user.Salt)
	if !ok {
		// todo: generic error
		//c.String(401, "bad username or password")
		c.String(401, "bad password")
		return
	}

	c.String(200, "auth succeeded")
}
