package com.gurumee.hello.lib.component;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class AppConfig {
    @Bean io.opentracing.Tracer initTracer() {
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
}
