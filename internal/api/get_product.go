package api

import (
	"context"

	"github.com/gotem2006/vitalmebel-product/internal/model"
	pb "github.com/gotem2006/vitalmebel-product/pkg/product"
)

func (c productApi) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	products, err := c.repo.SelectProduct(ctx)
	if err != nil{
		return nil, err
	}

	return &pb.GetProductResponse{Products: convertToPb(*products)}, nil
}


func convertToPb(products []model.Product) []*pb.Product{
	res := []*pb.Product{}
	for _, product := range products{
		res = append(res, &pb.Product{
			ProductId: product.Id,
			Tittle: product.Tittle,
			Cost: float32(product.Cost),
			Desc: product.Desc,
			Amount: product.Amount,
		})
	}
	return res
}