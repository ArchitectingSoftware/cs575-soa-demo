package edu.drexel.ws.scala.rest.demo

case class Publication ( id:      Option[Integer] = None,
                         title:   Option[String] = None,
                         cite:    Option[String] = None,
                         slides:  Option[String]= None,
                         link:    Option[String]= None,
                         url:     Option[String]= None,
                         pubDate: Option[String]= None,
                         `abstract`: Option[String]= None,
                         error:   Option[PublicationError]= None)
