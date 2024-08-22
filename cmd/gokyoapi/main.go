package main

import (
	"database/sql"
	"gokyoapi/internal/service"
	"gokyoapi/internal/web"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./gokyo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	GokyoService := service.NewGokyoService(db)
	GokyoHandlers := web.NewGokyoHandler(GokyoService)
	router := http.NewServeMux()
	router.HandleFunc("GET /status", GokyoHandlers.GetStatus)
	router.HandleFunc("GET /gokyo", GokyoHandlers.GetGokyo)
	router.HandleFunc("POST /gokyo", GokyoHandlers.CreateGokyo)
	router.HandleFunc("DELETE /gokyo/{id}", GokyoHandlers.DeleteGokyo)
	http.ListenAndServe(":8080", router)
}
