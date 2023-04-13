package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"

	_ "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag"

	"kmf/internal/app"

	_ "github.com/denisenkom/go-mssqldb"
)

type Config struct {
	Addr       string `json:"addr"`
	ConnString string `json:"conn_string"`
}

func main() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	db := OpenDB(config.ConnString)

	app.Run(config.Addr, db)
}

func OpenDB(conn string) *sql.DB {
	db, err := sql.Open("mssql", conn)
	if err != nil {
		log.Fatal("Eror open db connection: ", err)
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Eror ping db connection: ", err)
	}

	return db
}
