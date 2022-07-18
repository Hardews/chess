package dao

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mock   sqlmock.Sqlmock
	gormDb *gorm.DB
)

func init() {
	var err error
	var db *sql.DB

	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	gormDb, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	dB = gormDb

}

func TestWriteIn(t *testing.T) {
	mock.ExpectExec("INSERT INTO user").
		WithArgs("1225101", "123456").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// now we execute our method
	if err := WriteIn("1225101", "123456"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestCheckPassword(t *testing.T) {
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	rows := mock.NewRows([]string{"username", "password"}).AddRow(
		"1225101",
		password,
	)
	mock.ExpectQuery("^SELECT password FROM `user` WHERE username = ?").
		WithArgs("1225101").
		WillReturnRows(rows)

	// now we execute our method
	if err, res := CheckPassword("1225101"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	} else {
		fmt.Println(res)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}
