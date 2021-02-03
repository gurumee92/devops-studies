package main

import (
	"log"
	"net/http"
	"strings"
	"studydts/exercise1/people"
)

var repo *people.Repository

func main() {
	repo = people.NewRepository()
	defer repo.Close()

	http.HandleFunc("/sayHello/", handleSayHello)
	log.Println("Listen on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSayHello(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/sayHello/")
	greeting, err := SayHello(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(greeting))
}

func SayHello(name string) (string, error) {
	person, err := repo.GetPerson(name)

	if err != nil {
		return "", err
	}

	return FormatGreeting(
		person.Name,
		person.Title,
		person.Description,
	), nil
}

func FormatGreeting(name, title, description string) string {
	response := "Hello, "

	if title != "" {
		response += title + " "
	}

	response += name + "!"

	if description != "" {
		response += " " + description
	}

	response += "\n"
	return response
}
