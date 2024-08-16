package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

var methods = [7]string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"}

type HttpResponse struct {
	Status     string              `json:"status"`
	StatusCode int                 `json:"statusCode"`
	Header     map[string][]string `json:"header"`
	Body       string              `json:"body"`
	Time       time.Duration       `json:"time"`
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Methods() [7]string {
	return methods
}

func (a *App) SendHttpRequest(url string) HttpResponse {
	// Time required for request
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	end := time.Now()
	// Total time
	log.Printf("Time required for request: %v\n", end.Sub(start))

	buffer, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return HttpResponse{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(buffer[:]),
		Time:       end.Sub(start),
	}
}

func (a *App) CopyToClipboard(text string) {

}
