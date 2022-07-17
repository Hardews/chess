package main

import (
	"chess/api"
	"chess/dao"
)

func main() {
	//dao.InitDB()
	dao.InitClient()
	api.InitRouter()
}
