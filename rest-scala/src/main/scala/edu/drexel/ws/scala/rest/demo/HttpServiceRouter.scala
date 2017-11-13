package edu.drexel.ws.scala.rest.demo

import akka.actor.{ActorSystem, Props}
import akka.http.scaladsl.model.{ContentTypes, HttpEntity}
import akka.http.scaladsl.server.Directives._
import akka.http.scaladsl.server.PathMatchers.IntNumber
import akka.http.scaladsl.server.Route
import akka.stream.Materializer
import akka.util.Timeout
import org.json4s._
import org.json4s.JsonDSL._
import org.json4s.jackson.JsonMethods._
import org.json4s.jackson.Serialization

import scala.concurrent.{ExecutionContext, Future}

object HttpServiceRouter {

  def mainRoute()(implicit sys: ActorSystem, mat: Materializer, dis: ExecutionContext): Route = {

    implicit val formats = DefaultFormats
    implicit val serializationFormats = Serialization.formats(NoTypeHints)

    implicit val system = ActorSystem()
    import akka.pattern.ask

      path("publications") {
        get {
          complete {
            DBMock.getAll
          }
        }
      } ~
      path("publications" / IntNumber) { pubId =>
        get {
          complete {
            DBMock.getPub(pubId)
          }
        }
      }
    }
}

