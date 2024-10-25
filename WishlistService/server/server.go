package server

import (
	"InternetShop/WishlistService/db"
	wishlist "InternetShop/WishlistService/proto"
	"context"
	"database/sql"
)

type Server struct {
	Db *sql.DB
	wishlist.UnimplementedWishlistServer
}

func (s *Server) GetWishlist(ctx context.Context, req *wishlist.Empty) (*wishlist.Response, error) {
	items, err := db.GetDataItems(s.Db)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return &wishlist.Response{Items: nil}, nil
	}
	response := &wishlist.Response{
		Items: make([]*wishlist.Item, len(items)),
	}
	for i, item := range items {
		response.Items[i] = &wishlist.Item{
			Id:     item.ID,
			Name:   item.Name,
			Type:   item.Type,
			Amount: item.Amount,
			Cost:   item.Cost,
		}
	}

	return response, nil
}

func (s *Server) AddToWishlist(ctx context.Context, id *wishlist.Id) (*wishlist.Ok, error) {
	ok := wishlist.Ok{
		Ok: false,
	}
	if err := db.AddToWishlist(s.Db, int(id.Id)); err != nil {
		return &ok, err
	}
	ok.Ok = true
	return &ok, nil
}
func (s *Server) DeleteFromWishlist(ctx context.Context, id *wishlist.Id) (*wishlist.Ok, error) {
	ok := wishlist.Ok{
		Ok: false,
	}
	if err := db.DeleteFromWishlist(s.Db, int(id.Id)); err != nil {
		return &ok, err
	}
	ok.Ok = true
	return &ok, nil
}
