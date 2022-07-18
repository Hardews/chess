package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	engine := gin.Default()

	engine.POST("/login", Login)
	engine.POST("/register", Register)

	engine.GET("/new", NewRoom)
	engine.GET("/join/:num", JoinRoom)
	engine.GET("/chess/:num", ShowChess)
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

func SetUpRouter() *gin.Engine {
	engine := gin.Default()

	engine.POST("/login", Login)
	engine.POST("/register", Register)

	engine.GET("/new", NewRoom)
	engine.GET("/join/:num", JoinRoom)
	engine.GET("/chess/:num", ShowChess)
	engine.GET("/show", ShowRoom)

	return engine
}
