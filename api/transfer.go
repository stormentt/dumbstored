package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stormentt/dumbstored/db"
)

func StartTransfer(c *gin.Context) {
	user_id := c.GetInt("user_id")
	length := c.GetHeader("Length")
	if len(length) == 0 {
		c.String(400, "no Length header")
		return
	}

	final_size, err := strconv.ParseInt(length, 10, 64)
	if err != nil {
		c.String(400, "bad Length header")
		return
	}

	id, err := db.CreateTransfer(user_id, final_size)
	if err != nil {
		c.String(500, "internal server error")
		return
	}

	rstruct := struct {
		TransferID int
	}{id}

	c.JSON(201, rstruct)
}

func ContinueTransfer(c *gin.Context) {
	param_id := c.Param("id")
	transfer_id, err := strconv.ParseInt(param_id, 10, 32)
	if err != nil {
		c.String(400, "transfer_id must be an integer")
		return
	}

	rstruct := struct {
		TransferID int
	}{int(transfer_id)}

	c.JSON(200, rstruct)
}
