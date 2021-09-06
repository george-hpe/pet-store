package server

import (
	"context"
	"fmt"
	"sync"

	"github.com/gk-hpe/pet-store/database"
	"github.com/gk-hpe/pet-store/petstorepb"
	"github.com/google/uuid"
)

// Backend implements the protobuf interface.
type Backend struct {
	petstorepb.UnsafeStoreServiceServer
	mu *sync.RWMutex
	db database.Repository
}

// New initializes a new Backend struct.
func New() *Backend {
	return &Backend{
		mu: &sync.RWMutex{},
		db: database.NewRepository(),
	}
}

// AddProduct adds a new product to the store.
func (s *Backend) AddProduct(ctx context.Context, req *petstorepb.Product) (*petstorepb.AddProductResponse, error) {
	product := database.Product{
		ID:       uuid.New().String(),
		Name:     req.Name,
		Category: req.Category,
		URL:      req.PhotoUrl,
		Status:   req.Status.Enum().String(),
	}

	id, err := s.db.AddProduct(&product)
	if err != nil {
		fmt.Println(err)
	}
	return &petstorepb.AddProductResponse{Id: id}, nil
}

// ListProduct fetches all products from the store.
func (s *Backend) ListProduct(context.Context, *petstorepb.ItemRequest) (*petstorepb.ItemResponse, error) {

	res, err := s.db.ListProduct()
	if err != nil {
		fmt.Println(err)
	}

	var response petstorepb.ItemResponse
	for _, v := range res {
		response.Products = append(response.Products,
			&petstorepb.Product{
				Id:       v.ID,
				Name:     v.Name,
				Category: v.Category,
				PhotoUrl: v.URL,
				Status:   getStatus(v.Status),
			})
	}

	return &response, nil
}

func getStatus(status string) petstorepb.Product_Status {
	switch status {
	case "PENDING":
		return 1
	case "SOLD":
		return 2
	default:
		return 0
	}
}
