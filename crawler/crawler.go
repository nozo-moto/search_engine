package crawler

import (
	"log"

	"github.com/nozo-moto/search_engine/crawler/crawle"
	"github.com/nozo-moto/search_engine/page"
	"github.com/pkg/errors"
)

type CrawleUseCase struct {
	PageRepo page.PageRepository
}

func NewCrawleUseCase(pageRepo page.PageRepository) *CrawleUseCase {
	return &CrawleUseCase{
		PageRepo: pageRepo,
	}
}

func (c *CrawleUseCase) Crawle() error {
	log.Println("moving cralwer")

	// DBから Contentがnullのデータを取得
	pages, err := c.PageRepo.ContentNullPage()
	if err != nil {
		return errors.Wrap(err, "pagereop contentnullpage error")
	}
	log.Println("null page = ", pages)

	// webpage からデータを取得してくる
	var crawledPage []*crawle.CrawlePage
	for _, page := range pages {
		result, err := run(page.URL, "")
		if err != nil {
			return err
		}
		crawledPage = append(crawledPage, result...)
	}
	log.Println("got pages", crawledPage)

	var dbPages []*page.Page
	for _, p := range crawledPage {
		dbPages = append(dbPages,
			&page.Page{
				URL:     p.URL,
				Content: p.TEXT,
				TITLE:   p.TITLE,
			},
		)
	}

	var result []*page.Page
	for _, p := range dbPages {
		// TODO バルクインサートするべき
		r, err := c.PageRepo.Add(p)
		if err != nil {
			return err
		}
		result = append(result, r)
	}

	log.Println("finished crawler")
	return nil
}

var pages []*crawle.CrawlePage

func run(url, title string) ([]*crawle.CrawlePage, error) {
	page, err := crawle.NewPage(url, title)
	if err != nil {
		return nil, err
	}
	err = page.GetTEXT()
	if err != nil {
		return nil, err
	}
	err = page.GetLink()
	if err != nil {
		return nil, err
	}
	pages = append(pages, page)
	if len(page.Tolink) == 0 {
		return pages, nil
	}
	for _, pageurl := range page.Tolink {
		_, err := run(pageurl, "")
		if err != nil {
			return nil, err
		}
	}
	log.Println("finished run", pages)
	return pages, nil
}
