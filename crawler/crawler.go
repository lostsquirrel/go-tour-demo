package main

import (
    "fmt"
	"net/http"
		"sync"
	"golang.org/x/net/html"
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
	z := html.NewTokenizer(resp.Body)
	for  {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		} else {
			fmt.Println(tt)
		}
	}
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
