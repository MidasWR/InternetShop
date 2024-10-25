package itemsdb

import (
	AddService "InternetShop/AdderCatalogService/proto"
	"InternetShop/CatalogService/itemdb"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Add(item *AddService.Item) error {
	con := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", itemdb.Host, itemdb.Port, itemdb.User, itemdb.Password, itemdb.Dbname)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	if _, err = db.Exec(`INSERT INTO itemsdb(name,type,amount,cost) VALUES ($1,$2,$3,$4);`, item.Name, item.Type, item.Amount, item.Cost); err != nil {
		return err
	}
	logrus.Println("access adding item")
	return nil
}
func Delete(id *AddService.Id) error {
	con := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", itemdb.Host, itemdb.Port, itemdb.User, itemdb.Password, itemdb.Dbname)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	if _, err = db.Exec(`DELETE FROM itemsdb WHERE id = $1;`, id); err != nil {
		return err
	}
	logrus.Println("access deleting item")
	return nil
}
