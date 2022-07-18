package api

import (
	"chess/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

//const (
//	ready   = "1"
//	noReady = "0"
//)

var (
	ROOM = make(map[string][]*service.Client)
	mu   sync.Mutex
)

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

func (c *Client) getInstruction(client, rival *service.Client) {
	for true {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if flag := websocket.IsCloseError(err, websocket.CloseNormalClosure); flag {
				err = nil
				//err = service2.DelRoom(roomNum)
				//if err != nil {
				//	fmt.Println(err)
				//	return
				//}
				return
			}
			fmt.Println(err)
		}

		str := strings.Split(string(msg), " ")

		if len(str) != 4 {
			continue
		}

		IChoose, INum, IDirection, IStep := str[0], str[1], str[2], str[3]
		choose, err := strconv.Atoi(IChoose)
		num, err := strconv.Atoi(INum)
		direction, err := strconv.Atoi(IDirection)
		step, err := strconv.Atoi(IStep)
		err, flag := client.Operation(choose, num, direction, step, rival)
		if err != nil {
			fmt.Println("parse error : ", err)
		}
		if !flag {
			if err == service.ErrOfCanNotMove {
				c.conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
				continue
			}
			break
		}
	}
}

// NewRoom 新房间
func NewRoom(ctx *gin.Context) {
	// 升级协议
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 记录房间
	mu.Lock()
	roomNum := "Hardews"
	//err = service2.NewRoom(roomNum)
	if err != nil {
		log.Fatalln(err)
	}
	client := &Client{roomNum: roomNum, conn: conn}

	// 启动象棋初始化服务
	var chessClient []*service.Client
	red, green := service.NewClient()
	chessClient = append(chessClient, red)
	chessClient = append(chessClient, green)

	ROOM[roomNum] = chessClient
	mu.Unlock()

	go client.getInstruction(red, green)
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
	mu.Lock()
	//roomNum := ctx.Param("num")
	roomNum := "Hardews"
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
	mu.Unlock()
	client := &Client{roomNum: roomNum, conn: conn}
	green := clients[1]
	green.State = 0
	go client.getInstruction(green, clients[0])
}

func ShowChess(ctx *gin.Context) {
	roomNum := ctx.Param("num")
	//roomNum := "Hardews"
	clients, ok := ROOM[roomNum]
	if !ok {
		ctx.JSON(200, gin.H{
			"msg": "no this room",
		})
		return
	}

	fmt.Fprintln(ctx.Writer, "房间"+roomNum+"的比赛")
	fmt.Fprint(ctx.Writer, "\n\n\n")

	service.Print(ctx, clients[0])
}
