from flask import Flask
import opentracing

from database import Person
from lib.tracing import init_tracer
import json

app = Flask("service-bigbrother")
init_tracer("service-bigbrother")

@app.route("/getPerson/<name>")
def get_person(name):
    with opentracing.tracer.start_active_span("/getPerson") as scope:    
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