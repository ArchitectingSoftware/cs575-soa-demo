name := "rest-scala"

version := "0.1"

scalaVersion := "2.12.4"

val circeVersion = "0.9.0-M2"

libraryDependencies ++= Seq(
  "org.json4s"        %% "json4s-jackson"% "3.5.3",
  "com.typesafe.akka" %% "akka-http" % "10.0.10",
  "com.typesafe.akka" %% "akka-stream" % "2.5.4",
  "com.typesafe.akka" %% "akka-actor"  % "2.5.4"
)


Revolver.settings
