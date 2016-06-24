package serversql

import (
	"absensi-element/session"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql" //ignore
)

var db *sql.DB

type Absensi struct {
	ID        int64
	Nama      string
	Wilayah   string
	Tanggal   time.Time
	Hari      string
	JumlahJam string
}

type User struct {
	ID       int64  `json:"id"`
	Nama     string `json:"nama"`
	Wilayah  string `json:"wilayah"`
	Tipe     string `json:"tipe"`
	Password string `json:"password"`
}

func OpenDB() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(localhost:3306)/webasensi")
	if err != nil {
		panic(err)
	}
	session.Register(db)

}
func (u *User) Insert() error {
	query := "insert into user(nama,wilayah,tipe,password) values(?,?,?,?)"
	result, err := db.Exec(query, u.Nama, u.Wilayah, u.Tipe, u.Password)
	if err != nil {
		return err
	}
	effect, err := result.RowsAffected()
	if err != nil || effect < 1 {
		return fmt.Errorf("tidak ada inputan")

	}
	u.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update(ID int64, data map[string]string) error {
	var columnset string
	for k, v := range data {

		columnset = columnset + k + "=" + v + ","
	}
	columnset = strings.TrimRight(columnset, ",")
	query := fmt.Sprintf("update user set %s where ID=%d", columnset, ID)
	result, err := db.Exec(query, u.Nama, u.Wilayah, u.Tipe, u.Password)
	if err != nil {
		return err
	}
	effect, err := result.RowsAffected()
	if err != nil || effect < 1 {
		return fmt.Errorf("tidak ada inputan")
	}
	u.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Get(ID int64) error {
	query := "select Nama,Wilayah,Password,Tipe from user where ID=?"
	err := db.QueryRow(query, ID).Scan(&u.Nama, &u.Wilayah, &u.Password, &u.Tipe)
	if err != nil {
		return err

	}
	return nil
}

func (u *User) GetBy(Nama string) error {
	query := "select ID,Nama,Wilayah,Password,Tipe from user where Nama=?"
	err := db.QueryRow(query, Nama).Scan(&u.ID, &u.Nama, &u.Wilayah, &u.Password, &u.Tipe)
	if err != nil {
		return err

	}
	return nil
}

func (u *Absensi) Insert() error {
	query := "insert into user values(?,?,?,?,?)"
	result, err := db.Exec(query, u.Nama, u.Wilayah, u.Tanggal, u.Hari, u.JumlahJam)
	if err != nil {
		return err
	}
	effect, err := result.RowsAffected()
	if err != nil || effect < 1 {
		return fmt.Errorf("tidak ada inputan")

	}
	u.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (u *Absensi) Update(ID int64, data map[string]string) error {
	var columnset string
	for k, v := range data {

		columnset = columnset + k + "=" + v + ","
	}
	columnset = strings.TrimRight(columnset, ",")
	query := fmt.Sprintf("update user set %s where ID=%d", columnset, ID)
	result, err := db.Exec(query, u.Nama, u.Wilayah, u.Tanggal, u.Hari, u.JumlahJam)
	if err != nil {
		return err
	}
	effect, err := result.RowsAffected()
	if err != nil || effect < 1 {
		return fmt.Errorf("tidak ada inputan")
	}
	u.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (u *Absensi) Get(ID int64) error {
	query := "select Nama,Wilayah,Password,Tipe from user where ID=?"
	err := db.QueryRow(query, ID).Scan(&u.Nama, &u.Wilayah, &u.Tanggal, &u.Hari, u.JumlahJam)
	if err != nil {
		return err
	}
	return nil
}
