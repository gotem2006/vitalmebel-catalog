package repo


import(
	"context"

	pb "github.com/gotem2006/vitalmebel-product/pkg/product"
)


func (r *repo) InsertProduct(ctx context.Context, product *pb.Product) error{
	query := psql.Insert("product").
	Columns("tittle", "cost", "amount", "description").
	Values(product.Tittle,product.Cost, product.Amount, product.Desc).
	Suffix("RETURNING id").
	RunWith(r.db)

	if err := query.QueryRowContext(ctx).Scan(&product.ProductId); err != nil{
		return err
	}

	return nil
}