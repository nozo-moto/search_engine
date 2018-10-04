package crawle

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func Test_geturlfrompage(t *testing.T) {
	// TODO
	// Web-intの大学環境にいないと接続できないあれなのでなんとかする
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test webint",
			args: args{
				url: "https://web-int.u-aizu.ac.jp/official/index.html",
			},
			want: []string{
				"http://web-int.u-aizu.ac.jp/labs/istc/ipc/index.html",
				"http://www.u-aizu.ac.jp/e-current/e-internal.html",
			},
			wantErr: false,
		},
		{
			name: "test is not webint",
			args: args{
				url: "https://www.example.com",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := geturlfrompage(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("geturlfrompage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("geturlfrompage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGuessEncoding(t *testing.T) {
	f, err := os.Open(filepath.Join("testdata", "euc.html"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	testEUCdoc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test euc",
			args: args{
				text: testEUCdoc.Text(),
			},
			want:    "EUC-JP",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GuessEncoding(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("GuessEncoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GuessEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeToUTF8(t *testing.T) {
	f, err := os.Open(filepath.Join("testdata", "euc.html"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	testEUCdoc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	f, err = os.Open(filepath.Join("testdata", "utf8.html"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	utf8doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test euc",
			args: args{
				text: testEUCdoc.Text(),
			},
			want:    utf8doc.Text(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeToUTF8(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeToUTF8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodeToUTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gettext(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{
			name: "example.com",
			args: args{
				url: "https://www.example.com/",
			},
			want:    "",
			want1:   "Example Domain",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1, err := gettext(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("gettext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got1 != tt.want1 {
				t.Errorf("gettext() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
