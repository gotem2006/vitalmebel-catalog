package api

import (
	"context"
	"fmt"

	pb "github.com/gotem2006/vitalmebel-product/pkg/product"
)

func (c productApi) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	err := c.repo.InsertProduct(ctx, req.Product)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return &pb.CreateProductResponse{
		Res: req.Product,
	}, nil
}
