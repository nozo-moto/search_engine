package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nozo-moto/search_engine/crawler"
	"github.com/nozo-moto/search_engine/db"
	handle "github.com/nozo-moto/search_engine/handler"
	"github.com/nozo-moto/search_engine/page"
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

	mysqlAdapter := db.NewPageMySQLAdapter(dbx)
	pageAdapter := handle.NewPageAdapter(
		page.NewPageUseCase(
			mysqlAdapter,
		),
		crawler.NewCrawleUseCase(
			mysqlAdapter,
		),
	)

	// query parameterで q=検索したい文字列 limit=検索数
	router.Handle("/api/v1/page", handler(pageAdapter.GET)).Methods("GET")

	// toppageを追加するuPI {url: "hwertyui.com"}
	router.Handle("/api/v1/toppage", handler(pageAdapter.AddTopPage)).Methods("POST")

	// Crawler を動かすendpoint
	router.Handle("/api/v1/crawler", handler(pageAdapter.MoveCrawler)).Methods("GET")

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("front/dist/"))))
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("front/dist/css"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("front/dist/js"))))

	http.ListenAndServe(":8080", router)
}
