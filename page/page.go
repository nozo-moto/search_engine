package page

import (
	"github.com/pkg/errors"
)

type PageRepository interface {
	Add(page *Page) (*Page, error)
	Search(query string, limit int) ([]*Page, error)
}

type PageUseCase struct {
	PageRepo PageRepository
}

func NewPageUseCase(pr PageRepository) *PageUseCase {
	return &PageUseCase{
		PageRepo: pr,
	}
}

type Page struct {
	ID      int64
	URL     string
	Content string
}

func NewPage(url, content string) *Page {
	return &Page{
		URL:     url,
		Content: content,
	}
}

func (p *PageUseCase) Add(page *Page) (*Page, error) {
	page, err := p.PageRepo.Add(page)
	if err != nil {
		return nil, errors.Wrap(err, "usecase add error")
	}

	return page, nil
}

func (p *PageUseCase) Search(query string, limit int) ([]*Page, error) {
	// TODO impl
	return nil, nil
}
