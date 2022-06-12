package dao

func NewRoom(roomNum string) error {
	cmd := rdb.SAdd("room", roomNum)
	if err := cmd.Err(); err != nil {
		return err
	}
	return nil
}

func ShowRoom() ([]string, error) {
	cmd := rdb.SMembers("room")
	return cmd.Result()
}

func DelRoom(roomNum string) error {
	cmd := rdb.Del(roomNum)
	return cmd.Err()
}
