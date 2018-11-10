package edu.drexel.ws.kotlin.rest.demo

import com.fasterxml.jackson.core.type.TypeReference
import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.module.kotlin.*
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.core.io.ClassPathResource
import org.springframework.stereotype.Service
import java.util.HashMap
import java.util.stream.Collectors


@Service
object DBMock {

    private val pCache: Map<Int,Publication>
        get() {
            val mapper = jacksonObjectMapper();
            val pubDB = HashMap<Int, Publication>()

            val jsonDB = ClassPathResource("pubs.json").file
            val pubList = mapper.readValue<List<Publication>>(jsonDB)

            return pubList.associate { p -> Pair(p.id, p)}
        }

    val allPubs: Collection<Publication>
        get() {
            return pCache.values
        }

    fun getPubById(id: Int): Publication? {
        return pCache[id]
    }
}
