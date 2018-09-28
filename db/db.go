package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	driverName = "mysql"
)

type Page struct {
	ID      int64  `db:"ID"`
	URL     string `db:"URL"`
	Content string `db:"CONTENT"`
	DB      *sqlx.DB
}

func ConnectToDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Connect(driverName, "")
	if err != nil {
		return db, err
	}
	return db, nil
}

func (p *Page) Insert() (*Page, error) {

	stmt, err := p.DB.Prepare(`INSERT INTO Page (URL, CONTENT) VALUES (?, ?)`)
	defer stmt.Close()
	if err != nil {
		return p, errors.Wrap(err, "page insert error")
	}
	result, err := stmt.Exec(p.URL, p.Content)
	if err != nil {
		return p, errors.Wrap(err, "stmt exec error")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return p, errors.Wrap(err, "last insertid error")
	}
	p.ID = id

	return p, nil
}
