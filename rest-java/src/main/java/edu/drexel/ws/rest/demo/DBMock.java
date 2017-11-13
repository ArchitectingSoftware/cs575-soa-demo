package edu.drexel.ws.rest.demo;


import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.core.io.ClassPathResource;
import org.springframework.stereotype.Service;

import java.io.File;
import java.util.Collection;
import java.util.HashMap;
import java.util.List;

@Service
public class DBMock {

    private HashMap<Integer, Publication> pubCache = null;

    public HashMap<Integer, Publication> getFullPubList() {
        return queryPubs();
    }

    public Publication getPubById(int id) {
        HashMap<Integer, Publication> pubDB = queryPubs();
        Integer index = new Integer(id);
        return pubDB.get(index);
    }

    public Collection<Publication> getAllPubs() {
        HashMap<Integer, Publication> pubDB = queryPubs();
        return pubDB.values();
    }

    private HashMap<Integer, Publication> queryPubs() {
        if (pubCache != null) return pubCache;

        ObjectMapper mapper = new ObjectMapper();
        HashMap<Integer, Publication> pubDB = new HashMap<Integer, Publication>();

        try {
            File jsonDB = new ClassPathResource("pubs.json").getFile();

            List<Publication> pubList = mapper.readValue(jsonDB, new TypeReference<List<Publication>>() {});

            //List<HashMap<String,Object>> pubList2 = mapper.readValue(jsonDB, new TypeReference<List<HashMap<String,Object>>>(){});

            for (Publication p : pubList) {
                Integer idx = new Integer(p.getId());
                pubDB.put(idx, p);
            }
            pubCache = pubDB;
            return pubDB;
        } catch (Exception e) {
            e.printStackTrace();
            pubCache = null;
            return null;
        }
    }

}
