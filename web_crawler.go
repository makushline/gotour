package main

import (
    "fmt"
    "sync"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, seen map[string]bool, ch chan string) {
    defer close(ch)
    if depth <= 0 || seen[url] {
        return
    }
    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        ch <- err.Error()
        return
    }
    seen[url] = true
    ch <- fmt.Sprintf("found: %s %q", url, body)
    results := make([]chan string, len(urls))
    for i, u := range urls {
        results[i] = make(chan string)
        go Crawl(u, depth-1, fetcher, seen, results[i])
    }
    for i := range results {
        for s := range results[i] {
            ch <- s
        }
    }
    return
}

type synchronizedUrlCache struct {
    cache map[string]*fakeResult
    mux sync.Mutex
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f synchronizedUrlCache) Fetch(url string) (string, []string, error) {
    f.mux.Lock()
    defer f.mux.Unlock()
    if res, ok := f.cache[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
    
}

var cache = map[string]*fakeResult{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}
// fetcher is a populated fakeFetcher.
var fetcher = synchronizedUrlCache{cache: cache}

func main() {
    ch := make(chan string)
    go Crawl("http://golang.org/", 4, fetcher, make(map[string]bool), ch)
    for s := range ch {
        fmt.Println(s)
    }
}