package api

import (
	service2 "chess/service"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string
	Password string
}

func Login(ctx *gin.Context) {
	var user User
	var res bool
	user.Username, res = ctx.GetPostForm("username")
	if !res {
		ctx.JSON(200, "输入的账号为空")
		return
	}
	user.Password, res = ctx.GetPostForm("password")
	if !res {
		ctx.JSON(200, "输入的密码为空")
		return
	}

	// 检查账号是否正确
	err, res := service2.CheckPassword(user.Username, user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(200, "无此账号")
			return
		}
		fmt.Println(err)
		ctx.JSON(200, "internet error")
		return
	}
	if res {
		token, flag := SetToken(user.Username)
		if !flag {
			ctx.JSON(500, "internet error")
			return
		}
		ctx.JSON(200, gin.H{
			"msg": token,
		})
		ctx.JSON(200, "successful")
	} else {
		ctx.JSON(200, "密码错误")
		return
	}

}

func Register(ctx *gin.Context) {
	var user User
	user.Username, _ = ctx.GetPostForm("username")
	user.Password, _ = ctx.GetPostForm("password")

	if user.Username == "" {
		ctx.JSON(200, "用户名为空")
		return
	}
	if user.Password == "" {
		ctx.JSON(200, "密码为空")
		return
	}

	err, flag := service2.CheckUsername(user.Username)
	if err != nil {
		ctx.JSON(500, "internet error")
		fmt.Println("check username failed, error: ", err)
		return
	}
	if flag == false {
		ctx.JSON(200, "用户名已存在!")
		return
	}

	err, user.Password = service2.Encryption(user.Password)
	if err != nil {
		ctx.JSON(500, "internet error")
		fmt.Println("encryption failed , err :", err)
		return
	}

	err = service2.WriteIn(user.Username, user.Password)
	if err != nil {
		ctx.JSON(500, "internet error")
		fmt.Println("register failed,err:", err)
		return
	}

	ctx.JSON(200, "successful")

}
