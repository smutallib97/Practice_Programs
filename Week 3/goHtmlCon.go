/*
 Write a Go program that reads a list of URLs from a file and fetches their HTML contents
 concurrently using Go routines.
 Store the HTML contents in a map with URLs as keys and their respective HTML as values.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func fetchURL(url string, ch chan<- map[string]string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		wg.Done()
		return
	}
	defer resp.Body.Close()

	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		wg.Done()
		return
	}

	data := make(map[string]string)
	data[url] = string(htmlBytes)

	ch <- data
	wg.Done()
}

func main() {
	urls := []string{
		"https://www.example.com",
		"https://www.google.com",
		"https://www.openai.com",
		"https://www.github.com",
	}

	ch := make(chan map[string]string)
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go fetchURL(url, ch, &wg)
	}

	wg.Wait()
	close(ch)

	htmlContents := make(map[string]string)
	for data := range ch {
		for url, html := range data {
			htmlContents[url] = html
		}
	}

	for url, html := range htmlContents {
		fmt.Printf("URL: %s\nHTML Content:\n%s\n\n", url, html)
	}
}
