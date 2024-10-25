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
	dbname   = "wishlist"
	port     = "5432"
)

func OpenDB() *sql.DB {
	con := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", con)
	if err != nil {
		tracerr.PrintSourceColor(err)
		return nil
	}
	logrus.Println("Access opening db")
	return db
}
func AddToWishlist(db *sql.DB, id int) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	if _, err := db.Exec(`INSERT INTO wishlist (name, type, amount, cost)
	SELECT name, type, amount, cost
	FROM itemsdb
	WHERE id = $1;
	`, id); err != nil {
		return err
	}
	return nil
}
func DeleteFromWishlist(db *sql.DB, id int) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	if _, err := db.Exec(`DROP * FROM wishlist WHERE id = $1,id`); err != nil {
		return err
	}
	return nil
}

type Item struct {
	ID     int32   `json:"id"`
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Amount int32   `json:"amount"`
	Cost   float32 `json:"cost"`
}

func GetDataItems(db *sql.DB) ([]Item, error) {
	var items []Item
	rows, err := db.Query(`SELECT * FROM wishlist`)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	defer rows.Close()
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Type, &item.Amount, &item.Cost); err != nil {
			return nil, tracerr.Wrap(err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, tracerr.Wrap(err)
	}
	logrus.Println("Access got items")
	return items, nil
}
