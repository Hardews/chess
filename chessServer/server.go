package chessServer

import (
	"fmt"
	"math/rand"
	"time"
)

type Client struct {
	attribute int
	user      *chess
}

func (c *Client) Sync(client *Client) {
	c.user.arr = client.user.arr
}

func NewClient() (*Client, *Client) {
	redUser, greenUser, cb := newUser()
	var arr [10][9]storage
	for i := range cb {
		if cb[i].val.name != "" {
			arr[cb[i].row][cb[i].col] = cb[i].val
		}
	}
	redUser.arr = arr
	greenUser.arr = arr
	return &Client{user: redUser, attribute: 0}, &Client{attribute: 1, user: greenUser}
}

func (c *Client) Print() {
	switch c.attribute {
	case 1:
		for _, v := range c.user.arr {
			for _, s := range v {
				if s.name == "" {
					fmt.Print("   ")
					continue
				}
				fmt.Print(s.name, " ")
			}
			fmt.Println()
		}
	case 0:
		for i := 9; i >= 0; i-- {
			for j := 8; j >= 0; j-- {
				if c.user.arr[i][j].name == "" {
					fmt.Print("   ")
					continue
				}
				fmt.Print(c.user.arr[i][j].name, " ")
			}
			fmt.Println()
		}
	}

}

func (c *Client) GunMove(num, direction, stepsCount int) (err error) {
	return c.user.gunMove(num, direction, stepsCount)
}

func (c *Client) VehicleMove(num, direction, stepsCount int) (err error) {
	return c.user.vehicleMove(num, direction, stepsCount)
}

func (c *Client) MinisterMove(num, direction int) error {
	return c.user.ministerMove(num, direction)
}

func (c *Client) HorseMove(num, direction int) error {
	return c.user.horseMove(num, direction)
}

func (c *Client) CommanderMove(direction int) error {
	return c.user.commanderMove(direction)
}

func (c *Client) GuardMove(num, direction int) error {
	return c.user.guardMove(num, direction)
}

func (c *Client) SoldierMove(num, direction int) error {
	return c.user.soldierMove(num, direction)
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
