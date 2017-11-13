package edu.drexel.ws.soap.demo;

import edu.drexel.ws.messages.PublicationRequest;
import edu.drexel.ws.messages.PublicationResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.ws.server.endpoint.annotation.Endpoint;
import org.springframework.ws.server.endpoint.annotation.PayloadRoot;
import org.springframework.ws.server.endpoint.annotation.RequestPayload;
import org.springframework.ws.server.endpoint.annotation.ResponsePayload;


@Endpoint
public class PublicationEndpoint {

    private  PublicationService pubService;

    @Autowired
    public PublicationEndpoint(PublicationService pubService) {
        this.pubService = pubService;
    }

    @PayloadRoot(localPart = "PublicationRequest", namespace = "http://drexel.edu/ws/messages")
    @ResponsePayload
    public PublicationResponse getPublications(@RequestPayload PublicationRequest pubRequest) {

        //Use the publication service POJO to handle reading from the JSON database
        return pubService.getPublicationsServiceRouter(pubRequest);
    }

/*    @Autowired
    public PublicationEndpoint(PublicationRepository countryRepository) {
        this.publicationRepository = countryRepository;
    }

    @PayloadRoot(namespace = NAMESPACE_URI, localPart = "getCountryRequest")
    @ResponsePayload
    public GetPublicationResponse getCountry(@RequestPayload GetPublicationRequest request) {
        GetPublicationResponse response = new GetPublicationResponse();
        response.setCountry(publicationRepository.findCountry(request.getName()));

        return response;
    }*/
}

