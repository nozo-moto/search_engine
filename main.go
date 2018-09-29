package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nozo-moto/search_engine/db"
	"github.com/nozo-moto/search_engine/page"
	"github.com/nozo-moto/search_engine/presenter/adapter"
)

type Server struct {
	router *mux.Route
}

func main() {
	router := mux.NewRouter()
	dbx, err := db.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	defer dbx.Close()

	pageAdapter := adapter.NewPageAdapter(
		page.NewPageUseCase(
			db.NewPageMySQLAdapter(dbx),
		),
	)

	router.Handle("/api/v1/page", handler(pageAdapter.GET)).Methods("GET")
	http.ListenAndServe(":8080", router)
}
