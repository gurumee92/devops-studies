from flask import Flask
from flask import request
from flask_opentracing import FlaskTracer
import opentracing
from opentracing.ext import tags

from database import Person
from lib.tracing import init_tracer, flask_to_scope

app = Flask("service-formatter")
init_tracer("service-formatter")
flask_tracer = FlaskTracer(opentracing.tracer, True, app)


@app.route("/formatGreeting")
def handle_format_greeting():
    with flask_to_scope(flask_tracer, request) as scope: 
        name = request.args.get('name')
        title = request.args.get('title')
        description = request.args.get('description')

        return format_greeting(
            name=name,
            title=title,
            description=description,
        )

def format_greeting(name, title, description):
    with opentracing.tracer.start_active_span("foramt-greeting") as scope:
        greeting  = (scope.span.get_baggage_item("greeting") or "Hello") + " "

        if title:
            greeting += title + " "
        
        greeting += name + "!"

        if description:
            greeting += " " + description
        
        return greeting + "\n"


if __name__ == "__main__":
    app.run(port=8082)