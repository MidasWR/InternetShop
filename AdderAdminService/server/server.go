package server

import (
	"InternetShop/AdderAdminService/db"
	adderAdmin "InternetShop/AdderAdminService/proto"
	"context"
	"database/sql"
)

type Server struct {
	Db *sql.DB
	adderAdmin.UnimplementedAdderAdminServer
}

func (s *Server) AddAdmin(ctx context.Context, id *adderAdmin.Id) (*adderAdmin.Ok, error) {
	Id := id.Id
	Ok, err := db.AddAdmin(int(Id), s.Db)
	if err != nil {
		return nil, err
	}
	ok := adderAdmin.Ok{}
	ok.Ok = Ok
	return &ok, nil
}
func (s *Server) DeleteAdmin(ctx context.Context, id *adderAdmin.Id) (*adderAdmin.Ok, error) {
	Id := id.Id
	Ok, err := db.DelAdmin(int(Id), s.Db)
	if err != nil {
		return nil, err
	}
	ok := adderAdmin.Ok{}
	ok.Ok = Ok
	return &ok, nil
}
