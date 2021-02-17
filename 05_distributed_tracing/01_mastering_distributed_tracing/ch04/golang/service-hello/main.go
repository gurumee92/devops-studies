package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"studydts/lib/model"
	"studydts/lib/othttp"
	"studydts/lib/tracing"
	"studydts/lib/xhttp"

	opentracing "github.com/opentracing/opentracing-go"
	otTag "github.com/opentracing/opentracing-go/ext"
	otLog "github.com/opentracing/opentracing-go/log"
)

func main() {
	tracer, closer := tracing.Init("service-hello")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/sayHello/", handleSayHello)
	othttp.ListenAndServe(":8080", "/sayHello")
}

func handleSayHello(w http.ResponseWriter, r *http.Request) {
	span := opentracing.SpanFromContext(r.Context())
	name := strings.TrimPrefix(r.URL.Path, "/sayHello/")
	greeting, err := SayHello(r.Context(), name)
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

func get(ctx context.Context, operationName, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	otTag.SpanKindRPCClient.Set(span)
	otTag.HTTPUrl.Set(span, url)
	otTag.HTTPMethod.Set(span, "GET")
	opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	return xhttp.Do(req)
}

func getPeron(ctx context.Context, name string) (*model.Person, error) {
	url := "http://localhost:8081/getPerson/" + name
	res, err := get(ctx, "getPerson", url)
	if err != nil {
		return nil, err
	}

	var person model.Person
	if err := json.Unmarshal(res, &person); err != nil {
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
	res, err := get(ctx, "formatGreeting", url)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
