package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	engine := gin.Default()

	engine.POST("/login", Login)
	engine.POST("/register", Register)

	engine.GET("/new", newRoom)
	engine.GET("/join/:num", JoinRoom)
	engine.GET("/show", ShowRoom)

	base := engine.Group("/")
	{
		base.Use(JwtToken)
	}

	err := engine.Run()
	if err != nil {
		fmt.Println("engine failed,err:", err)
		return
	}
}
