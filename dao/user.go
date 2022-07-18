package dao

func CheckPassword(username string) (error, string) {
	var checkPwd struct {
		username string
		password string
	}
	tx := dB.Table("user").Select("password").Where("username = ?", username).Scan(&checkPwd)

	return tx.Error, checkPwd.password
}

func CheckUsername(username string) error {
	tx := dB.Table("user").Select("username").Where("username = ", username)
	return tx.Error
}

func WriteIn(username, password string) error {
	tx := dB.Table("user").Select("username", "password").Create([]string{username, password})
	return tx.Error
}
