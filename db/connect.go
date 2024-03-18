package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "admin"
    password = "admin"
    dbname   = "testapi"
)


var DB *pgxpool.Pool

func Connect() error{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	conn, err:= pgxpool.New(context.Background(), psqlInfo)
	if err != nil {
		conn.Close()
		return err
	}
	defer conn.Close()

	err= conn.Ping(context.Background())
	if err != nil {
		fmt.Println("error to connect DB")
		conn.Close()
		return err
	}
	DB = conn

	return nil
}