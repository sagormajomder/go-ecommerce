package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.Config) string {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s ", cnf.DB.User, cnf.DB.Password, cnf.DB.Host, cnf.DB.Port, cnf.DB.Name)

	if !cnf.DB.EnableSSLMode {
		return connString + " sslmode=disable"
	}

	return connString
}

func NewConnection(cnf *config.Config) (*sqlx.DB, error) {
	dbSouce := GetConnectionString(cnf)

	dbCon, err := sqlx.Connect("postgres", dbSouce)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil
}
