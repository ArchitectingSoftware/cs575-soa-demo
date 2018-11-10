name := "rest-scala"

version := "0.1"

scalaVersion := "2.12.7"

val circeVersion = "0.9.0-M2"

libraryDependencies ++= Seq(
  "org.json4s"        %% "json4s-jackson"% "3.6.2",
  "com.typesafe.akka" %% "akka-http" % "10.1.5",
  "com.typesafe.akka" %% "akka-stream" % "2.5.18",
  "com.typesafe.akka" %% "akka-actor"  % "2.5.18"
)


Revolver.settings
