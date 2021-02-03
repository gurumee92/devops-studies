package main

import (
	"log"
	"net/http"
	"strings"
	"studydts/exercise1/people"
	"studydts/lib/tracing"

	"github.com/opentracing/opentracing-go"
	otLog "github.com/opentracing/opentracing-go/log"
)

var repo *people.Repository
var tracer opentracing.Tracer

func main() {
	repo = people.NewRepository()
	defer repo.Close()

	tr, closer := tracing.Init("go-2-hello")
	defer closer.Close()
	tracer = tr

	http.HandleFunc("/sayHello/", handleSayHello)
	log.Println("Listen on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSayHello(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpan("say-hello")
	defer span.Finish()

	name := strings.TrimPrefix(r.URL.Path, "/sayHello/")
	greeting, err := SayHello(name, span)

	if err != nil {
		span.SetTag("error", true)
		span.LogFields(otLog.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	span.SetTag("response", greeting)
	w.Write([]byte(greeting))
}

func SayHello(name string, span opentracing.Span) (string, error) {
	person, err := repo.GetPerson(name)

	if err != nil {
		return "", err
	}

	span.LogKV(
		"name", person.Name,
		"title", person.Title,
		"description", person.Description,
	)

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
