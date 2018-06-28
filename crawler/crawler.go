package main

import (
    "fmt"
	"net/http"
    "sync"
			"github.com/PuerkitoBio/goquery"
	"log"
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
	doc.Find(".commentlist li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title

		fmt.Println(s.Find(".img-hash").Text())
		fmt.Println(s.Find(".tucao-like-container").ChildrenFiltered("span").First().Text())
		fmt.Println(s.Find(".tucao-unlike-container").ChildrenFiltered("span").First().Text())
		//fmt.Println("")
		//band := s.Find("a").Text()
		//title := s.Find("i").Text()
		//fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})

	//fmt.Println(body)
	return "", []string{}, err
}
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func crawl(url string, depth int, fetcher Fetcher, us *UrlStore)  {
	if depth <= 0 || us.exists(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go crawl(u, depth-1, fetcher, us)
	}
}

func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	us := NewUrlStore()
	crawl(url, depth, fetcher, us)
	//fmt.Printf("found: %s %q\n", url, body)
	
	return
}

func main() {
	Crawl("http://jandan.net/ooxx", 2, MFetcher{})
}
