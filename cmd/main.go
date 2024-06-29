package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"groupie-tracker/web"
)

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес веб-сервера")
	flag.Parse()

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := web.NewApplication(errorLog, infoLog)
	server := web.NewServer(addr, errorLog, app.Routes())

	infoLog.Printf("Launching the server on an http://localhost%s", *addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		errorLog.Printf("error when starting the server on http://localhost%s", *addr)
		errorLog.Println(err)
	}
}
