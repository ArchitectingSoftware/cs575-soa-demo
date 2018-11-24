package edu.drexel.ws.ktor.rest

import java.util.*
import com.google.gson.*

data class RequestError(var id:Int, var msg:String)

data class Publication (
    var id: Int,
    var title: String,
    var cite: String,
    var slides: String?,
    var link: String?,
    var url: String?,
    var pubDate: Date?,
    var `abstract`: String?,
    var error: RequestError?)

object PubUtil{

    private val pCache = LoadFromJson();

    fun LoadFromJson() : Array<Publication>{
        val jsonDB = javaClass.getResource("/pubs.json").readText();
        val gson = Gson()

        val pA = gson.fromJson<Array<Publication>>(jsonDB, Array<Publication>::class.java)

        return pA;
    }

    fun getAll() : Array<Publication>{
        return pCache
    }

    fun getById(id: Int): Publication?{
        if (id > 0 && id <= pCache.size)
            return pCache[id-1]
        else
            return null
    }
}
