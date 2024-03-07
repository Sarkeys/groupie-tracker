package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Relation struct {
	DatesLocation map[string][]string `json:"datesLocations"`
}

func GetArtists() ([]Artists, error) {
	var artists []Artists
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return artists, err
	}
	err = json.Unmarshal(bytes, &artists)
	if err != nil {
		return artists, err
	}
	defer response.Body.Close()
	return artists, nil
}

func Artist(id int) (Artists, error) {
	var artist Artists
	response, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id))
	if err != nil {
		return artist, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return artist, err
	}
	err = json.Unmarshal(bytes, &artist)
	if err != nil {
		return artist, err
	}
	defer response.Body.Close()
	return artist, nil
}

func GetRelations(id int) (Relation, error) {
	var relat Relation
	response, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id))
	if err != nil {
		return relat, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return relat, err
	}
	err = json.Unmarshal(bytes, &relat)
	if err != nil {
		return relat, err
	}
	defer response.Body.Close()
	return relat, nil
}
