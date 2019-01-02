package page

import (
	"encoding/json"

	redisCache "github.com/nozo-moto/search_engine/redis"
	"github.com/pkg/errors"
)

type PageRepository interface {
	Add(page *Page) (*Page, error)
	Search(query string, limit int) ([]*Page, error)
	ContentNullPage() ([]*Page, error)
	DeleteNullPage() error
	AddTopPage(url string) error
}

type PageUseCase struct {
	PageRepo     PageRepository
	RedisAdapter *redisCache.RedisAdapter
}

func NewPageUseCase(pr PageRepository, redis *redisCache.RedisAdapter) *PageUseCase {
	return &PageUseCase{
		PageRepo:     pr,
		RedisAdapter: redis,
	}
}

type Page struct {
	ID      int64
	URL     string
	Content string
	TITLE   string
	Desc    string
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
	result := []*Page{}
	resultByte, err := p.RedisAdapter.GetSearch(query, string(limit))
	if err != nil {
		result, err = p.PageRepo.Search(query, limit)
		if err != nil {
			return nil, errors.Wrap(err, "pagerepo search error")
		}
		go func() {
			resultByte, err = json.Marshal(result)
			if err != nil {
				panic(err)
			}
			err = p.RedisAdapter.SetSearch(query, string(limit), resultByte)
			if err != nil {
				panic(err)
			}
		}()
		return result, nil
	}
	err = json.Unmarshal(resultByte, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *PageUseCase) ContentNullPage() ([]*Page, error) {
	pages, err := p.PageRepo.ContentNullPage()
	if err != nil {
		return nil, err
	}
	return pages, nil
}
func (p *PageUseCase) DeleteNullPage() error {
	return p.PageRepo.DeleteNullPage()
}

func (p *PageUseCase) AddTopPage(url string) error {
	return p.PageRepo.AddTopPage(url)
}
