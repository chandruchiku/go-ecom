package main

import (
	"database/sql"
	"log"

	"github.com/chandruchiku/go-ecom/cmd/api"
	"github.com/chandruchiku/go-ecom/config"
	"github.com/chandruchiku/go-ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBConfig.User,
		Passwd:               config.Envs.DBConfig.Passwd,
		Addr:                 config.Envs.DBConfig.Addr,
		DBName:               config.Envs.DBConfig.Name,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}

	initStorage(db)

	server := api.New(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("error pinging db: %v", err)
	}
	log.Printf("Connected to db")
}
