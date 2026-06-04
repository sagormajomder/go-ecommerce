package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "user=postgres password=Sagor112 host=localhost port=5432 dbname=ecommerce sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	dbSouce := GetConnectionString()

	dbCon, err := sqlx.Connect("postgres", dbSouce)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil
}
