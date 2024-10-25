package server

import (
	adderAdmin "InternetShop/AdderAdminService/proto"
	AddService "InternetShop/AdderCatalogService/proto"
	service "InternetShop/CatalogService/proto"
	"InternetShop/ClientService/config"
	"InternetShop/ClientService/data_users"
	wishlist "InternetShop/WishlistService/proto"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
	"time"
)

type AuthMicroService struct {
	Config *config.Config
	Db     *sql.DB
	Router *mux.Router
	Logger *logrus.Logger
	Token  string
}

func (a *AuthMicroService) Run() (error, error) {
	a.Router.HandleFunc("/registration", RegPage(a)).Methods("POST")
	a.Router.HandleFunc("/login", LoginPage(a)).Methods("POST")
	a.Router.HandleFunc("/catalog", CatalogPage(a)).Methods("GET", "POST")
	a.Router.HandleFunc("/catalog/adder", CatalogAdderPage(a)).Methods("POST")
	a.Router.HandleFunc("/admin/adder", AdderAdminPage(a)).Methods("POST")
	a.Router.HandleFunc("/catalog/wishlist", WishlistPage(a)).Methods("GET")
	addr := fmt.Sprintf("%s:%s", a.Config.Host, a.Config.Port)
	return http.ListenAndServe(addr, a.Router), a.Db.Close()
}
func RegPage(a *AuthMicroService) http.HandlerFunc {
	type RegUser struct {
		Username string `json:"username"`
		Login    string `json:"login"`
		Password string `json:"password"`
		Role     string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ru := RegUser{}
		if err := json.NewDecoder(r.Body).Decode(&ru); err != nil {
			tracerr.PrintSourceColor(err)
			a.Logger.Println(err)
		}
		ru.Role = "consumer"
		if ok, err := data_users.CheckUsernameExists(ru.Username, a.Db); err != nil {
			tracerr.PrintSourceColor(err)
			a.Logger.Fatal(err)
		} else if ok {
			err := errors.New("this username already exists")
			a.Logger.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			tracerr.PrintSourceColor(err)
		}
		if ok, err := data_users.CheckLoginExists(ru.Login, a.Db); err != nil {
			tracerr.PrintSourceColor(err)
			a.Logger.Fatal(err)
		} else if ok {
			err := errors.New("this login already exists")
			a.Logger.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			tracerr.PrintSourceColor(err)
		}
		HashPassword, err := Crypto(ru.Password)
		if err != nil {
			a.Logger.Println(err)
			tracerr.PrintSourceColor(err)
		} else {
			if err := data_users.AddNewMember(ru.Username, ru.Login, HashPassword, ru.Role, a.Db); err != nil {
				a.Logger.Println(err)
				tracerr.PrintSourceColor(err)
			}
		}

	}
}
func LoginPage(a *AuthMicroService) http.HandlerFunc {
	type LoginUser struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		Role     string
		ID       int
	}
	return func(w http.ResponseWriter, r *http.Request) {
		lu := LoginUser{}
		if err := json.NewDecoder(r.Body).Decode(&lu); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = tracerr.Wrap(err)
			tracerr.PrintSourceColor(err)
			a.Logger.Println(err)
		}
		if ok, err := data_users.CheckLoginExists(lu.Login, a.Db); err != nil {
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				err = tracerr.Wrap(err)
				tracerr.PrintSourceColor(err)
				a.Logger.Fatal(err)
			} else if !ok {
				err := errors.New("this login does not exist")
				a.Logger.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				err = tracerr.Wrap(err)
				tracerr.PrintSourceColor(err)
			}
		}
		if HashedPassword, err := data_users.GetHashPasswordByLogin(a.Db, lu.Login); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			a.Logger.Println(err)
			err = tracerr.Wrap(err)
			tracerr.PrintSourceColor(err)
		} else {
			if err := CheckPasswordHash(lu.Password, HashedPassword); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				a.Logger.Println(err)
				err = tracerr.Wrap(err)
				tracerr.PrintSourceColor(err)
			} else {
				w.WriteHeader(http.StatusOK)
				lu.Role, err = data_users.GetRoleByLogin(a.Db, lu.Login)
				if err != nil {
					a.Logger.Println(err)
					err = tracerr.Wrap(err)
					tracerr.PrintSourceColor(err)
				}
				lu.ID, err = data_users.GetIDByLogin(a.Db, lu.Login)
				if err != nil {
					a.Logger.Println(err)
					err = tracerr.Wrap(err)
					tracerr.PrintSourceColor(err)
				} else if lu.ID < 1 {
					err := errors.New("user not found")
					a.Logger.Println(err)
					err = tracerr.Wrap(err)
					tracerr.PrintSourceColor(err)
				}
				if a.Token, err = GenerateJWT(lu.Role, lu.ID); err != nil {
					a.Logger.Println(err)
					err = tracerr.Wrap(err)
					tracerr.PrintSourceColor(err)
				}
			}
		}

	}
}
func CatalogPage(a *AuthMicroService) http.HandlerFunc {
	type Claims struct {
		Role string `json:"role"`
		ID   int    `json:"sub"`
		jwt.RegisteredClaims
	}
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := ParseJWT(a.Token, w)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = tracerr.Wrap(err)
			a.Logger.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprintln(w, claims); err != nil {
			err = tracerr.Wrap(err)
			tracerr.PrintSourceColor(err)
		}
		conn, err := grpc.Dial(":8090", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(60*time.Second))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := service.NewCatalogClient(conn)
		action := GetAction(r)
		switch action {
		case "catalog":
			items, err := c.GetItemsFromDB(context.Background(), &service.Empty{})
			if err != nil {
				err = tracerr.Wrap(err)
				tracerr.PrintSourceColor(err)
			}
			for _, item := range items.Items {
				fmt.Fprintln(w, item)
			}
		case "addToWishlist":
			ids := r.URL.Query().Get("id")
			id := service.IdItem{}
			Id, err := strconv.Atoi(ids)
			if err != nil {
				err = tracerr.Wrap(err)
				tracerr.PrintSourceColor(err)
			}
			id.Id = int32(Id)
			ok, err := c.AddItemToWishlist(context.Background(), &id)
			if err != nil {
				err = tracerr.Wrap(err)
				tracerr.PrintSourceColor(err)
			}
			if !ok.Ok {
				w.WriteHeader(http.StatusNotImplemented)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		default:
			w.WriteHeader(404)
		}

	}
}
func CatalogAdderPage(a *AuthMicroService) http.HandlerFunc {
	type Claims struct {
		Role string `json:"role"`
		ID   int    `json:"sub"`
		jwt.RegisteredClaims
	}
	in := AddService.Item{}
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := ParseJWT(a.Token, w)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = tracerr.Wrap(err)
			a.Logger.Println(err)
		}
		if claims.Role != "admin" {
			w.WriteHeader(http.StatusForbidden)
			a.Logger.Println("Not enough rights to log in")
		}
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprintln(w, claims); err != nil {
			err = tracerr.Wrap(err)
			tracerr.PrintSourceColor(err)
		}
		conn, err := grpc.Dial(":8060", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(60*time.Second))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			err = tracerr.Wrap(err)
			tracerr.PrintSourceColor(err)
		}
		c := AddService.NewAdderCatalogClient(conn)
		ok := &AddService.Ok{}
		if ok, err = c.AddItemToCatalog(context.Background(), &in); err != nil {
			err = tracerr.Wrap(err)
			tracerr.PrintSourceColor(err)
		}
		if !ok.Ok {
			w.WriteHeader(http.StatusBadGateway)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
func AdderAdminPage(a *AuthMicroService) http.HandlerFunc {
	type Claims struct {
		Role string `json:"role"`
		ID   int    `json:"sub"`
		jwt.RegisteredClaims
	}
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := ParseJWT(a.Token, w)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = tracerr.Wrap(err)
			a.Logger.Println(err)
		}
		if claims.Role != "admin" {
			w.WriteHeader(http.StatusForbidden)
			a.Logger.Println("Not enough rights to log in")
		}
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprintln(w, claims); err != nil {
			err = tracerr.Wrap(err)
			tracerr.PrintSourceColor(err)
		}
		conn, err := grpc.Dial(":8050", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(60*time.Second))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := adderAdmin.NewAdderAdminClient(conn)
		switch GetAction(r) {
		case "add":
			ids := r.URL.Query().Get("id")
			id := adderAdmin.Id{}
			Id, err := strconv.Atoi(ids)
			if err != nil {
				err = tracerr.Wrap(err)
				tracerr.PrintSourceColor(err)
			}
			id.Id = int32(Id)
			ok, err := c.AddAdmin(context.Background(), &id)
			if err != nil {
				err = tracerr.Wrap(err)
				tracerr.PrintSourceColor(err)
			}
			if !ok.Ok {
				w.WriteHeader(http.StatusNotImplemented)
			}
		case "delete":
			ids := r.URL.Query().Get("id")
			id := adderAdmin.Id{}
			Id, err := strconv.Atoi(ids)
			if err != nil {
				err = tracerr.Wrap(err)
				tracerr.PrintSourceColor(err)
			}
			id.Id = int32(Id)
			ok, err := c.DeleteAdmin(context.Background(), &id)
			if err != nil {
				err = tracerr.Wrap(err)
				tracerr.PrintSourceColor(err)
			}
			if !ok.Ok {
				w.WriteHeader(http.StatusNotImplemented)
			}
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}

	}
}
func WishlistPage(a *AuthMicroService) http.HandlerFunc {
	type Claims struct {
		Role string `json:"role"`
		ID   int    `json:"sub"`
		jwt.RegisteredClaims
	}
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := ParseJWT(a.Token, w)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = tracerr.Wrap(err)
			a.Logger.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprintln(w, claims); err != nil {
			err = tracerr.Wrap(err)
			tracerr.PrintSourceColor(err)
		}
		conn, err := grpc.Dial(":8050", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(60*time.Second))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := wishlist.NewWishlistClient(conn)
		id := wishlist.Id{}
		switch GetAction(r) {
		case "delete":
			ids := r.URL.Query().Get("id")
			Id, err := strconv.Atoi(ids)
			if err != nil {
				err = tracerr.Wrap(err)
				tracerr.PrintSourceColor(err)
				w.WriteHeader(http.StatusBadGateway)
			}
			id.Id = int32(Id)
			ok, err := c.DeleteFromWishlist(context.Background(), &id)
			if err != nil {
				err = tracerr.Wrap(err)
				tracerr.PrintSourceColor(err)
				w.WriteHeader(http.StatusNotImplemented)
			}
			if !ok.Ok {
				w.WriteHeader(http.StatusNotImplemented)
			}
		case "get":
			items, err := c.GetWishlist(context.Background(), nil)
			if err != nil {
				err = tracerr.Wrap(err)
				tracerr.PrintSourceColor(err)
			}
			for _, item := range items.Items {
				fmt.Fprintln(w, item)
			}
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}
