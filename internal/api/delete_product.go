package api

import (
	"context"

	pb "github.com/gotem2006/vitalmebel-product/pkg/product"
)

func (c productApi) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	if err := c.repo.DeleteProduct(ctx, req.ProductId); err != nil {
		return nil, err
	}
	return &pb.DeleteProductResponse{}, nil
}
