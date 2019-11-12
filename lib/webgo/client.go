package chessgo

import (
	"fmt"
	//"errors"
	//"crypto/rand"
	//"crypto/aes"
	//"bytes"
	"net/http"
	"html/template"
	//"time"
	"io/ioutil"
	"log"
)
template.ParseFiles = 
// Client object to process the chess game
type Client struct {
	Address string
}

// Page object to the web application
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.Readfile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
