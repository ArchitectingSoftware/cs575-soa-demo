package edu.drexel.ws.rest.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ServerWebExchange;

import java.util.Collection;


@RestController
public class PublicationsController {

    @Autowired
    DBMock dbService;

    @GetMapping("/publications")
    Iterable<Publication> getAllPubs(ServerWebExchange server) {
        Collection<Publication> pubList = dbService.getAllPubs(); //dbService.getAllPubs();
        if (pubList == null)
            server.getResponse().setStatusCode(HttpStatus.SERVICE_UNAVAILABLE);

        return pubList;
        // can also return a Flux<Publication> return Flux.fromIterable(pubList);
    }


    @GetMapping("/publications/{id}")
    Publication findById(@PathVariable Integer id, ServerWebExchange server) {
        Publication pub = dbService.getPubById(id);
        if(pub == null)
            server.getResponse().setStatusCode(HttpStatus.NOT_FOUND);

        return pub;
        // can also return a Mono<Publication> and return Mono.justOrEmpty(pub);
    }
}
