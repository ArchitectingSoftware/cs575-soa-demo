package edu.drexel.ws.scala.rest.demo

case class PubResults (
                         results: Array[Publication],
                         metadata: MetaData
                       )

case class MetaData(
                     count : Int,
                     error : Boolean,
                     statusCode: Int,
                     message: String = ""
                   )
