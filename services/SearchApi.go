package services

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	models "gateway/models"
	"io/ioutil"
	"log"
	"net/http"
)

func GeneralSearch(busqueda string) (movies []models.Movie, songs []models.Song, tvShows []models.TvShow, persons []models.Person) {
	songs, movies = SearchSongsAndMovies(busqueda)
	tvShows = SearchTvShow(busqueda)
	persons = SearchPerson(busqueda)
	return
}
func SearchPerson(busqueda string) (persons []models.Person) {
	//Itunes URL
	url := "http://www.crcind.com/csp/samples/SOAP.Demo.cls?soap_method=GetByName&name=" + busqueda
	//Call Construction
	response, err := http.Get(url)
	//Have a error?
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	//Read
	responseData, err := ioutil.ReadAll(response.Body)
	//Have a error with the response?
	if err != nil {
		log.Fatal(err)
		return
	}
	//Convert to Object
	envelope := models.Envelope{}
	xml.Unmarshal(responseData, &envelope)
	var person models.Person
	for i := 0; i < len(envelope.Body.GetByNameResponse.GetByNameResult.Diffgram.ListByName.SQL); i++ {
		person.Name = envelope.Body.GetByNameResponse.GetByNameResult.Diffgram.ListByName.SQL[i].Name
		person.DOB = envelope.Body.GetByNameResponse.GetByNameResult.Diffgram.ListByName.SQL[i].DOB
		person.SSN = envelope.Body.GetByNameResponse.GetByNameResult.Diffgram.ListByName.SQL[i].SSN
		persons = append(persons, person)
	}
	return
}
func SearchTvShow(busqueda string) (tvShows []models.TvShow) {
	//Itunes URL
	url := "https://api.tvmaze.com/search/shows?q=" + busqueda
	//Call Construction
	response, err := http.Get(url)
	//Have a error?
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	//Read
	responseData, err := ioutil.ReadAll(response.Body)
	//Have a error with the response?
	if err != nil {
		log.Fatal(err)
		return
	}
	//Convert to Object
	var responseObject []models.TvShowResponse
	json.Unmarshal(responseData, &responseObject)
	var tvShow models.TvShow
	for i := 1; i < len(responseObject); i++ {
		if responseObject[i].Score > 0.5 {
			tvShow.Id = responseObject[i].Show.Id
			tvShow.Name = responseObject[i].Show.Name
			tvShow.Status = responseObject[i].Show.Status
			tvShow.Language = responseObject[i].Show.Language
			tvShow.Genres = responseObject[i].Show.Genres
			tvShow.ScoreMatch = responseObject[i].Show.ScoreMatch
			tvShow.Premiered = responseObject[i].Show.Premiered
			tvShow.Ended = responseObject[i].Show.Ended
		}

	}
	tvShows = append(tvShows, tvShow)
	return
}
func SearchSongsAndMovies(busqueda string) (songs []models.Song, movies []models.Movie) {
	//Itunes URL
	url := "https://itunes.apple.com/search?term=" + busqueda
	//Call Construction
	response, err := http.Get(url)
	//Have a error?
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	//Read
	responseData, err := ioutil.ReadAll(response.Body)
	//Have a error with the response?
	if err != nil {
		log.Fatal(err)
		return
	}
	//Convert to Object
	var responseObject models.MovieResponse
	json.Unmarshal(responseData, &responseObject)
	var movie models.Movie
	var song models.Song
	for i := 1; i < responseObject.ResultCont; i++ {
		if responseObject.Results[i].Kind == "song" {
			song.ArtistName = responseObject.Results[i].ArtistName
			song.TrackName = responseObject.Results[i].TrackName
			song.TrackPrice = responseObject.Results[i].TrackPrice
			songs = append(songs, song)
		} else {
			movie.ArtistName = responseObject.Results[i].ArtistName
			movie.TrackName = responseObject.Results[i].TrackName
			movie.TrackPrice = responseObject.Results[i].TrackPrice
			movies = append(movies, movie)
		}
	}
	return
}
