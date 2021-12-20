package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Member struct {
	Name             string
	Email            string
	RegistrationDate string
}

type App struct {
	Members []Member
}

func (app *App) run() {
	r := mux.NewRouter()

	r.HandleFunc("/", app.Index).Methods("GET")
	r.HandleFunc("/", app.AddMember).Methods("POST")

	loggedRouter := handlers.LoggingHandler(os.Stderr, r)

	err := http.ListenAndServe(":8080", loggedRouter)

	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) Index(w http.ResponseWriter, r *http.Request) {
	templates := GetTemplates()

	t, err := template.ParseFS(templates, "templates/index.html.tmpl")

	if err != nil {
		log.Panic(err)
	}

	errorCookie, err := r.Cookie("errorMessage")

	errorMessage := ""

	if err == nil {
		errorMessage = errorCookie.Value
		errorCookie.Value = ""
		http.SetCookie(w, errorCookie)
	}

	err = t.Execute(w, IndexContext{
		Members:      &app.Members,
		ErrorMessage: errorMessage,
	})

	if err != nil {
		log.Panic(err)
	}
}

func (app *App) AddMember(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")

	validName, err := regexp.MatchString(".+", name)

	if err != nil {
		log.Panic(err)
	}

	if !validName {
		respondWithError(w, r, "The name is not a proper name!")
		return
	}

	email := r.FormValue("email")

	validEmail, err := regexp.MatchString(".+\\@.+\\..+", email)

	if err != nil {
		log.Panic(err)
	}

	if !validEmail {
		respondWithError(w, r, "The email is invalid!")
		return
	}

	for _, member := range app.Members {
		if member.Email == email {
			respondWithError(w, r, "A user with this email is already registered!")
			return
		}
	}

	currentTime := time.Now()

	app.Members = append(app.Members, Member{
		Name:             name,
		Email:            email,
		RegistrationDate: currentTime.Format(time.ANSIC),
	})

	http.Redirect(w, r, "/", http.StatusFound)
}

func respondWithError(w http.ResponseWriter, r *http.Request, err string) {
	log.Printf("Responding with error: %v", err)

	http.SetCookie(w, &http.Cookie{
		Name:     "errorMessage",
		Value:    err,
		SameSite: http.SameSiteDefaultMode,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
