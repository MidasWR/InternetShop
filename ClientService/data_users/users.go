package data_users

import (
	"database/sql"
	"errors"
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

func NewDB() *sql.DB {
	con := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", con)
	if err != nil {
		tracerr.PrintSourceColor(err)
		return nil
	}
	logrus.Println("Access opening db")
	return db
}
func CheckUsernameExists(username string, db *sql.DB) (bool, error) {
	var username2 string
	if err := db.QueryRow(`SELECT username FROM usersdb WHERE username=$1`, username).Scan(&username2); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return false, err
		}
		return false, nil
	}
	logrus.Println("username checked without errors")
	return username == username2, nil
}
func CheckLoginExists(login string, db *sql.DB) (bool, error) {
	var login2 string
	if err := db.Ping(); err != nil {
		return false, err
	}
	if err := db.QueryRow(`SELECT login FROM usersdb WHERE login=$1`, login).Scan(&login2); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return false, err
		}
		return false, nil
	}
	logrus.Println("login checked without errors")
	return login == login2, nil
}
func AddNewMember(username string, login string, hashPass string, role string, db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}
	if _, err := db.Exec(`INSERT INTO usersdb (username, login, password_hashed, role) VALUES ($1, $2, $3, $4)`,
		username, login, hashPass, role); err != nil {
		return err
	}
	logrus.Println("new member created")
	return nil
}
func GetHashPasswordByLogin(db *sql.DB, login string) (string, error) {
	var hash string
	if err := db.Ping(); err != nil {
		return "", err
	}
	if err := db.QueryRow(`SELECT password_hashed FROM usersdb WHERE login=$1`, login).Scan(&hash); err != nil {
		return "", err
	}
	logrus.Println("hash password got access")
	return hash, nil
}
func GetRoleByLogin(db *sql.DB, login string) (string, error) {
	var role string
	if err := db.Ping(); err != nil {
		return "", err
	}
	if err := db.QueryRow(`SELECT role FROM usersdb WHERE login=$1`, login).Scan(&role); err != nil {
		return "", err
	}
	logrus.Println("role got access")
	return role, nil
}
func GetIDByLogin(db *sql.DB, login string) (int, error) {
	var id int
	if err := db.Ping(); err != nil {
		return 0, err
	}
	if err := db.QueryRow(`SELECT id FROM usersdb WHERE login=$1`, login).Scan(&id); err != nil {
		return 0, err
	}
	logrus.Println("id got access")
	return id, nil
}
