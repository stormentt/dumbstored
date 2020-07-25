package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stormentt/dumbstored/api"
	"github.com/stormentt/dumbstored/config"
	"github.com/stormentt/dumbstored/db"
)

func main() {
	conStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.C.PostgresUser,
		config.C.PostgresPass,
		config.C.PostgresHost,
		config.C.PostgresPort,
		config.C.PostgresName,
	)

	err := db.Connect(conStr)
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.POST("/register", api.Register)
	r.GET("/check", api.Check)

	httpStr := fmt.Sprintf(":%v", config.C.Port)
	fmt.Println(httpStr)
	r.Run(httpStr)
}
