package main

import (
	"groupie-tracker/web"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHome(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := web.NewApplication(errorLog, infoLog)

	// os.Chdir("../../")
	test_list := []struct {
		path   string
		method string
		code   int
	}{
		{
			path:   "/",
			method: "GET",
			code:   200,
		},
		{
			path:   "/main",
			method: "GET",
			code:   404,
		},
		{
			path:   "/",
			method: "DELETE",
			code:   405,
		},
		{
			path:   "/",
			method: "POST",
			code:   405,
		},
		{
			path:   "/",
			method: "PUT",
			code:   405,
		},
		{
			path:   "/Home",
			method: "POST",
			code:   404,
		},
		{
			path:   "/Home",
			method: "GET",
			code:   404,
		},
		{
			path:   "/main.go",
			method: "POST",
			code:   404,
		},
		{
			path:   "/artist",
			method: "GET",
			code:   404,
		},
	}
	for _, cases := range test_list {
		req := httptest.NewRequest(cases.method, cases.path, nil)
		w := httptest.NewRecorder()
		app.Home(w, req)
		resp := w.Result()
		if resp.StatusCode != cases.code {
			t.Errorf("Expected status %v; got %v", cases.code, resp.Status)
		}
	}
}

func TestArtists(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := web.NewApplication(errorLog, infoLog)

	test_list := []struct {
		path   string
		method string
		code   int
	}{
		{
			path:   "/artists?id=21",
			method: "GET",
			code:   200,
		},
		{
			path:   "/artists?id=-1",
			method: "GET",
			code:   404,
		},
		{
			path:   "/artists?id=21",
			method: "DELETE",
			code:   405,
		},
		{
			path:   "/artists?id=2",
			method: "POST",
			code:   405,
		},
		{
			path:   "/artists?id=13",
			method: "PUT",
			code:   405,
		},
		{
			path:   "/artists?id=12&name=Akzhol",
			method: "GET",
			code:   404,
		},
		{
			path:   "/artists?id=715",
			method: "GET",
			code:   404,
		},
		{
			path:   "/artists?id=000001",
			method: "GET",
			code:   404,
		},
		{
			path:   "/artists?id=01",
			method: "GET",
			code:   404,
		},
		{
			path:   "/artists?id=17",
			method: "GET",
			code:   200,
		},
		{
			path:   "/artists?id=15",
			method: "DELETE",
			code:   405,
		},
		{
			path:   "/artists?id=LOOL",
			method: "GET",
			code:   404,
		},
		{
			path:   "/artists?id='A'",
			method: "GET",
			code:   404,
		},
		{
			path:   "/artists?id=12/3",
			method: "GET",
			code:   404,
		},
		{
			path:   "/artists?id===15",
			method: "GET",
			code:   404,
		},
		{
			path:   "/artists????id=15",
			method: "GET",
			code:   404,
		},
		{
			path:   "/artists?id=main.go",
			method: "GET",
			code:   404,
		},
	}
	for _, cases := range test_list {
		req, err := http.NewRequest(cases.method, cases.path, nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		app.Artist(w, req)
		resp := w.Result()
		if resp.StatusCode != cases.code {
			t.Errorf("Expected status %v; got %v", cases.code, resp.Status)
		}
	}
}

func TestGetResponse(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := web.NewApplication(errorLog, infoLog)
	w := httptest.NewRecorder()
	bk := app.GetResponse(w)
	if len(bk) != apiLen {
		t.Errorf("Expected len %v; got %v", apiLen, len(bk))
	}
}
