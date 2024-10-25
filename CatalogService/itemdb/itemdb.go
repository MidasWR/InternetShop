package itemdb

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
)

const (
	Host     = "localhost"
	User     = "midas"
	Password = "123qwer321QWER"
	Dbname   = "postgres"
	Port     = "5432"
)

func OpenDB() *sql.DB {
	con := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)
	db, err := sql.Open("postgres", con)
	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	}
	err = db.Ping()
	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	}
	logrus.Println("Access opening db")
	return db
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
	rows, err := db.Query(`SELECT * FROM itemsdb`)
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
