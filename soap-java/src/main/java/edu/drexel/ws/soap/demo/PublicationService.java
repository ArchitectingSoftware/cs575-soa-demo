package edu.drexel.ws.soap.demo;

import edu.drexel.ws.messages.*;

public interface PublicationService {
    PublicationResponse getPublicationsServiceRouter(PublicationRequest req);
    PublicationResponse getAllPublications();
    PublicationResponse getAPublication(int id);
}
