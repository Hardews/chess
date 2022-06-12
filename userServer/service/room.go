package service

import "chess/userServer/dao"

func NewRoom(roomNum string) error {
	return dao.NewRoom(roomNum)
}

func ShowRoom() ([]string, error) {
	return dao.ShowRoom()
}
