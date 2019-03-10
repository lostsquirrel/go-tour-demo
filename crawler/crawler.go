package main

import (
    "fmt"
	"net/http"
    "sync"
			"github.com/PuerkitoBio/goquery"
	"log"
		"io"
	"os"
	"bytes"
		)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type MFetcher struct {
}

type UrlStore struct {
	m map[string]bool
	mux sync.Mutex
}

func NewUrlStore() *UrlStore {
	us := &UrlStore{}
	us.m = make(map[string]bool)
	return us
}

func (us *UrlStore) exists(url string) bool {
	us.mux.Lock()
	_, ok := us.m[url]
	if !ok {
		us.m[url] = true
	}
	defer us.mux.Unlock()
	return ok
}

func (f MFetcher) Fetch(url string) (body string, urls []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("fail to fetch %v", url)
		return
	}
	defer resp.Body.Close()
	//content, err := ioutil.ReadAll(resp.Body)
	//body = string(content)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var buffer bytes.Buffer
	doc.Find(".commentlist").Each(func(i int, ol *goquery.Selection) {
		// For each item found, get the band and title

		ol.Children().Each(func(i int, li *goquery.Selection) {
			//fmt.Println(i)
			li.Find(".img-hash").Each(func(i int, image *goquery.Selection) {
				//fmt.Println(i)
				buffer.WriteString(image.Text())
				fmt.Println(image.Text())
				buffer.WriteString(",")
			})


		})

		//fmt.Println(s.Find(".tucao-like-container").ChildrenFiltered("span").First().Text())
		//fmt.Println(s.Find(".tucao-unlike-container").ChildrenFiltered("span").First().Text())
		//fmt.Println("")
		//band := s.Find("a").Text()
		//title := s.Find("i").Text()
		//fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
	pp,_ := doc.Find(".previous-comment-page").First().Attr("href")
	//fmt.Println(pp)
	//fmt.Println(body)
	return buffer.String(), []string{fmt.Sprintf("http:%s", pp)}, err
}
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func crawl(url string, depth int, fetcher Fetcher, us *UrlStore, out chan string)  {
	if depth <= 0 || us.exists(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	out <- body
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		crawl(u, depth - 1, fetcher, us, out)
	}

}
func outHandler(out chan string, path string)  {
	for body := range out {
		fmt.Println(body)
	}
}

func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	path := "G:\\downloads\\beauty\\xxoo"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// path/to/whatever does not exist
		fh, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer closeFile(fh)
	}

	us := NewUrlStore()
	out := make(chan string, 100)
	crawl(url, depth, fetcher, us, out)
	close(out)
	//fmt.Printf("found: %s %q\n", url, body)
	//go outHandler(out, path)
}

func main() {
	Crawl("http://jandan.net/ooxx", 200, MFetcher{})
}

func closeFile(fh io.Closer) {
	if err := fh.Close(); err != nil {
		panic(err)
	}
}