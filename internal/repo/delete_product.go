package repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)


func (r *repo) DeleteProduct(ctx context.Context, product_id int64) error{
	query := psql.Delete("product").
	Where(sq.Eq{"id": product_id}).
	RunWith(r.db)

	if _, err := query.QueryContext(ctx); err != nil{
		return err
	}

	return nil
}