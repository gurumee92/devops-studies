from flask import Flask
from flask import request
import opentracing
from opentracing.ext import tags

from database import Person
from lib.tracing import init_tracer

app = Flask("service-formatter")
init_tracer("service-formatter")

@app.route("/formatGreeting")
def handle_format_greeting():
    span_ctx = opentracing.tracer.extract(
        opentracing.Format.HTTP_HEADERS,
        request.headers,
    )
    with opentracing.tracer.start_active_span(
        "/formatGreeting",
        child_of=span_ctx,
        tags={tags.SPAN_KIND: tags.SPAN_KIND_RPC_SERVER},
    ) as scope:    
        name = request.args.get('name')
        title = request.args.get('title')
        description = request.args.get('description')

        return format_greeting(
            name=name,
            title=title,
            description=description,
        )

def format_greeting(name, title, description):
    with opentracing.tracer.start_active_span("foramt-greeting"):   
        greeting  = "Hello, "

        if title:
            greeting += title + " "
        
        greeting += name + "!"

        if description:
            greeting += " " + description
        
        return greeting + "\n"


if __name__ == "__main__":
    app.run(port=8082)