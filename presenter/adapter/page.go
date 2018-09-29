package adapter

import (
	"fmt"
	"net/http"

	"github.com/nozo-moto/search_engine/page"
)

type PageAdapter struct {
	Usecase *page.PageUseCase
}

func NewPageAdapter(pageUsecase *page.PageUseCase) *PageAdapter {
	return &PageAdapter{
		Usecase: pageUsecase,
	}
}

func (p *PageAdapter) GET(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("hello")
	return nil
}
