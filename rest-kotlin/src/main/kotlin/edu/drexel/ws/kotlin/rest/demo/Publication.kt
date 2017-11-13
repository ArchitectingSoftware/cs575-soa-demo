package edu.drexel.ws.kotlin.rest.demo

import com.fasterxml.jackson.annotation.JsonInclude
import java.util.*

data class RequestError(var id:Int, var msg:String)

@JsonInclude(JsonInclude.Include.NON_NULL)
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