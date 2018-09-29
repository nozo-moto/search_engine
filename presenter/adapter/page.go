package adapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/nozo-moto/search_engine/page"
	"github.com/nozo-moto/search_engine/utils"
	"github.com/pkg/errors"
)

type PageAdapter struct {
	Usecase *page.PageUseCase
}

type Page struct {
	ID      int64  `json:"id"`
	URL     string `json:"url"`
	Content string `json:"content"`
}

func (p *Page) domain() *page.Page {
	return &page.Page{
		ID:      p.ID,
		URL:     p.URL,
		Content: p.Content,
	}
}

func NewPage(page *page.Page) *Page {
	return &Page{
		ID:      page.ID,
		URL:     page.URL,
		Content: page.Content,
	}
}

func NewPages(pages []*page.Page) []*Page {
	var result []*Page
	for _, page := range pages {
		result = append(result, NewPage(page))
	}
	return result
}

func NewPageAdapter(pageUsecase *page.PageUseCase) *PageAdapter {
	return &PageAdapter{
		Usecase: pageUsecase,
	}
}

func (p *PageAdapter) GET(w http.ResponseWriter, r *http.Request) error {
	query := r.FormValue("q")
	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		return errors.Wrap(err, "strconv atoi error")
	}
	pages, err := p.Usecase.Search(query, limit)
	if err != nil {
		return fmt.Errorf("usecase search error %v", err)
	}
	return utils.JSON(w, http.StatusOK, NewPages(pages))
}

func (p *PageAdapter) AddTopPage(w http.ResponseWriter, r *http.Request) error {
	var page Page
	if err := json.NewDecoder(r.Body).Decode(&page); err != nil {
		return errors.Wrap(err, "addtoppage json newdecode error")
	}
	_, err := p.Usecase.Add(page.domain())
	if err != nil {
		return err
	}
	return utils.JSON(w, http.StatusCreated, nil)
}
