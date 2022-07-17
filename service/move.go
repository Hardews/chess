package service

import (
	"errors"
	"strconv"
	"strings"
)

// direction 方向
const (
	north      = 1
	northeast  = 12
	northeastE = 122
	east       = 2
	southeast  = 23
	southeastE = 233
	south      = 3
	southwest  = 34
	southwestE = 344
	west       = 4
	northwest  = 41
	northwestE = 411
)

var ErrOfCanNotMove = errors.New("can not move")

func getLocation(location string) (row int, col int) {
	locArr := strings.Split(location, ",")
	row, _ = strconv.Atoi(locArr[0])
	col, _ = strconv.Atoi(locArr[1])
	return
}

// 因为是镜像，所以要翻转
func oppositeDirection(direction int) int {
	switch direction {
	case north:
		return 3
	case northeast:
		return 34
	case northeastE:
		return 344
	case east:
		return 4
	case southeast:
		return 41
	case southeastE:
		return 411
	case south:
		return 1
	case southwest:
		return 12
	case southwestE:
		return 122
	case west:
		return 2
	case northwest:
		return 23
	case northwestE:
		return 233
	default:
		return 0
	}
}

func newRowAndCol(row, col int) (string, string) {
	return strconv.Itoa(row), strconv.Itoa(col)
}

func (c *chess) gunMove(num, direction, stepsCount int) error {
	row, col := getLocation(c.gun.location[num])
	var count = 0
	var nRow, nCol = row, col
	var chessInRoad []int

	if c.arr[row][col].attribute == "green" {
		direction = oppositeDirection(direction)
	}

	switch direction {
	case north:
		nRow = row - stepsCount
		if row-stepsCount < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		// 遍历炮所在的这一列
		for i := row - 1; i >= row-stepsCount; i-- {
			for j, sto := range c.arr[i] {
				if j == col {
					if sto.name != "" {
						count++
						chessInRoad = append(chessInRoad, i)
					}
				}
			}
		}
	case east:
		nCol = col + stepsCount
		if col+stepsCount > 8 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		// 遍历炮所在的这一行
		for n := col + 1; n <= col+stepsCount; n++ {
			// 如果前进的路上有棋子
			if c.arr[row][n].name != "" {
				// 计数加一
				count++
				// 将该棋子加入到切片以方便后边处理
				chessInRoad = append(chessInRoad, n)
			}
		}
	case south:
		nRow = row + stepsCount
		if row+stepsCount > 9 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		// 遍历炮所在的这一列
		for i := row + 1; i <= row+stepsCount; i++ {
			for j, sto := range c.arr[i] {
				if j == col {
					if sto.name != "" {
						count++
						chessInRoad = append(chessInRoad, i)
					}
				}
			}
		}
	case west:
		nCol = col - stepsCount
		if col-stepsCount < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		// 遍历炮所在的这一行
		for n := col - 1; n >= col-stepsCount; n-- {
			// 如果前进的路上有棋子
			if c.arr[row][n].name != "" {
				// 计数加一
				count++
				// 将该棋子加入到切片以方便后边处理
				chessInRoad = append(chessInRoad, n)
			}
		}
	default:
		return ErrOfCanNotMove
	}

	switch {
	// 如果大于2,炮不能前进
	case count > 2 || count == 1:
		return ErrOfCanNotMove
	case count == 2:
		// 等于2，判断是否为吃子，不是则不能移动
		if nRow == chessInRoad[1] || nCol == chessInRoad[1] {
			// 如果是友方棋子，不能吃
			if c.arr[row][col].attribute == c.arr[nRow][nCol].attribute {
				return ErrOfCanNotMove
			}
		} else {
			return ErrOfCanNotMove
		}
	case count == 0:
		// 等于0 直接移动
	}

	// 这里是棋子原来的位置变为空位，然后棋子移动到的位置变为棋子的位置
	val := c.arr[row][col]
	c.arr[row][col] = storage{}
	c.arr[nRow][nCol] = val
	// 更新棋子的坐标
	iRow, iCol := newRowAndCol(nRow, nCol)
	c.gun.location[num] = iRow + "," + iCol
	return nil
}

