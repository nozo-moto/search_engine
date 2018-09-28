package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nozo-moto/search_engine/page"
	"github.com/pkg/errors"
)

const (
	driverName = "mysql"
)

type PageMySQLAdapter struct {
	ID      int64  `db:"ID"`
	URL     string `db:"URL"`
	Content string `db:"CONTENT"`
	DB      *sqlx.DB
}

func NewPageMySQLAdapter(db *sqlx.DB) *PageMySQLAdapter {
	return &PageMySQLAdapter{
		DB: db,
	}
}

func (p *PageMySQLAdapter) domain() *page.Page {
	return &page.Page{
		ID:      p.ID,
		URL:     p.URL,
		Content: p.Content,
	}
}

func ConnectToDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Connect(driverName, "root:password@tcp(0.0.0.0:13306)/search_engine")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (p *PageMySQLAdapter) Add(page *page.Page) (*page.Page, error) {
	stmt, err := p.DB.Prepare(`INSERT INTO Page (URL, CONTENT) VALUES (?, ?)`)
	defer stmt.Close()
	if err != nil {
		return nil, errors.Wrap(err, "page insert error")
	}
	result, err := stmt.Exec(page.URL, page.Content)
	if err != nil {
		return nil, errors.Wrap(err, "stmt exec error")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "last insertid error")
	}
	page.ID = id

	return page, nil
}
