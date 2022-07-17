package service

func newUser() (*chess, *chess, []chessboard) {
	cb, r, g := InitChessboard()
	return r, g, cb
}

// 棋子
type chess struct {
	arr [10][9]storage //棋盘数组
	minister
	vehicle   //车
	commander //帅
	guard     //士
	gun       //炮
	soldier   //兵
	horse     //马
}

type chessboard struct {
	row int     // 行
	col int     // 列
	val storage // 值：哪个棋子
}

// minister 象/相
type minister struct {
	name     string
	location [2]string
}

// vehicle 车
type vehicle struct {
	name     string
	location [2]string
}

// commander 帅/将
type commander struct {
	name     string
	location string
}

// guard 士
type guard struct {
	name     string
	location [2]string
}

// gun 炮
type gun struct {
	name     string
	location [2]string
}

// soldier 兵/卒
type soldier struct {
	name     string
	location [5]string
}

// horse 马
type horse struct {
	name     string
	location [2]string
}

type storage struct {
	name      string
	attribute string // 棋子属性 后来想到应该用0 1 比较好，但是懒得改了
}

func InitChessboard() ([]chessboard, *chess, *chess) {
	// 红方
	var red = &chess{
		minister:  minister{"相", [2]string{"0,2", "0,6"}},
		vehicle:   vehicle{"车", [2]string{"0,0", "0,8"}},
		commander: commander{"帅", "0,4"},
		guard:     guard{"士", [2]string{"0,3", "0,5"}},
		gun:       gun{"炮", [2]string{"2,1", "2,7"}},
		soldier:   soldier{"兵", [5]string{"3,0", "3,2", "3,4", "3,6", "3,8"}},
		horse:     horse{"马", [2]string{"0,1", "0,7"}},
	}
	// 绿方
	var green = &chess{
		minister:  minister{"象", [2]string{"9,2", "9,6"}},
		vehicle:   vehicle{"车", [2]string{"9,0", "9,8"}},
		commander: commander{"将", "9,4"},
		guard:     guard{"士", [2]string{"9,3", "9,5"}},
		gun:       gun{"炮", [2]string{"7,1", "7,7"}},
		soldier:   soldier{"卒", [5]string{"6,0", "6,2", "6,4", "6,6", "6,8"}},
		horse:     horse{"马", [2]string{"9,1", "9,7"}},
	}
	var chessArr [10][9]storage
	var chessNode = chessboard{
		row: 10,
		col: 9,
		val: storage{
			name:      "",
			attribute: "",
		},
	}

	var cb []chessboard

	chessArr[0][0] = storage{name: green.vehicle.name, attribute: "green"}
	chessArr[0][1] = storage{name: green.horse.name, attribute: "green"}
	chessArr[0][2] = storage{name: green.minister.name, attribute: "green"}
	chessArr[0][3] = storage{name: green.guard.name, attribute: "green"}
	chessArr[0][4] = storage{name: green.commander.name, attribute: "green"}
	chessArr[0][5] = storage{name: green.guard.name, attribute: "green"}
	chessArr[0][6] = storage{name: green.minister.name, attribute: "green"}
	chessArr[0][7] = storage{name: green.horse.name, attribute: "green"}
	chessArr[0][8] = storage{name: green.vehicle.name, attribute: "green"}

	chessArr[2][1] = storage{name: green.gun.name, attribute: "green"}
	chessArr[2][7] = storage{name: green.gun.name, attribute: "green"}

	chessArr[3][0] = storage{name: green.soldier.name, attribute: "green"}
	chessArr[3][2] = storage{name: green.soldier.name, attribute: "green"}
	chessArr[3][4] = storage{name: green.soldier.name, attribute: "green"}
	chessArr[3][6] = storage{name: green.soldier.name, attribute: "green"}
	chessArr[3][8] = storage{name: green.soldier.name, attribute: "green"}

	chessArr[9][0] = storage{name: red.vehicle.name, attribute: "red"}
	chessArr[9][1] = storage{name: red.horse.name, attribute: "red"}
	chessArr[9][2] = storage{name: red.minister.name, attribute: "red"}
	chessArr[9][3] = storage{name: red.guard.name, attribute: "red"}
	chessArr[9][4] = storage{name: red.commander.name, attribute: "red"}
	chessArr[9][5] = storage{name: red.guard.name, attribute: "red"}
	chessArr[9][6] = storage{name: red.minister.name, attribute: "red"}
	chessArr[9][7] = storage{name: red.horse.name, attribute: "red"}
	chessArr[9][8] = storage{name: red.vehicle.name, attribute: "red"}

	chessArr[7][1] = storage{name: red.gun.name, attribute: "red"}
	chessArr[7][7] = storage{name: red.gun.name, attribute: "red"}

	chessArr[6][0] = storage{name: red.soldier.name, attribute: "red"}
	chessArr[6][2] = storage{name: red.soldier.name, attribute: "red"}
	chessArr[6][4] = storage{name: red.soldier.name, attribute: "red"}
	chessArr[6][6] = storage{name: red.soldier.name, attribute: "red"}
	chessArr[6][8] = storage{name: red.soldier.name, attribute: "red"}

	cb = append(cb, chessNode)

	for i, v := range chessArr {
		for j, s := range v {
			if s.name != "" {
				chessNode = chessboard{
					row: i,
					col: j,
					val: s,
				}
				cb = append(cb, chessNode)
			}
		}
	}

	return cb, red, green
}
