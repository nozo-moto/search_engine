package adapter

import (
	"net/http"

	"github.com/nozo-moto/search_engine/page"
	"github.com/nozo-moto/search_engine/utils"
)

type PageAdapter struct {
	Usecase *page.PageUseCase
}

type Page struct {
	ID      int64  `json:"id"`
	URL     string `json:"url"`
	Content string `json:"content"`
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
	limit := 10
	pages, err := p.Usecase.Search(query, limit)
	if err != nil {
		return err
	}

	return utils.JSON(w, http.StatusOK, NewPages(pages))
}
