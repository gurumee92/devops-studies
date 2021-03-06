from flask import Flask, request
from flask_opentracing import FlaskTracer
import opentracing
from opentracing.ext import tags
from opentracing_instrumentation.client_hooks import install_all_patches

import requests
import json

from database import Person
from lib.tracing import init_tracer, flask_to_scope


app = Flask("service-hello")
init_tracer("service-hello")
install_all_patches()
flask_tracer = FlaskTracer(opentracing.tracer, True, app)

@app.route("/sayHello/<name>")
def say_hello(name):
     with flask_to_scope(flask_tracer, request) as scope:  
        person = get_person(name)
        resp = format_greeting(person)
        scope.span.set_tag('response', resp)
        return resp


def _get(url, params=None):
    r = requests.get(url, params=params)
    assert r.status_code == 200
    return r.text


def get_person(name):
    with opentracing.tracer.start_active_span("get-person") as scope:   
        url = 'http://localhost:8081/getPerson/%s' % name
        res = _get(url)
        person = json.loads(res)
        scope.span.log_kv({
            'name': person['name'],
            'title': person['title'],
            'description': person['description'],
        })
        return person


def format_greeting(person):
    with opentracing.tracer.start_active_span("foramt-greeting"):   
        url = 'http://localhost:8082/formatGreeting'
        return _get(url, params=person)


if __name__ == "__main__":
    app.run(port=8080)