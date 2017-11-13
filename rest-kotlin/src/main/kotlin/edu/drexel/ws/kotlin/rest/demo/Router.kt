package edu.drexel.ws.kotlin.rest.demo

import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration

import org.springframework.web.reactive.function.server.router
import org.springframework.http.MediaType

@Configuration
class Router(private val serviceHandler: Handler){

    @Bean
    fun boot() = router {
        (accept(MediaType.APPLICATION_JSON) and "/").nest {
            GET("/publications/{id}", serviceHandler::getPubById)
            GET("/publications", serviceHandler::getAll)
            GET("/pubstream", serviceHandler::getStream)
        }
    }

}