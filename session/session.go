package session

import (
	"database/sql"
	"fmt"
	"time"
)

var dbs *sql.DB

type Session struct {
	ID    int
	User  int64
	Waktu time.Time
}

func Register(db *sql.DB) {
	dbs = db
}

func (u *Session) GetUserID(ID int) error {
	query := "select   ID,user,waktu_login from user where User=?"
	err := dbs.QueryRow(query, ID).Scan(&u.ID, &u.User, &u.Waktu)
	if err != nil {
		return err

	}
	return nil
}

func (u *Session) Insert() error {
	query := "insert into session (user, waktu_login) values(?,?)"
	result, err := dbs.Exec(query, u.User, u.Waktu)
	if err != nil {
		return err
	}
	effect, err := result.RowsAffected()
	if err != nil || effect < 1 {
		return fmt.Errorf("tidak ada inputan")

	}
	return nil
}
