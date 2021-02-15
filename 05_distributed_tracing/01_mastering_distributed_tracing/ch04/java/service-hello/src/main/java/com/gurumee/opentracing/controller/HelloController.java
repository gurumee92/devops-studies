package com.gurumee.opentracing.controller;

import com.gurumee.opentracing.model.Person;
import com.gurumee.opentracing.model.PersonRepository;
import io.opentracing.Scope;
import io.opentracing.Span;
import io.opentracing.Tracer;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.util.LinkedHashMap;
import java.util.Map;
import java.util.Optional;

@RestController
@RequiredArgsConstructor
public class HelloController {
    private final PersonRepository personRepository;
    private final Tracer tracer;

    @GetMapping("/sayHello/{name}")
    public String sayHello(@PathVariable String name) {
        Span span = tracer.buildSpan("say-hello").start();
        try (Scope s = tracer.scopeManager().activate(span, false)){
            Person person = getPerson(name);
            Map<String, String> fields = new LinkedHashMap<>();
            fields.put("name", person.getName());
            fields.put("title", person.getTitle());
            fields.put("description", person.getDescription());
            span.log(fields);

            String response = formatGreeting(person);
            span.setTag("response", response);

            return response;
        } finally {
            span.finish();
        }
    }

    private String formatGreeting(Person person) {
        Span span = tracer.buildSpan("format-greeting").start();
        try (Scope s = tracer.scopeManager().activate(span, false)){
            String response = "Hello, ";
            if (person.getTitle() != null && !person.getTitle().isBlank()) {
                response += person.getTitle() + " ";
            }

            response += person.getName() + "!";

            if (person.getTitle() != null && !person.getDescription().isBlank()) {
                response += " " + person.getDescription();
            }

            return response + "\n";
        } finally {
            span.finish();
        }
    }

    private Person getPerson(String name) {
        Span span = tracer.buildSpan("get-person").start();
        try (Scope s = tracer.scopeManager().activate(span, false)){
            Optional<Person> personOptional = personRepository.findById(name);
            return personOptional.orElseGet(() -> new Person(name));
        } finally {
            span.finish();
        }
    }

}
