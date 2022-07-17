package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Client struct {
	State     int // 状态， 1是这个用户等待加入
	PreState  int // 准备状态 ， 1是已准备
	attribute int // "0 red 1 green
	user      *chess
}

func (c *Client) Operation(choose, num, direction, step int, Client *Client) (err error, info bool) {
	var atr string
	if c.attribute == 0 {
		atr = "red"
	} else {
		atr = "green"
	}
	fmt.Println(atr + " move")

	c.sync(Client)

	switch choose {
	case 1:
		err = c.gunMove(num, direction, step)
		if err != nil {
			fmt.Println(err)
		}
	case 2:
		err = c.horseMove(num, direction)
		if err != nil {
			fmt.Println(err)
		}
	case 3:
		err = c.guardMove(num, direction)
		if err != nil {
			fmt.Println(err)
		}
	case 4:
		err = c.vehicleMove(num, direction, step)
		if err != nil {
			fmt.Println(err)
		}
	case 5:
		err = c.ministerMove(num, direction)
		if err != nil {
			fmt.Println(err)
		}
	case 6:
		err = c.commanderMove(direction)
		if err != nil {
			fmt.Println(err)
		}
	case 7:
		err = c.soldierMove(num, direction)
		if err != nil {
			fmt.Println(err)
		}
	}
	if err != nil {
		return err, false
	}
	return nil, true
}

func (c *Client) isWin(client *Client) (bool, int) {
	row, col := getLocation(c.user.commander.location)
	var atr string
	if c.attribute == 0 {
		atr = "red"
	} else {
		atr = "green"
	}

	row1, col1 := getLocation(client.user.commander.location)
	var atr1 string
	if client.attribute == 0 {
		atr1 = "red"
	} else {
		atr1 = "green"
	}

	if c.user.arr[row][col].attribute != atr {
		return true, c.attribute
	} else if client.user.arr[row1][col1].attribute != atr1 {
		return true, client.attribute
	}
	return false, 2
}

func (c *Client) sync(client *Client) {
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
	return &Client{State: 0, PreState: 0, attribute: 0, user: redUser}, &Client{State: 1, PreState: 0, attribute: 1, user: greenUser}
}

func Print(ctx *gin.Context, c *Client) {
	for _, v := range c.user.arr {
		for _, s := range v {
			if s.name == "" {
				fmt.Fprint(ctx.Writer, "   ")
				continue
			}
			fmt.Fprint(ctx.Writer, s.name+" ")
		}
		fmt.Fprintln(ctx.Writer, "")
	}
}

func (c *Client) gunMove(num, direction, stepsCount int) (err error) {
	return c.user.gunMove(num, direction, stepsCount)
}

func (c *Client) vehicleMove(num, direction, stepsCount int) (err error) {
	return c.user.vehicleMove(num, direction, stepsCount)
}

func (c *Client) ministerMove(num, direction int) error {
	return c.user.ministerMove(num, direction)
}

func (c *Client) horseMove(num, direction int) error {
	return c.user.horseMove(num, direction)
}

func (c *Client) commanderMove(direction int) error {
	return c.user.commanderMove(direction)
}

func (c *Client) guardMove(num, direction int) error {
	return c.user.guardMove(num, direction)
}

func (c *Client) soldierMove(num, direction int) error {
	return c.user.soldierMove(num, direction)
}
