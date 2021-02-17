package com.gurumee.opentracing.component;

import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;

@Configuration
public class AppConfig {
    @Bean
    public io.opentracing.Tracer initTracer() {
        io.jaegertracing.Configuration.SamplerConfiguration samplerConfiguration = new io.jaegertracing.Configuration.SamplerConfiguration()
                .withType("const").withParam(1)
                ;
        io.jaegertracing.Configuration.ReporterConfiguration reporterConfiguration = new io.jaegertracing.Configuration.ReporterConfiguration()
                .withLogSpans(true)
                ;
        return new io.jaegertracing.Configuration("service-hello")
                .withSampler(samplerConfiguration)
                .withReporter(reporterConfiguration)
                .getTracer();
    }
    @Bean
    public RestTemplate restTemplate() {
        return new RestTemplateBuilder().build();
    }
}
