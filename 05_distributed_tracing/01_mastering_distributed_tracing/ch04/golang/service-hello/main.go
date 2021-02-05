package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"studydts/lib/model"
	"studydts/lib/tracing"

	"github.com/opentracing/opentracing-go"
	otLog "github.com/opentracing/opentracing-go/log"
)

func main() {
	tracer, closer := tracing.Init("service-hello")
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

// SayHello is..
func SayHello(ctx context.Context, name string) (string, error) {
	person, err := getPeron(ctx, name)
	if err != nil {
		return "", err
	}

	return formatGreeting(ctx, person)
}

func getPeron(ctx context.Context, name string) (*model.Person, error) {
	res, err := http.Get("http://localhost:8081/getPerson/" + name)
	if err != nil {
		fmt.Println(err)

		return nil, err
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var person model.Person
	if err := json.Unmarshal(bytes, &person); err != nil {
		return nil, err
	}

	return &person, nil
}

func formatGreeting(ctx context.Context, person *model.Person) (string, error) {
	v := url.Values{}
	v.Set("name", person.Name)
	v.Set("title", person.Title)
	v.Set("description", person.Description)
	url := "http://localhost:8082/formatGreeting?" + v.Encode()
	res, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	return string(bytes), nil
}
