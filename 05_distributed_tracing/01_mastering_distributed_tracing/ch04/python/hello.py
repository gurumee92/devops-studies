from flask import Flask
from database import Person
from lib.tracing import init_tracer
import opentracing

app = Flask("service-hello")
init_tracer("service-hello")

@app.route("/sayHello/<name>")
def say_hello(name):
    with opentracing.tracer.start_span("say-hello"):    
        person = get_person(name)
        resp = format_greeting(
            name=person.name,
            title=person.title,
            description=person.description,        
        )
        return resp


def get_person(name):
    person = Person.get(name)

    if person is None:
        person = Person()
        person.name = name
    
    return person


def format_greeting(name, title, description):
    greeting  = "Hello, "

    if title:
        greeting += title + " "
    
    greeting += name + "!"

    if description:
        greeting += " " + description
    
    return greeting + "\n"


if __name__ == "__main__":
    app.run(port=8080)