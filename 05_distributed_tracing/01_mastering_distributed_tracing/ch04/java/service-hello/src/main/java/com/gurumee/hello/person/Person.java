package com.gurumee.hello.person;

import lombok.*;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name="people")
@NoArgsConstructor @AllArgsConstructor
@Getter @Setter @ToString
public class Person {
    @Id
    private String name;

    @Column(nullable = false)
    private String title;

    @Column(nullable = false)
    private String description;

    public Person(String name) {
        this.name = name;
    }
}
