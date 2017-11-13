package edu.drexel.ws.soap.demo;


import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import edu.drexel.ws.messages.PublicationType;
import org.springframework.core.io.ClassPathResource;
import org.springframework.stereotype.Service;

import java.io.File;
import java.util.Collection;
import java.util.HashMap;
import java.util.List;

@Service
public class DBMock {

    private HashMap<Integer,PublicationType> pubCache = null;

    public HashMap<Integer,PublicationType> getFullPubList()
    {
        return queryPubs();
    }

    public PublicationType getPubById(int id)
    {
        HashMap<Integer,PublicationType> pubDB = queryPubs();
        Integer index = new Integer(id);
        return pubDB.get(index);
    }

    public Collection<PublicationType> getAllPubs()
    {
        HashMap<Integer,PublicationType> pubDB = queryPubs();
        return pubDB.values();
    }

    private HashMap<Integer,PublicationType> queryPubs()
    {
        if(pubCache != null) return pubCache;

        ObjectMapper mapper = new ObjectMapper();
        HashMap<Integer,PublicationType> pubDB = new HashMap<Integer,PublicationType>();

        try{

            File jsonDB  = new ClassPathResource("pubs.json").getFile();

            List<HashMap<String,Object>> pubList = mapper.readValue(jsonDB, new TypeReference<List<HashMap<String,Object>>>(){});

            for( HashMap<String,Object> p : pubList)
            {
                PublicationType pType = new PublicationType();
                Integer idx = new Integer(p.get("id").toString());
                pType.setAbstract(p.get("abstract").toString());
                pType.setCite(p.get("cite").toString());
                pType.setId(idx.intValue());
                pType.setTitle(p.get("title").toString());

                pubDB.put(idx,pType);
            }

            pubCache = pubDB;
            return pubDB;
        }catch(Exception e)
        {
            e.printStackTrace();
            pubCache = null;
            return null;
        }
    }
}

