package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stormentt/dumbstored/auth"
	"github.com/stormentt/dumbstored/db"
)

func Auth(c *gin.Context) {
	username, password, ok := auth.DecodeHeader(c.GetHeader("Authorization"))
	if !ok {
		c.String(400, "bad Authorization header format")
		c.Abort()
		return
	}

	user, ok, err := db.GetUserByName(username)
	if err != nil {
		c.String(500, "internal server error", err)
		c.Abort()
		return
	}

	badLoginString := "bad username or password"
	badLoginCode := 401

	if !ok {
		// defeat the username enumeration timing attack
		auth.DummyCheckPassword()
		c.String(badLoginCode, badLoginString)
		c.Abort()
		return
	}

	ok = auth.CheckPassword(password, user.Password, user.Salt)
	if !ok {
		c.String(badLoginCode, badLoginString)
		c.Abort()
		return
	}

	fmt.Printf("user id %d\n", user.ID)
	c.Set("user_id", user.ID)
	c.Next()
}
