package main

import (
	"chess/userServer/api"
	"chess/userServer/dao"
)

func main() {
	dao.InitDB()
	dao.InitClient()
	api.InitRouter()
}
