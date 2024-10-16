package repo

import (
	"context"

	"github.com/gotem2006/vitalmebel-product/internal/model"
)



func (r *repo) SelectProduct(ctx context.Context) (*[]model.Product, error){
	query := psql.Select("id", "tittle", "cost", "amount", "description").
	From("product").
	RunWith(r.db)

	sql, _, err := query.ToSql()
	if err != nil{
		return nil, err
	}
	products := &[]model.Product{}
	if err := r.db.SelectContext(ctx, products, sql); err != nil{
		return nil, err
	}


	return products, nil
}