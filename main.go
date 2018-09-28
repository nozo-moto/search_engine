package main

import (
	"fmt"

	"github.com/nozo-moto/search_engine/db"
	"github.com/nozo-moto/search_engine/page"
)

func main() {
	dbx, err := db.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	defer dbx.Close()

	pageusecase := page.NewPageUseCase(
		db.NewPageMySQLAdapter(dbx),
	)
	page := &page.Page{
		URL:     "http://example.com",
		Content: "test",
	}

	page, err = pageusecase.Add(page)
	if err != nil {
		panic(err)
	}
	fmt.Println(page)

}
