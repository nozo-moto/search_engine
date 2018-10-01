package crawler

import (
	"fmt"

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
	// TODO

	// DBから Contentがnullのデータを取得
	pages, err := c.PageRepo.ContentNullPage()
	if err != nil {
		return errors.Wrap(err, "pagereop contentnullpage error")
	}
	fmt.Println(pages)

	// HTMLを取得

	// リンク一覧を取得

	// HTMLからテキストを取得

	return nil
}
