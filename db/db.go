package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nozo-moto/search_engine/page"
	"github.com/pkg/errors"
)

const (
	driverName = "mysql"
)

type PageMySQLAdapter struct {
	ID      int64          `db:"ID"`
	Title   string         `db:"TITLE"`
	URL     string         `db:"URL"`
	Content sql.NullString `db:"CONTENT"`
	TITLE   string         `db:"TITLE"`
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
		Content: p.Content.String,
		TITLE:   p.TITLE,
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
	stmt, err := p.DB.Prepare(`INSERT INTO Page (URL, CONTENT, TITLE) VALUES (?, ?, ?)`)
	defer stmt.Close()
	if err != nil {
		return nil, errors.Wrap(err, "page insert error")
	}

	result, err := stmt.Exec(page.URL, page.Content, page.TITLE)
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

func (p *PageMySQLAdapter) AddTopPage(url string) error {
	stmt, err := p.DB.Prepare(`INSERT INTO Page(URL, TITLE) VALUES(?, "")`)
	if err != nil {
		return errors.Wrap(err, "page insert error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(url)
	if err != nil {
		return err
	}
	return nil
}

func (p *PageMySQLAdapter) Search(query string, limit int) ([]*page.Page, error) {
	var pages []*PageMySQLAdapter
	err := p.DB.Select(&pages, `SELECT * FROM Page WHERE MATCH ( CONTENT ) AGAINST (? IN NATURAL LANGUAGE MODE) LIMIT ?;`, query, limit)
	if err != nil {
		return nil, errors.Wrap(err, "page search error")
	}

	var result []*page.Page
	for _, page := range pages {
		result = append(result, page.domain())
	}

	return result, nil
}

func (p *PageMySQLAdapter) ContentNullPage() ([]*page.Page, error) {
	var pages []*PageMySQLAdapter
	err := p.DB.Select(&pages, `SELECT * FROM Page WHERE CONTENT IS NULL`)
	if err != nil {
		return nil, errors.Wrap(err, "contentnull page error")
	}
	var result []*page.Page
	for _, page := range pages {
		result = append(result, page.domain())
	}
	return result, nil
}

func (p *PageMySQLAdapter) DeleteNullPage() error {
	_, err := p.DB.Exec(`DELETE FROM Page WHERE CONTENT IS NULL`)
	if err != nil {
		return err
	}
	return nil
}
