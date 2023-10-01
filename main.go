package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yuki0418/go-gin-tutorial/db"
	"github.com/yuki0418/go-gin-tutorial/server"
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	dbManager := db.DatabaseManager{}
	dbManager.Init()

	server.Init()
}