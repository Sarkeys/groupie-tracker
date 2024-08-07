package web

import (
	"bytes"
	"embed"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	Config   Config
}

type Config struct {
	ArtistsURL string
}

type ApplicationError struct {
	Message string
	Code    int
}

func NewApplication(errorLog, infoLog *log.Logger) *Application {
	return &Application{
		errorLog: errorLog,
		infoLog:  infoLog,
		Config: Config{
			ArtistsURL: "https://groupietrackers.herokuapp.com/api/artists",
		},
	}
}

func NewServer(addr *string, errorLog *log.Logger, mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}

//go:embed templates/*.html
var fs embed.FS

var (
	templates *template.Template
	err       error
)

func init() {
	templates, err = template.ParseFS(fs, "templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
}

func (app *Application) InternalServerError(w http.ResponseWriter, err error) {
	app.errorLog.Println(err)
	app.Errors(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	app.Errors(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}

func (app *Application) BadRequest(w http.ResponseWriter) {
	app.ClientError(w, http.StatusBadRequest)
}

func (app *Application) MethodNotAllowed(w http.ResponseWriter) {
	app.ClientError(w, http.StatusMethodNotAllowed)
}

func (app *Application) Errors(w http.ResponseWriter, errorMessage string, errorCode int) {
	buf := new(bytes.Buffer)
	if err := templates.ExecuteTemplate(buf, "error.html", ApplicationError{
		Message: errorMessage,
		Code:    errorCode,
	}); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}
