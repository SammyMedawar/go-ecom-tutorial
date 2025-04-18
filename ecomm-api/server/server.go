package server

import (
	"context"

	"github.com/medawarsammy/go-ecom-tutorial/ecomm-api/storer"
)

type Server struct {
	storer *storer.MySQLStorer
}

func NewServer(storer *storer.MySQLStorer) *Server {
	return &Server{
		storer: storer,
	}
}

func (s *Server) CreateProduct(ctx context.Context, p *storer.Product) (*storer.Product, error) {
	return s.storer.CreateProduct(ctx, p)
}

func (s *Server) GetProduct(ctx context.Context, id int64) (*storer.Product, error) {
	return s.storer.GetProduct(ctx, id)
}

func (s *Server) UpdateProduct(ctx context.Context, p *storer.Product) (*storer.Product, error) {
	return s.storer.UpdateProduct(ctx, p)
}

func (s *Server) DeleteProduct(ctx context.Context, id int64) error {
	return s.storer.DeleteProduct(ctx, id)
}

func (s *Server) ListProducts(ctx context.Context) ([]storer.Product, error) {
	return s.storer.ListProducts(ctx)
}
