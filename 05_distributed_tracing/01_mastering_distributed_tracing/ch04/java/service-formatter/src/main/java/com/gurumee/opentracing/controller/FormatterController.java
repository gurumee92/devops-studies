package com.gurumee.opentracing.controller;

import io.opentracing.Scope;
import io.opentracing.Span;
import io.opentracing.Tracer;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequiredArgsConstructor
public class FormatterController {
    private final Tracer tracer;

    @GetMapping("/formatGreeting")
    public String formatGreeting(@RequestParam String name,
                                 @RequestParam String title,
                                 @RequestParam String description) {
        Span span = tracer.buildSpan("/formatGreeting").start();
        try (Scope s = tracer.scopeManager().activate(span, false)){
            String response = "Hello, ";
            if (!title.isBlank()) {
                response += title + " ";
            }

            response += name + "!";

            if (!description.isBlank()) {
                response += " " + description;
            }

            return response + "\n";
        } finally {
            span.finish();
        }
    }
}
