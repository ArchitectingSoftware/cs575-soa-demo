package edu.drexel.ws.kotlin.rest.demo

import org.springframework.web.reactive.function.server.ServerRequest
import org.springframework.web.reactive.function.server.ServerResponse
import org.springframework.http.MediaType.APPLICATION_JSON
import org.springframework.stereotype.Component
import org.springframework.web.reactive.function.BodyInserters
import org.springframework.web.reactive.function.server.bodyToServerSentEvents
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono
import java.time.Duration

@Component
class Handler{


    fun getPubById(req:ServerRequest) : Mono<ServerResponse>{
        val id = req.pathVariable("id").toInt()
        val p = DBMock.getPubById(id);
        if (p == null){
            return ServerResponse.notFound().build()
        }
        return ServerResponse.ok().body(BodyInserters.fromObject(p))
    }

    fun getAll(req:ServerRequest) : Mono<ServerResponse>{
        val p = DBMock.allPubs
        if (p == null){
            return ServerResponse.notFound().build()
        }
        return ServerResponse.ok().body(BodyInserters.fromObject(p))
    }


    fun getStream(req:ServerRequest) : Mono<ServerResponse>{
        val p = DBMock.allPubs
        if (p == null){
            return ServerResponse.notFound().build()
        }

        val dataStream = Flux.fromIterable(p)
        val stream = dataStream.delayElements(Duration.ofSeconds(2))

        return ServerResponse.ok()
                .contentType(APPLICATION_JSON)
                .bodyToServerSentEvents(stream)
    }
}

