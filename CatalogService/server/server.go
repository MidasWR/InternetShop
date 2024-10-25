package server

import (
	"InternetShop/CatalogService/itemdb"
	service "InternetShop/CatalogService/proto"
	wishlist "InternetShop/WishlistService/proto"
	"context"
	"database/sql"
	"google.golang.org/grpc"
	"log"
	"time"
)

type Server struct {
	Db  *sql.DB
	Ctx context.Context
	service.UnimplementedCatalogServer
}

func (s *Server) mustEmbedUnimplementedCatalogServer() {
	panic("implement me")
}

func (s *Server) AddItemToWishlist(ctx context.Context, id *service.IdItem) (*service.Response2, error) {
	conn, err := grpc.Dial(":8050", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(60*time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := wishlist.NewWishlistClient(conn)
	Id := wishlist.Id{
		Id: id.Id,
	}
	ok, err := c.AddToWishlist(context.Background(), &Id)
	if err != nil {
		return &service.Response2{Ok: false}, err
	}
	if !ok.Ok {
		return &service.Response2{Ok: false}, nil
	}
	return &service.Response2{Ok: true}, nil
}

func (s *Server) GetItemsFromDB(ctx context.Context, req *service.Empty) (*service.Response1, error) {
	items, err := itemdb.GetDataItems(s.Db)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return &service.Response1{Items: nil, Success: false}, nil
	}

	response := &service.Response1{
		Items:   make([]*service.Item, len(items)),
		Success: true,
	}
	for i, item := range items {
		response.Items[i] = &service.Item{
			Id:     item.ID,
			Name:   item.Name,
			Type:   item.Type,
			Amount: item.Amount,
			Cost:   item.Cost,
		}
	}

	return response, nil
}
