// This modules creates a webserver that listens for incoming HTTP requests
// and writes it to a file. It also provides a web interface to view the file.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"os"
	"example.com/httprepeater/models"
	"example.com/httprepeater/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/TwiN/go-color"
)


func main() {
	stage := os.Getenv("STAGE")
	log.Printf(color.Ize(color.Green, fmt.Sprintf("Starting application for %s environment", stage )))
	storage.SetupDBConnection(stage)
	log.Default().Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe(":8000", router()))
}

// create new router 
// define routes and middleware
func router() http.Handler {
	router:= chi.NewRouter()
	router.Use(middleware.Logger)
	
	router.Route("/api", func(r chi.Router) {
		r.Get("/person/{id}", GetPerson)
		r.Post("/person", CreatePerson)
		r.Post("/person/email/{id}", ChangeMail)
		r.Put("/person/{id}", UpdatePerson)
		r.Put("/person/birthday/{id}", Birthday)
		r.Delete("/person/{id}", DeletePerson)
	})
	log.Printf(color.Ize(color.Green, "Startup successfull - awaiting traffic"))
	return router
}



func CreatePerson(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var person models.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Default().Println("Json decode failed")
	}
	
	// create person in database
	storage.NewPerson(person)	
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	// get persion with id in url from database
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	person := storage.GetPersonByID(id)
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Name: %s\nAge: %d\nBirthday: %s\nEmail: %s\n", person.Name, person.Age, person.Birthday, person.Email)
	w.Write([]byte(response))	
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	// get id from url and convert to uint
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	
	// decode request body
	var person models.Person
	json.NewDecoder(r.Body).Decode(&person)

	// update person in database
	storage.UpdatePerson(id, person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	// get id from url and convert to uint
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	// delete person in database
	storage.DeletePersonById(id)
}

func Birthday(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	person := storage.GetPersonByID(id)
	person.CelebrateBirthday()
	storage.UpdatePerson(id, person)
}

func ChangeMail(w http.ResponseWriter, r *http.Request)  {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	person := storage.GetPersonByID(id)
	json.NewDecoder(r.Body).Decode(&person)
	person.ChangeMail(person.Email)
	storage.UpdatePerson(id, person)
}