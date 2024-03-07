package handlers

import (
	"fmt"
	"groupie-tracker/api"
	"html/template"
	"net/http"
	"strconv"
)

type AllData struct {
	Artists []api.Artists
}

type ArtistRelation struct {
	api.Artists
	api.Relation
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, "Ooops! Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Errors(w, "Ooops! Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	artists, err := api.GetArtists()

	if err != nil {
		Errors(w, "Ooops! Internal Server Error", http.StatusInternalServerError)
		return
	}
	allInfo := AllData{artists}
	templates, err := template.ParseGlob("./web/template/home.html")
	if err != nil {
		Errors(w, "Ooops! Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = templates.ExecuteTemplate(w, "home.html", allInfo)
	if err != nil {
		Errors(w, "Ooops! Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func Artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Errors(w, "Ooops! Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/artist" {
		Errors(w, "Ooops! Not Found", http.StatusNotFound)
		return
	}

	artists, err := api.GetArtists()
	if err != nil {
		Errors(w, "Ooops! Internal Server Error", http.StatusInternalServerError)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 || id > len(artists) {
		Errors(w, "Ooops! Not Found", http.StatusNotFound)
		return
	}

	artistf, err := api.Artist(id)
	if err != nil {
		Errors(w, "Ooops! Internal Server Error", http.StatusInternalServerError)
		return
	}
	rel, err := api.GetRelations(id)
	if err != nil {
		Errors(w, "Ooops! Internal Server Error", http.StatusInternalServerError)
		return
	}

	info := ArtistRelation{artistf, rel}

	templates, err := template.ParseGlob("./web/template/artist.html")
	if err != nil {
		Errors(w, "Ooops! Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = templates.ExecuteTemplate(w, "artist.html", info)
	if err != nil {
		fmt.Println("gg")
		Errors(w, "Ooops! Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func Errors(w http.ResponseWriter, msgs string, status int) {
	templates, err1 := template.ParseFiles("./web/template/error.html")
	if err1 != nil {
		Errors(w, "Ooops! Internal Server Error", http.StatusInternalServerError)
		return
	}
	Error := struct {
		Msg  string
		Code int
	}{
		msgs,
		status,
	}
	err := templates.ExecuteTemplate(w, "error.html", Error)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
