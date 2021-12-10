package handlers

import (
	"encoding/json"
	"fmt"
	models "gateway/models"
	services "gateway/services"
	"log"
	"net/http"

	_ "gateway/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Search API
// @version 1.0
// @description Service for search movies, songs, tvshows and people
// @host localhost:8080
// @BasePath /

func Api() {
	router := mux.NewRouter()
	router.HandleFunc("/search/{search}", Search).Methods("GET")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}

// GetSearch godoc
// @Summary Get items for search
// @Description Get movies, songs, tv shows and people for search
// @Tags search
// @Accept  plain
// @Produce  json
// @Success 200 {object} models.JsonResponse
// @Param search query string true "value to search"
// @Router /search/{search} [get]
func Search(w http.ResponseWriter, r *http.Request) {
	var response models.JsonResponse
	params := mux.Vars(r)
	busqueda := params["search"]
	response.Movies, response.Songs, response.TvShow, response.Person = services.GeneralSearch(busqueda)
	json.NewEncoder(w).Encode(response)
}

//param name,param type,data type,is mandatory?,comment attribute(optional)
