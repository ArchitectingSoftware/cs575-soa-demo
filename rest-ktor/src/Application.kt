package edu.drexel.ws.ktor.rest

import io.ktor.application.*
import io.ktor.response.*
import io.ktor.request.*
import io.ktor.features.*
import io.ktor.routing.*
import io.ktor.http.*
import io.ktor.gson.*
import com.fasterxml.jackson.databind.*
import io.ktor.jackson.*

fun main(args: Array<String>): Unit = io.ktor.server.netty.EngineMain.main(args)

@Suppress("unused") // Referenced in application.conf
@kotlin.jvm.JvmOverloads
fun Application.module(testing: Boolean = false) {

    val cache = PubUtil.getAll()

    install(CORS) {
        method(HttpMethod.Options)
        method(HttpMethod.Put)
        method(HttpMethod.Delete)
        method(HttpMethod.Patch)
        header(HttpHeaders.Authorization)
        header("MyCustomHeader")
        allowCredentials = true
        anyHost() // @TODO: Don't do this in production if possible. Try to limit it.
    }

    install(ContentNegotiation) {
        gson {
        }

        jackson {
            enable(SerializationFeature.INDENT_OUTPUT)
        }
    }

    routing {
        get("/") {
            call.respondText("HELLO WORLD!", contentType = ContentType.Text.Plain)
        }

        get("/publications") {
            call.respond(PubUtil.getAll())
        }

        get("/publications/{id}") {
            val idField = call.parameters["id"] ?: throw IllegalArgumentException("Parameter id not found")
            val id = try{
                idField.toInt()
            } catch(e:Exception){
                throw IllegalArgumentException("Parameter Id must be an integer")
            }

            val pub = PubUtil.getById(id)
            if (pub == null)
                call.respondText("Supplied Parameter with value ${id} is not in range", status=HttpStatusCode.NotFound)
            else
                call.respond(pub!!)
        }

        get("/json/gson") {
            call.respond(mapOf("hello" to "world"))
        }

        get("/json/jackson") {
            call.respond(mapOf("hello" to "world"))
        }
    }
}

