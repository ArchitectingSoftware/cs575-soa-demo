package edu.drexel.ws.scala.rest.demo

import akka.actor.ActorSystem
import akka.http.scaladsl.Http
import akka.http.scaladsl.server.Route
import akka.http.scaladsl.model.{ContentTypes, HttpEntity}
import akka.http.scaladsl.server.HttpApp
import akka.stream.ActorMaterializer
import akka.util.Timeout
import com.typesafe.config.ConfigFactory

import scala.concurrent.duration._


object Boot extends App{
  // we need an ActorSystem to host our application in
  implicit val system = ActorSystem("http-server")

  implicit val mat = ActorMaterializer()
  implicit val dispatch = system.dispatcher

  val config = ConfigFactory.load()
  val appConfig = config.getConfig("application")

  implicit val timeout = Timeout(5.seconds)
  // start a new HTTP server on port 8080 with our service actor as the handler
  //IO(Http) ? Http.Bind(service, interface="localhost", port = appConfig.getInt("port"))

  val host = sys.env.get("BC_SERVER_HOST").getOrElse(appConfig.getString("host"))
  val port = sys.env.get("BC_SERVER_PORT").getOrElse(appConfig.getString("port")).toInt

  System.out.println("starting server")
  val server = Http().bindAndHandle(
    handler = HttpServiceRouter.mainRoute(),
    interface = host,
    port = port
  )


  server.failed.foreach( err => {
    println(err.toString)
  })

  System.out.println(s"server started on ${host}:${port}")

}


