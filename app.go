package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var methods = [7]string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"}

type FEReturn struct {
	Data        any         `json:"data"`
	Error       string      `json:"error"`
	RequestData RequestData `json:"requestData"`
}

type HttpResponse struct {
	Status     string              `json:"status"`
	StatusCode int                 `json:"statusCode"`
	Header     map[string][]string `json:"header"`
	Body       string              `json:"body"`
	Time       time.Duration       `json:"time"`
}

type RequestData struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

type DirEntry struct {
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	IsDir       bool        `json:"isDir"`
	RequestData RequestData `json:"requestData"`
}

// App struct
type App struct {
	ctx         context.Context
	currentFile string
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

func (a *App) validData(data RequestData) bool {
	allowedMethod := []string{"get", "post", "put", "delete", "patch", "options", "head"}

	// Check if method not in list
	if !contains(strings.ToLower(data.Method), allowedMethod) {
		return false
	}
	return true
}

func (a *App) ListDir() []DirEntry {
	// Get files in home dir
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dir, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	dirs := make([]DirEntry, 0)
	for _, d := range dir {
		// Check if name ends with

		if !d.IsDir() && strings.HasSuffix(d.Name(), ".goget.json") {
			data, err := a.getRequestDataFromFile(dirname + "/" + d.Name())
			if err == nil {
				dirs = append(dirs, DirEntry{Name: d.Name(), IsDir: d.IsDir(), RequestData: data, Path: dirname + "/" + d.Name()})
			}

		} else if d.IsDir() {
			dirs = append(dirs, DirEntry{Name: d.Name(), IsDir: d.IsDir(), Path: dirname + "/" + d.Name()})
		}
	}
	return dirs
}

func (a *App) getRequestDataFromFile(filepath string) (RequestData, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	// Read file
	buffer, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	// Convert to json
	var data RequestData
	err = json.Unmarshal(buffer, &data)
	if err != nil {
		log.Fatal(err)
	}

	if a.validData(data) {
		return data, nil
	}

	return RequestData{}, errors.New("Invalid data")
}

func (a *App) SetFile(filepath string) FEReturn {

	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return FEReturn{Error: "File does not exist"}
	}

	a.currentFile = filepath
	return FEReturn{}
}

func (a *App) GetFile() string {
	return a.currentFile
}

func (a *App) GetRequestDataFromFile(filepath string) FEReturn {
	data, err := a.getRequestDataFromFile(filepath)
	if err != nil {
		return FEReturn{Error: err.Error()}
	}
	return FEReturn{RequestData: data}
}
