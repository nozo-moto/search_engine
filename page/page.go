package page

import (
	"github.com/pkg/errors"
)

type PageRepository interface {
	Add(page *Page) (*Page, error)
	Search(query string, limit int) ([]*Page, error)
	ContentNullPage() ([]*Page, error)
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
	TITLE   string
}

func NewPage(url, content, title string) *Page {
	return &Page{
		URL:     url,
		Content: content,
		TITLE:   title,
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
	pages, err := p.PageRepo.Search(query, limit)
	if err != nil {
		return nil, errors.Wrap(err, "pagerepo search error")
	}
	return pages, nil
}

func (p *PageUseCase) ContentNullPage() ([]*Page, error) {
	pages, err := p.PageRepo.ContentNullPage()
	if err != nil {
		return nil, err
	}
	return pages, nil
}