func (c *chess) vehicleMove(num, direction, stepsCount int) error {
	row, col := getLocation(c.vehicle.location[num])
	var count = 0
	var nRow, nCol = row, col
	var chessInRoad []int

	switch direction {
	case north:
		nRow = row - stepsCount
		if nRow < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		// 遍历车所在的这一列
		for i := row - 1; i <= row-stepsCount; i-- {
			for j, sto := range c.arr[i] {
				if j == col {
					if sto.name != "" {
						count++
						chessInRoad = append(chessInRoad, i)
					}
				}
			}
		}
	case east:
		nCol = col + stepsCount
		if col+stepsCount > 8 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		// 遍历炮所在的这一行
		for n := col + 1; n >= col+stepsCount; n++ {
			// 如果前进的路上有棋子
			if c.arr[row][n].name != "" {
				// 计数加一
				count++
				// 将该棋子加入到切片以方便后边处理
				chessInRoad = append(chessInRoad, n)
			}
		}
	case south:
		nRow = row + stepsCount
		if row+stepsCount > 9 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		// 遍历炮所在的这一列
		for i := row + 1; i <= row+stepsCount; i++ {
			for j, sto := range c.arr[i] {
				if j == col {
					if sto.name != "" {
						count++
						chessInRoad = append(chessInRoad, i)
					}
				}
			}
		}
	case west:
		nCol = col - stepsCount
		if col-stepsCount < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		// 遍历车所在的这一行
		for n := col - 1; n <= col-stepsCount; n-- {
			// 如果前进的路上有棋子
			if c.arr[row][n].name != "" {
				// 计数加一
				count++
				// 将该棋子加入到切片以方便后边处理
				chessInRoad = append(chessInRoad, n)
			}
		}

	default:
		return ErrOfCanNotMove
	}

	switch {
	// 如果大于1,车不能前进
	case count > 1:
		return ErrOfCanNotMove
	case count == 1:
		// 等于1，判断是否为吃子，不是则不能移动
		if nCol == chessInRoad[0] || nRow == chessInRoad[0] {
			// 如果是右方棋子，不能吃
			if c.arr[row][col].attribute == c.arr[nRow][nCol].attribute {
				return ErrOfCanNotMove
			}
		} else {
			return ErrOfCanNotMove
		}
	case count == 0:
	// 等于0 直接移动
	default:
		return ErrOfCanNotMove
	}

	// 这里是棋子原来的位置变为空位，然后棋子移动到的位置变为棋子的位置
	val := c.arr[row][col]
	c.arr[row][col] = storage{}
	c.arr[nRow][nCol] = val
	iRow, iCol := newRowAndCol(nRow, nCol)
	c.vehicle.location[num] = iRow + "," + iCol
	return nil
}

func (c *chess) commanderMove(direction int) error {
	row, col := getLocation(c.commander.location)
	var nCol = col
	var nRow = row
	var stepsCount = 1
	var count = 0

	if c.arr[row][col].attribute == "green" {
		direction = oppositeDirection(direction)
	}

	switch direction {
	case north:
		nRow = row - stepsCount
		switch c.arr[row][col].attribute {
		case "green":
			if nRow < 0 {
				// 如果移动的步数超过棋盘范围
				return ErrOfCanNotMove
			}
		case "red":
			if nRow < 7 {
				// 如果移动的步数超过棋盘范围
				return ErrOfCanNotMove
			}
		}

	case east:
		nCol = col + stepsCount
		if nCol > 5 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}
	case south:
		nRow = row + stepsCount
		switch c.arr[row][col].attribute {
		case "green":
			if nRow > 2 {
				// 如果移动的步数超过棋盘范围
				return ErrOfCanNotMove
			}
		case "red":
			if nRow > 9 {
				// 如果移动的步数超过棋盘范围
				return ErrOfCanNotMove
			}
		}
	case west:
		nCol = col - stepsCount
		if nCol < 3 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}
		if c.arr[nRow][nCol].name != "" {
			count++
		}
	default:
		return ErrOfCanNotMove
	}

	switch {
	case count == 1:
		// 等于1，判断是否为吃子，不是则不能移动
		// 如果是友方棋子，不能吃
		if c.arr[row][col].attribute == c.arr[nRow][nCol].attribute {
			return ErrOfCanNotMove
		}

	case count == 0:
	// 等于0 直接移动
	default:
		return ErrOfCanNotMove
	}
	// 这里是棋子原来的位置变为空位，然后棋子移动到的位置变为棋子的位置
	val := c.arr[row][col]
	c.arr[row][col] = storage{}
	c.arr[nRow][nCol] = val
	iRow, iCol := newRowAndCol(nRow, nCol)
	c.commander.location = iRow + "," + iCol
	return nil
}

