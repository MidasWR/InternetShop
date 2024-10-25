package main

import (
	"InternetShop/ClientService/config"
	"InternetShop/ClientService/data_users"
	server "InternetShop/ClientService/server"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
)

func main() {
	rest := server.AuthMicroService{
		Config: config.NewConfig(),
		Db:     data_users.NewDB(),
		Router: mux.NewRouter(),
		Logger: logrus.New(),
	}
	if err, er := rest.Run(); err != nil {
		rest.Logger.Println(err)
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	} else if er != nil {
		rest.Logger.Println(er)
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	} else {
		logrus.Println("Server listening on ", rest.Config.Port)
	}
}
