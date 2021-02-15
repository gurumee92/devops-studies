package com.gurumee.opentracing.model;

import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.Optional;

import static org.junit.Assert.*;

@RunWith(SpringRunner.class)
@DataJpaTest
public class PersonRepositoryTest {
    @Autowired
    private PersonRepository personRepository;

    @Before
    public void setUp() {
        personRepository.deleteAll();
        personRepository.save(Person.builder()
                .name("test")
                .description("test")
                .title("test")
                .build());
    }

    @Test
    public void test() {
        Optional<Person> personOptional = personRepository.findById("test");
        assertTrue(personOptional.isPresent());

        Person person = personOptional.get();
        assertEquals("test", person.getTitle());
        assertEquals("test", person.getDescription());
    }

    @Test
    public void test2() {
        Optional<Person> personOptional = personRepository.findById("test");
        assertTrue(personOptional.isPresent());

        Person person = personOptional.get();
        assertEquals("test", person.getTitle());
        assertEquals("test", person.getDescription());
    }
}