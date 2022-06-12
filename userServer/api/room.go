package api

import (
	"chess/chessServer"
	"chess/userServer/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var ROOM = make(map[string][]*chessServer.Client)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	roomNum string
	conn    *websocket.Conn
}

func (c *Client) getInstruction(roomNum string, client, rival *chessServer.Client) {
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if err == websocket.ErrCloseSent {
				err = nil
				err = service.DelRoom(roomNum)
				if err != nil {
					fmt.Println(err)
					return
				}
				return
			}
			log.Fatalln(err)
		}

		str := string(msg)
		if len(str) == 1 {
			if str == "1" {
				clients := ROOM[roomNum]
				for i, cli := range clients {
					if cli == client {
						clients[i].PreState = 1
					}
				}
				ROOM[roomNum] = clients
			} else if str == "0" {
				clients := ROOM[roomNum]
				for i, cli := range clients {
					if cli == client {
						clients[i].PreState = 0
					}
				}
				ROOM[roomNum] = clients
			}
		}

		clients := ROOM[roomNum]
		if clients[0].PreState == 1 && clients[1].PreState == 1 {
			time.Sleep(2 * time.Second)
			break
		}
	}

	client.Operation(0, 0, 0, 0, client)

	for true {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if err == websocket.ErrCloseSent {
				err = nil
				err = service.DelRoom(roomNum)
				if err != nil {
					fmt.Println(err)
					return
				}
				return
			}
			log.Fatalln(err)
		}

		str := strings.Split(string(msg), " ")

		if len(str) != 4 {
			continue
		}

		IChoose, INum, IDirection, IStep := str[0], str[1], str[2], str[3]
		choose, _ := strconv.Atoi(IChoose)
		num, _ := strconv.Atoi(INum)
		direction, _ := strconv.Atoi(IDirection)
		step, _ := strconv.Atoi(IStep)
		err, flag := client.Operation(choose, num, direction, step, rival)
		if !flag {
			if err == chessServer.ErrOfCanNotMove {
				c.conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
				continue
			}
			break
		}
		client.Print()
	}
}

// 新房间
func newRoom(ctx *gin.Context) {
	// 升级协议
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 记录房间
	roomNum := RandStringRunes(6)
	err = service.NewRoom(roomNum)
	if err != nil {
		log.Fatalln(err)
	}
	client := &Client{roomNum: roomNum, conn: conn}

	// 启动象棋初始化服务
	var chessClient []*chessServer.Client
	red, green := chessServer.NewClient()
	chessClient = append(chessClient, red)
	chessClient = append(chessClient, green)

	ROOM[roomNum] = chessClient

	go client.getInstruction(roomNum, red, green)
}

func ShowRoom(ctx *gin.Context) {
	rooms, err := service.ShowRoom()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, room := range rooms {
		ctx.JSON(200, gin.H{
			"room_name": room,
		})
	}
}

func JoinRoom(ctx *gin.Context) {
	roomNum := ctx.Param("num")
	clients := ROOM[roomNum]
	if clients[1].State == 0 {
		ctx.JSON(200, gin.H{
			"msg": "房间满人",
		})
		return
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{roomNum: roomNum, conn: conn}
	green := clients[1]
	green.State = 0
	go client.getInstruction(roomNum, green, clients[0])
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
