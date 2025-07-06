package main

import (
	"fmt"

	"github.com/meles-z/golang-graphql/app/infrastructure/db"
	"github.com/meles-z/golang-graphql/configs"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		fmt.Println("Fail to load config file:", err)
		return
	}
	_, err = db.InitDB(cfg.DB)
	if err != nil {
		fmt.Println("Fail to initialize database connection:", err)
		return
	}
}
