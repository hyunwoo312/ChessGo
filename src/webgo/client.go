package chessgo

import (
	"fmt"
	"errors"
	//"crypto/rand"
	//"crypto/aes"
	//"bytes"
	"database/sql"
	"net/http"
	"html/template"
	"time"
	"io/ioutil"
	"log"
	"runtime"
	"github.com/hyunwoo312/chessai/lib/chessgo/game"
)
template.ParseFiles = runtime.GOOS, errors.New("aa")
// Client object to process the chess game
type Client struct {
	Address string
	Chatlog *sql.DB
	Logs *log.Logger
	ServerUpTime time.Duration // time.Since() to get uptime
	Users []User
}

type User struct {
	UserID string
	Profile Profile
	Status byte
	CurrentGame int64
	DateJoined time.Time
}

type Profile struct {
	Wins int
	Losses int
	Draws int
	Games []game.Game 
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
