package main

import (
	"fmt"

	"github.com/nozo-moto/search_engine/db"
)

func main() {
	dbx, err := db.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	page := &db.Page{
		URL:     "http://example.com",
		Content: "test",
		DB:      dbx,
	}

	page, err = page.Insert()
	if err != nil {
		panic(err)
	}
	fmt.Println(page)

}
