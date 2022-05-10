package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"
)

// @title Bank Proj
// @version 1.0
// @description Swagger API for Golang Project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email example@email.com

// @license.name MIT
// @license.url https://github.com/VasiliKavaleu/Go-bank-proj-example

// @BasePath /
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
