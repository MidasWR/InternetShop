package main

import (
	"InternetShop/WishlistService/db"
	wishlist "InternetShop/WishlistService/proto"
	"InternetShop/WishlistService/server"
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
	wishlist.RegisterWishlistServer(s, &srv)
	listener, err := net.Listen("tcp", ":8050")
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
