from flask import Flask
import opentracing

import requests
import json

from database import Person
from lib.tracing import init_tracer


app = Flask("service-hello")
init_tracer("service-hello")

@app.route("/sayHello/<name>")
def say_hello(name):
    with opentracing.tracer.start_active_span("say-hello") as scope:    
        person = get_person(name)
        resp = format_greeting(person)
        scope.span.set_tag('response', resp)
        return resp


def get_person(name):
    with opentracing.tracer.start_active_span("get-person") as scope:   
        url = 'http://localhost:8081/getPerson/%s' % name
        resp = requests.get(url)
        
        assert resp.status_code == 200
        
        person = json.loads(resp.text)
        scope.span.log_kv({
            'name': person['name'],
            'title': person['title'],
            'description': person['description'],
        })
        return person


def format_greeting(person):
    with opentracing.tracer.start_active_span("foramt-greeting"):   
        url = 'http://localhost:8082/formatGreeting'
        resp = requests.get(url, params=person)
        assert resp.status_code == 200
        return resp.text


if __name__ == "__main__":
    app.run(port=8080)