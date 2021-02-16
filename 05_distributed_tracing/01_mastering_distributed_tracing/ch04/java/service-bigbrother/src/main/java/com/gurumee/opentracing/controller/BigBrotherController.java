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
public class BigBrotherController {
    private final PersonRepository personRepository;
    private final Tracer tracer;

    @GetMapping("/getPerson/{name}")
    public Person getPerson(@PathVariable String name) {
        Span span = tracer.buildSpan("/getPerson").start();
        try (Scope s = tracer.scopeManager().activate(span, false)){
            Person person = loadPerson(name);
            Map<String, String> fields = new LinkedHashMap<>();
            fields.put("name", person.getName());
            fields.put("title", person.getTitle());
            fields.put("description", person.getDescription());
            span.log(fields);
            return person;
        } finally {
            span.finish();
        }
    }


    private Person loadPerson(String name) {
        Span span = tracer.buildSpan("get-person").start();
        try (Scope s = tracer.scopeManager().activate(span, false)){
            Optional<Person> personOptional = personRepository.findById(name);
            return personOptional.orElseGet(() -> new Person(name));
        } finally {
            span.finish();
        }
    }

}
