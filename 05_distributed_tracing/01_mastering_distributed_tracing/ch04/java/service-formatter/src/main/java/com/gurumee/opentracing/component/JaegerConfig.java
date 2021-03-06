package com.gurumee.opentracing.component;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class JaegerConfig {
    @Bean
    public io.opentracing.Tracer initTracer() {
        io.jaegertracing.Configuration.SamplerConfiguration samplerConfiguration = new io.jaegertracing.Configuration.SamplerConfiguration()
                .withType("const").withParam(1)
                ;
        io.jaegertracing.Configuration.ReporterConfiguration reporterConfiguration = new io.jaegertracing.Configuration.ReporterConfiguration()
                .withLogSpans(true)
                ;
        return new io.jaegertracing.Configuration("service-formatter")
                .withSampler(samplerConfiguration)
                .withReporter(reporterConfiguration)
                .getTracer();
    }
}
