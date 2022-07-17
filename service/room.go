package service

import (
	"chess/dao"
	"math/rand"
	"time"
)

func NewRoom(roomNum string) error {
	return dao.NewRoom(roomNum)
}

func ShowRoom() ([]string, error) {
	return dao.ShowRoom()
}

func DelRoom(roomNum string) error {
	return dao.DelRoom(roomNum)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
