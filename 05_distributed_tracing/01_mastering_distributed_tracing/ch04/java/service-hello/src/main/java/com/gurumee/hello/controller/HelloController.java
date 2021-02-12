package com.gurumee.hello.controller;

import com.gurumee.hello.person.Person;
import com.gurumee.hello.person.PersonRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.util.Optional;

@RestController
@RequiredArgsConstructor
public class HelloController {
    private final PersonRepository personRepository;

    @GetMapping("/sayHello/{name}")
    public String sayHello(@PathVariable String name) {
        Person person = getPerson(name);
        String response = formatGreeting(person);
        return response;
    }

    private String formatGreeting(Person person) {
        String response = "Hello, ";
        if (person.getTitle() != null && !person.getTitle().isBlank()) {
            response += person.getTitle() + " ";
        }

        response += person.getName() + "!";

        if (person.getTitle() != null && !person.getDescription().isBlank()) {
            response += " " + person.getDescription();
        }

        return response + "\n";
    }

    private Person getPerson(String name) {
        Optional<Person> personOptional = personRepository.findById(name);
        return personOptional.orElseGet(() -> new Person(name));
    }

}
