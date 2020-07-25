package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stormentt/dumbstored/auth"
	"github.com/stormentt/dumbstored/db"
)

func Register(c *gin.Context) {
	username, password, ok := auth.DecodeHeader(c.GetHeader("Authorization"))
	if !ok {
		c.String(400, "bad Authorization header format")
		return
	}

	ok, err := db.CreateUser(username, password)
	if err != nil {
		c.String(500, "%s", err)
		return
	}

	if !ok {
		c.String(409, "username taken")
		return
	}

	user, ok, err := db.GetUserByName(username)
	if err != nil {
		c.String(500, "%s", err)
		return
	}

	if !ok {
		c.String(500, "very odd failure")
		return
	}

	c.String(http.StatusOK, "%s", user)
}
