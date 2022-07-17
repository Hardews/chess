package dao

func CheckPassword(username string) (error, string) {
	var checkPwd string
	tx, err := dB.Prepare("select username,password from user where username = ?")
	if err != nil {
		return err, checkPwd
	}
	err = tx.QueryRow(username).Scan(&username, &checkPwd)
	if err != nil {
		return err, checkPwd
	}
	return err, checkPwd
}

func CheckUsername(username string) error {
	var checkUsername string
	tx, err := dB.Prepare("select username from user where username = ?")
	if err != nil {
		return err
	}

	err = tx.QueryRow(username).Scan(checkUsername)
	if err != nil {
		return err
	}
	return err
}
func WriteIn(username, password string) error {
	tx, err := dB.Prepare("insert into user (username,password) values (?,?)")
	if err != nil {
		return err
	}
	_, err = tx.Exec(username, password)
	if err != nil {
		return err
	}
	return err
}
