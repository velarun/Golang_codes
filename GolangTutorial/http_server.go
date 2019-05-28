package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

type countHandler struct {
	mu sync.Mutex
	n  int
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.google.com", 301)
}

func returnGetArticles(w http.ResponseWriter, r *http.Request) {
	url := "http://country.io/capital.json"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Endpoint Hit: returnGetArticles")
	responseString := string(responseData)
	fmt.Fprintf(w, responseString)
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	url := "https://httpbin.org/post"

	data := []byte(`{"hello": "world"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", "httpbin.org")

	cookie := http.Cookie{Name: "cookie_name", Value: "cookie_value"}
	req.AddCookie(&cookie)

	client := &http.Client{Timeout: time.Second * 10}

	fmt.Println(req.Cookies())
	fmt.Println(req.Header)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	fmt.Fprintf(w, "%s\n", body)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/data", returnAllArticles)
	http.HandleFunc("/get", returnGetArticles)
	http.HandleFunc("/redirect", redirect)
	http.Handle("/get/count", new(countHandler))
	mux := http.NewServeMux()
	mux.HandleFunc("/send", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Trailer", "AtEnd1, AtEnd2")
		w.Header().Add("Trailer", "AtEnd3")

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		w.Header().Set("AtEnd1", "value 1")
		io.WriteString(w, "This HTTP response has both headers before this text and trailers at the end.\n")
		w.Header().Set("AtEnd2", "value 2")
		w.Header().Set("AtEnd3", "value 3")
	})
	log.Fatal(http.ListenAndServe(":8000", nil))

	// Simple static webserver
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}

func main() {
	handleRequests()
}
