package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
)

const (
	host     = "localhost"
	user     = "midas"
	password = "123qwer321QWER"
	dbname   = "usersdb"
	port     = "5432"
)

func OpenDB() *sql.DB {
	con := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", con)
	if err != nil {
		tracerr.PrintSourceColor(err)
		return nil
	}
	logrus.Println("access opening db")
	return db
}
func AddAdmin(id int, db *sql.DB) (bool, error) {
	err := db.Ping()
	if err != nil {
		return false, err
	}
	if _, err := db.Exec(`UPDATE usersdb SET role ='admin' WHERE id = $1`, id); err != nil {
		return false, err
	}
	logrus.Println("access add new admin with id ", id)
	return true, nil
}
func DelAdmin(id int, db *sql.DB) (bool, error) {
	err := db.Ping()
	if err != nil {
		return false, err
	}
	if _, err := db.Exec(`UPDATE usersdb SET role ='consumer' WHERE id = $1`, id); err != nil {
		return false, err
	}
	logrus.Println("access delete admin with id ", id)
	return true, nil
}
