package service

import (
	"chess/userServer/dao"
	"database/sql"
)

func CheckPassword(username, password string) (error, bool) {
	err, checkPwd := dao.CheckPassword(username)
	if err != nil {
		return err, false
	}
	err, res := Interpretation(checkPwd, password)
	return err, res
}

func CheckUsername(username string) (error, bool) {
	err := dao.CheckUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err, true
		}
		return err, false
	}
	return err, false
}

func WriteIn(username, password string) error {
	err := dao.WriteIn(username, password)
	if err != nil {
		return err
	}
	return err
}
