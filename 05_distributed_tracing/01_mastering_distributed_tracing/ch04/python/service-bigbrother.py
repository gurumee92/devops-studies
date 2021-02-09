from flask import Flask
from flask import request
import opentracing

from database import Person
from lib.tracing import init_tracer
import json

app = Flask("service-bigbrother")
init_tracer("service-bigbrother")

@app.route("/getPerson/<name>")
def get_person(name):
    span_ctx = opentracing.tracer.extract(
        opentracing.Format.HTTP_HEADERS,
        request.headers,
    )
    with opentracing.tracer.start_active_span(
        "/getPerson",
        child_of=span_ctx
    ) as scope:    
        person = Person.get(name)

        if person is None:
            person = Person()
            person.name = name

        scope.span.log_kv({
            "name": person.name,
            "title": person.title,
            "description": person.description, 
        })
        return json.dumps({
            "name": person.name,
            "title": person.title,
            "description": person.description, 
        })


if __name__ == "__main__":
    app.run(port=8081)