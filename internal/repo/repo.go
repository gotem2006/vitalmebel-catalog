package repo


import (

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	// "github.com/rs/zerolog/log"
	

)


type repo struct{
	db *sqlx.DB
}


var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)


func NewRepo(db *sqlx.DB) *repo{
	return &repo{
		db: db,
	}
}








