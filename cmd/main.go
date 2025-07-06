package main

import (
	"fmt"
	"net/http"

	"github.com/meles-z/soap-test/config"
	dbutils "github.com/meles-z/soap-test/internal/db_utils"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	_, err = dbutils.InitDb(cfg.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started on 8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi there"))
	})
	http.ListenAndServe(":8080", nil)
}
