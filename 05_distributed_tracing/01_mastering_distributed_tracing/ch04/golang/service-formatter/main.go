package main

import (
	"context"
	"log"
	"net/http"
	"studydts/lib/tracing"

	opentracing "github.com/opentracing/opentracing-go"
	otTag "github.com/opentracing/opentracing-go/ext"
)

func main() {
	tracer, closer := tracing.Init("service-formatter")
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/formatGreeting/", handleFormatGreeting)
	log.Print("Listen port :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func handleFormatGreeting(w http.ResponseWriter, r *http.Request) {
	spanCtx, _ := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header),
	)
	span := opentracing.GlobalTracer().StartSpan(
		"/formatGreeting",
		otTag.RPCServerOption(spanCtx),
	)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(r.Context(), span)
	name := r.FormValue("name")
	title := r.FormValue("title")
	description := r.FormValue("description")

	greeting := FormatGreeting(ctx, name, title, description)
	w.Write([]byte(greeting))
}

// FormatGreeting combines information about a person into a greeting string.
func FormatGreeting(
	ctx context.Context,
	name, title, description string,
) string {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"format-greeting",
	)
	defer span.Finish()

	greeting := span.BaggageItem("greeting")
	if greeting == "" {
		greeting = "Hello"
	}

	response := greeting + ", "
	if title != "" {
		response += title + " "
	}
	response += name + "!"
	if description != "" {
		response += " " + description
	}
	return response + "\n"
}
