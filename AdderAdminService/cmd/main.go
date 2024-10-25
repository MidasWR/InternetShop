package main

import (
	"InternetShop/AdderAdminService/db"
	adderAdmin "InternetShop/AdderAdminService/proto"
	"InternetShop/AdderAdminService/server"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
	"google.golang.org/grpc"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := server.Server{
		Db: db.OpenDB(),
	}
	adderAdmin.RegisterAdderAdminServer(s, &srv)
	listener, err := net.Listen("tcp", ":8040")
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
