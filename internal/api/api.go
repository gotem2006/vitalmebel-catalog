package api

import (
	"context"

	"github.com/gotem2006/vitalmebel-product/internal/model"
	pb "github.com/gotem2006/vitalmebel-product/pkg/product"
)

type Repo interface {
	SelectProduct(context.Context) (*[]model.Product, error)
	InsertProduct(context.Context, *pb.Product) error
	DeleteProduct(ctx context.Context, product_id int64) error
	PatchProduct(context.Context, *pb.Product) error
}

type productApi struct {
	pb.UnimplementedProductServiceServer
	repo Repo
}

func NewProductApi(repo Repo) pb.ProductServiceServer {
	return &productApi{
		repo: repo,
	}
}