func (c *chess) soldierMove(num, direction int) error {
	row, col := getLocation(c.soldier.location[num])
	var stepsCount = 1
	var nRow, nCol = row, col
	var count = 0

	// 兵不能后退
	if c.arr[row][col].attribute == "red" {
		if direction == south {
			return ErrOfCanNotMove
		}
	} else {
		direction = oppositeDirection(direction)
	}

	switch direction {
	case north:
		nRow = row - stepsCount
		if row-stepsCount < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[row-stepsCount][col].name != "" {
			count++
		}
	case east:
		nCol = col + stepsCount
		if col+stepsCount > 8 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[row][col+stepsCount].name != "" {
			count++
		}
	case south:
		nRow = row + stepsCount
		if row+stepsCount > 8 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[row+stepsCount][col].name != "" {
			count++
		}
	case west:
		nCol = col - stepsCount
		if col-stepsCount < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[row][col-stepsCount].name != "" {
			count++
		}
	}
	switch {
	case count == 1:
		// 等于1，判断是否为吃子，不是则不能移动
		// 如果是友方棋子，不能吃
		if c.arr[row][col].attribute == c.arr[nRow][nCol].attribute {
			return ErrOfCanNotMove
		}
		// 这里是棋子原来的位置变为空位，然后棋子移动到的位置变为棋子的位置
	case count == 0:
		// 等于0 直接移动
	default:
		return ErrOfCanNotMove
	}
	val := c.arr[row][col]
	c.arr[row][col] = storage{}
	c.arr[nRow][nCol] = val
	// 更新棋子的坐标
	iRow, iCol := newRowAndCol(nRow, nCol)
	c.soldier.location[num] = iRow + "," + iCol
	return nil
}

func (c *chess) guardMove(num, direction int) error {
	row, col := getLocation(c.guard.location[num])
	var nRow, nCol = row, col
	var count = 0

	if c.arr[row][col].attribute == "red" {

	}

	switch direction {
	case northeast:
		nRow, nCol = row-1, col+1
		switch c.arr[row][col].attribute {
		case "red":
			if col+1 > 5 || row-1 < 0 {
				// 如果移动的步数超过棋盘范围
				return ErrOfCanNotMove
			}
		case "green":
			if col+1 > 5 || row-1 < 7 {
				// 如果移动的步数超过棋盘范围
				return ErrOfCanNotMove
			}
		}

		if c.arr[row-1][col+1].name != "" {
			count++
		}

	case southeast:
		nRow, nCol = row+1, col+1
		switch c.arr[row][col].attribute {
		case "red":
			if col+1 > 5 || row+1 > 2 {
				// 如果移动的步数超过棋盘范围
				return ErrOfCanNotMove
			}
		case "green":
			if col+1 > 5 || row+1 > 9 {
				// 如果移动的步数超过棋盘范围
				return ErrOfCanNotMove
			}
		}

		if c.arr[row+1][col+1].name != "" {
			count++
		}

	case southwest:
		nRow, nCol = row+1, col-1

		switch c.arr[row][col].attribute {
		case "green":
			if col-1 < 3 || row+1 > 2 {
				// 如果移动的步数超过棋盘范围
				return ErrOfCanNotMove
			}
		case "red":
			if col-1 > 3 || row+1 > 9 {
				// 如果移动的步数超过棋盘范围
				return ErrOfCanNotMove
			}
		}

		if c.arr[row+1][col-1].name != "" {
			count++
		}

	case northwest:
		nRow, nCol = row-1, col-1

		switch c.arr[row][col].attribute {
		case "red":
			if col-1 < 3 || row-1 < 0 {
				// 如果移动的步数超过棋盘范围
				return ErrOfCanNotMove
			}
		case "green":
			if col-1 < 3 || row-1 < 7 {
				// 如果移动的步数超过棋盘范围
				return ErrOfCanNotMove
			}
		}

		if c.arr[row-1][col-1].name != "" {
			count++
		}

	}
	switch {
	case count == 1:
		// 等于1，判断是否为吃子，不是则不能移动
		// 如果是友方棋子，不能吃
		if c.arr[row][col].attribute == c.arr[nRow][nCol].attribute {
			return ErrOfCanNotMove
		}
	case count == 0:
	// 等于0 直接移动
	default:
		return ErrOfCanNotMove
	}

	// 这里是棋子原来的位置变为空位，然后棋子移动到的位置变为棋子的位置
	val := c.arr[row][col]
	c.arr[row][col] = storage{}
	c.arr[nRow][nCol] = val
	// 更新棋子的坐标
	iRow, iCol := newRowAndCol(nRow, nCol)
	c.guard.location[num] = iRow + "," + iCol
	return nil
}

