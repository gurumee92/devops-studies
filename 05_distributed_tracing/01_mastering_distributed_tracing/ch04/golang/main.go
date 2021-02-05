package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"studydts/lib/tracing"
	"studydts/people"

	"github.com/opentracing/opentracing-go"
	otLog "github.com/opentracing/opentracing-go/log"
)

var repo *people.Repository

// var tracer opentracing.Tracer

func main() {
	repo = people.NewRepository()
	defer repo.Close()

	tracer, closer := tracing.Init("go-2-hello")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/sayHello/", handleSayHello)
	log.Println("Listen on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSayHello(w http.ResponseWriter, r *http.Request) {
	span := opentracing.GlobalTracer().StartSpan("say-hello")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(r.Context(), span)
	name := strings.TrimPrefix(r.URL.Path, "/sayHello/")
	greeting, err := SayHello(ctx, name)

	if err != nil {
		span.SetTag("error", true)
		span.LogFields(otLog.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	span.SetTag("response", greeting)
	w.Write([]byte(greeting))
}

func SayHello(ctx context.Context, name string) (string, error) {
	person, err := repo.GetPerson(ctx, name)

	if err != nil {
		return "", err
	}

	opentracing.SpanFromContext(ctx).LogKV(
		"name", person.Name,
		"title", person.Title,
		"description", person.Description,
	)

	return FormatGreeting(
		ctx,
		person.Name,
		person.Title,
		person.Description,
	), nil
}

func FormatGreeting(
	ctx context.Context,
	name, title, description string,
) string {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"format-greeting",
	)
	defer span.Finish()
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
