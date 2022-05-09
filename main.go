package main

import (
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"
	"database/sql"
    _ "github.com/lib/pq"
)


func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configuration:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)


	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
	
}

