package server

import (
	"InternetShop/AdderCatalogService/itemsdb"
	AddService "InternetShop/AdderCatalogService/proto"
	"context"
	"database/sql"
)

type Server struct {
	AddService.UnimplementedAdderCatalogServer
	Db *sql.DB
}

func (s *Server) AddItemToCatalog(ctx context.Context, item *AddService.Item) (*AddService.Ok, error) {
	if err := itemsdb.Add(item); err != nil {
		ok := AddService.Ok{
			Ok: false,
		}
		return &ok, err
	}
	ok := AddService.Ok{
		Ok: true,
	}
	return &ok, nil
}
func (s *Server) DeleteItemFromCatalog(ctx context.Context, id *AddService.Id) (*AddService.Ok, error) {
	if err := itemsdb.Delete(id); err != nil {
		ok := AddService.Ok{
			Ok: false,
		}
		return &ok, err
	}
	ok := AddService.Ok{
		Ok: true,
	}
	return &ok, nil
}
