package main

import (
	"fmt"
	"log"
	"database/sql"
	
	"github.com/go-sql-driver/mysql"
	"github.com/luisfucros/go-api-tutorial/cmd/api"
	"github.com/luisfucros/go-api-tutorial/configs"
	"github.com/luisfucros/go-api-tutorial/db"

)

func main() {
	db, err := db.NewMySQLStorage(
		mysql.Config{
			User:                 configs.Envs.DBUser,
			Passwd:               configs.Envs.DBPassword,
			Addr:                 configs.Envs.DBAddress,
			DBName:               configs.Envs.DBName,
			Net:                  "tcp",
			AllowNativePasswords: true,
			ParseTime:            true,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
	err = server.Run(); if err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}