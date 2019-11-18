name := "rest-scala"

version := "0.1"

scalaVersion := "2.13.1"

val circeVersion = "0.9.0-M2"

libraryDependencies ++= Seq(
  "org.json4s"        %% "json4s-jackson"% "3.6.7",
  "com.typesafe.akka" %% "akka-http" % "10.1.10",
  "com.typesafe.akka" %% "akka-stream" % "2.6.0",
  "com.typesafe.akka" %% "akka-actor"  % "2.6.0"
)


Revolver.settings
