package main

import (
	"database/sql"
	"log"
	"github.com/justIGreK/simple_bank_course/db/api"
	db "github.com/justIGreK/simple_bank_course/db/sqlc"
	"github.com/justIGreK/simple_bank_course/db/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil{
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start servers:", err)
	}
}