func (c *chess) ministerMove(num, direction int) error {
	row, col := getLocation(c.minister.location[num])
	var nRow, nCol = row, col
	var count = 0

	if c.arr[row][col].attribute == "green" {
		direction = oppositeDirection(direction)
	}

	switch direction {
	case northeast:
		nRow, nCol = row-2, col+2
		if col+2 > 9 || row-2 < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[row-2][col+2].name != "" {
			count++
		}
	case southeast:
		nRow, nCol = row+2, col+2
		if col+2 > 9 || row+2 > 8 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[row+2][col+2].name != "" {
			count++
		}

	case southwest:
		nRow, nCol = row+2, col-2
		if row+2 > 9 || col-2 < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[row+2][col-2].name != "" {
			count++
		}

	case northwest:
		nRow, nCol = row-2, col-2
		if col-2 < 0 || row-2 < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[row-2][col-2].name != "" {
			count++
		}
	default:
		return ErrOfCanNotMove
	}

	switch {
	case count == 1:
		// 等于1，判断是否为吃子，不是则不能移动
		// 如果是友方棋子，不能吃
		if c.arr[row][col].attribute == c.arr[nRow][nCol].attribute {
			return ErrOfCanNotMove
		}
	case count == 0:
	// 等于0 直接移动
	default:
		return ErrOfCanNotMove
	}

	// 这里是棋子原来的位置变为空位，然后棋子移动到的位置变为棋子的位置
	val := c.arr[row][col]
	c.arr[row][col] = storage{}
	c.arr[nRow][nCol] = val
	// 更新棋子的坐标
	iRow, iCol := newRowAndCol(nRow, nCol)
	c.minister.location[num] = iRow + "," + iCol
	return nil
}

func (c *chess) horseMove(num, direction int) error {
	row, col := getLocation(c.horse.location[num])
	nRow, nCol := row, col
	var count = 0

	if c.arr[row][col].attribute == "green" {
		direction = oppositeDirection(direction)
	}

	switch direction {
	case northeast:
		nRow, nCol = row-2, col+1
		if col+1 > 8 || row-2 < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[row-2][col+1].name != "" {
			count++
		}

	case southwestE:
		nRow, nCol = row+1, col-2
		if row+1 > 8 || col-2 < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[nRow][nCol].name != "" {
			count++
		}

	case southeast:
		nRow, nCol = row+2, col+1
		if col+1 > 9 || row+2 > 8 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[row+2][col+1].name != "" {
			count++
		}

	case southeastE:
		nRow, nCol = row+1, col+2
		if col+2 > 9 || row+1 > 8 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[nRow][nCol].name != "" {
			count++
		}

	case southwest:
		nRow, nCol = row+2, col-1
		if row+2 > 9 || col-1 < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[row+2][col-1].name != "" {
			count++
		}

	case northeastE:
		nRow, nCol = row-1, col+2
		if col+2 > 9 || row-1 < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[nRow][nCol].name != "" {
			count++
		}

	case northwest:
		nRow, nCol = row-2, col-1
		if col-1 < 0 || row-2 < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[row-2][col-1].name != "" {
			count++
		}

	case northwestE:
		nRow, nCol = row-1, col-2
		if col-2 < 0 || row-1 < 0 {
			// 如果移动的步数超过棋盘范围
			return ErrOfCanNotMove
		}

		if c.arr[row-1][col-2].name != "" {
			count++
		}
	default:
		return ErrOfCanNotMove
	}

	switch {
	case count == 1:
		// 等于1，判断是否为吃子，不是则不能移动
		// 如果是友方棋子，不能吃
		if c.arr[row][col].attribute == c.arr[nRow][nCol].attribute {
			return ErrOfCanNotMove
		}
	case count == 0:
		// 等于0 直接移动
	default:
		return ErrOfCanNotMove
	}

	// 这里是棋子原来的位置变为空位，然后棋子移动到的位置变为棋子的位置
	val := c.arr[row][col]
	c.arr[row][col] = storage{}
	c.arr[nRow][nCol] = val
	// 更新棋子的坐标
	iRow, iCol := newRowAndCol(nRow, nCol)
	c.horse.location[num] = iRow + "," + iCol

	return nil
}
