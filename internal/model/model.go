package model



type Product struct{
	Id int64 `db:"id"`
	Tittle string `db:"tittle"`
	Cost float64 `db:"cost"`
	Amount int64 `db:"amount"`
	Desc string `db:"description"`
}