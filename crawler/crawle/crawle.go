package crawle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
	nkf "github.com/nozo-moto/go-nkf"
	"github.com/pkg/errors"
	"github.com/saintfish/chardet"
)

const (
	ToppagePath = "toppage.json"
)

var sawPages []string

type TopPages struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

func LoadTopPage() ([]TopPages, error) {
	// TODO
	// これ、DBから取ってくるような関数にする
	p := []TopPages{}
	bytes, err := ioutil.ReadFile("./" + ToppagePath)
	if err != nil {
		return p, err
	}
	err = json.Unmarshal(bytes, &p)
	if err != nil {
		return p, err
	}
	return p, nil
}

type CrawlePage struct {
	ID       string   `json:"id"`
	URL      string   `json:"url"`
	TITLE    string   `json:"title"`
	TEXT     string   `json:"text"`
	Tolink   []string `json:"tolink"`
	ToBelink []string `json:"to_belink"`
}

func NewPage(url, title string) (*CrawlePage, error) {
	uuid := uuid.New().String()
	return &CrawlePage{
		uuid, url, title, "", nil, nil,
	}, nil
}

func (p *CrawlePage) GetTEXT() error {
	text, err := gettext(p.URL)
	if err != nil {
		return err
	}
	p.TEXT = text
	return nil
}

func (p *CrawlePage) GetLink() error {
	links, err := geturlfrompage(p.URL)
	if err != nil {
		return err
	}
	p.Tolink = links
	return nil
}

func EncodeToUTF8(text string) (string, error) {

	return "", nil
}

func GuessEncoding(text string) (string, error) {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest([]byte(text))
	if err != nil {
		return "", err
	}
	return result.Charset, nil
}

func gettext(url string) (string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}
	// TODO: utf8に全て変換する
	//	text, err := charsetutil.DecodeString(doc.Find("body").Text(), "EUC-JP")
	//	if err != nil {
	//		return "", err
	//	}
	text := doc.Find("body").Text()
	textUTF8, err := nkf.ConvertText(text, "", "UTF8", "")
	if err != nil {
		return "", errors.Wrap(err, "nkf convert text error")
	}
	result := strings.Join(
		strings.Split(
			strings.TrimSpace(
				textUTF8,
			),
			"\n",
		), " ",
	)
	return result, nil
}

func geturlfrompage(url string) ([]string, error) {
	sawPages = append(sawPages, url)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprint(doc))
	}
	var result []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		r := regexp.MustCompile(`web-ext`)
		r2 := regexp.MustCompile(`web-int`)
		r3 := regexp.MustCompile(`u-aizu`)
		r4 := regexp.MustCompile(`html`)
		if (r.MatchString(link) == true || r2.MatchString(link) == true || r3.MatchString(link) == true) && contains(sawPages, link) != true && r4.MatchString(link) == true {
			result = append(
				result, link,
			)
		}
	})
	return result, nil
}
func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
