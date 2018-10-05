package db

import (
	"testing"

	"github.com/nozo-moto/search_engine/page"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		Name  string
		Pages []*page.Page
		Word  string
		Limit int
	}{
		{
			Name: "japanese success case",
			Pages: []*page.Page{
				{
					ID:      0,
					URL:     "example.com/test",
					Content: "日本語が動くかのテストだよ",
					Title:   "テスト",
				},
			},
			Word:  "日本語",
			Limit: 1,
		},
	}

	for i, tc := range cases {
		dbx, err := ConnectToDatabase()
		defer dbx.Close()
		if err != nil {
			t.Logf("Failed MySQL open: %v", err)
			t.Fail()
		}
		adapter := NewPageMySQLAdapter(dbx)
		p, err := adapter.Add(tc.Pages[i])
		if err != nil {
			t.Errorf("failed add: %v", err)
		}
		t.Logf("success add: %v", p)
		pages, err := adapter.Search(tc.Word, tc.Limit)
		if err != nil {
			t.Errorf("failed search: %v", err)
		}

		for j := range pages {
			if tc.Pages[j].URL != pages[j].URL {
				t.Errorf("test failed! expected %v, but got %v", tc.Pages[j].URL, pages[j].URL)
			}
			t.Logf("pages content: %v", pages[j].Content)
		}
	}
}
