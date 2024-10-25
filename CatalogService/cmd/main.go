package main

import (
	"InternetShop/CatalogService/itemdb"
	service "InternetShop/CatalogService/proto"
	"InternetShop/CatalogService/server"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
	"google.golang.org/grpc"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &server.Server{
		Db:  itemdb.OpenDB(),
		Ctx: context.Background(),
	}
	defer srv.Db.Close()
	service.RegisterCatalogServer(s, srv)
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSource(err)
		logrus.Fatalf("failed to listen: %v", err)
	} else {
		logrus.Println("Listening on " + listener.Addr().String())
	}

	if err := s.Serve(listener); err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSource(err)
		logrus.Fatalf("failed to serve: %v", err)
	} else {
		logrus.Println("Serving on " + listener.Addr().String())
	}

}
