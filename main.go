package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stormentt/dumbstored/api"
	"github.com/stormentt/dumbstored/config"
	"github.com/stormentt/dumbstored/db"
	"github.com/stormentt/dumbstored/middleware"
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
	authRoutes := r.Group("/", middleware.Auth)
	{
		authRoutes.GET("/check", api.Check)
		authRoutes.POST("/change-pw", api.ChangePassword)
		authRoutes.POST("/store", api.StartTransfer)
		authRoutes.POST("/store/:id", api.ContinueTransfer)
	}

	httpStr := fmt.Sprintf(":%v", config.C.Port)
	fmt.Println(httpStr)
	r.Run(httpStr)
}
