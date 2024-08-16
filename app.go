package main

import (
	"context"
	"io"
	"log"
	"net/http"
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

func (a *App) SendHttpRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	buffer, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(buffer[:])
}
