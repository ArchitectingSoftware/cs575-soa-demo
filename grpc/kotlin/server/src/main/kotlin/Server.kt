package pubsrpc.server

import io.grpc.Server
import io.grpc.ServerBuilder
import io.grpc.stub.StreamObserver
import pubsrpc.PubServiceGrpc
import java.util.logging.Logger
import java.io.IOException
import com.typesafe.config.Config
import com.typesafe.config.ConfigFactory
import pubsrpc.PubsRpcProto



class Server(private val config: Config) {
    private val logger: Logger
    private val server: Server

    init {
        val port = this.config.getInt("grpc.deployment.port")

        server = ServerBuilder.forPort(port)
            .addService(PubsServiceImpl())
            .build()
        logger = Logger.getLogger(Server::class.java.name)
    }


    @Throws(IOException::class)
    fun start() {

        server.start()

        logger.info("Server started, listening on ${this.server.port}")

        Runtime.getRuntime().addShutdownHook(object : Thread() {
            override fun run() {
                // Use stderr here since the logger may have been reset by its JVM shutdown hook.
                System.err.println("*** shutting down gRPC server since JVM is shutting down")

                this@Server.stop()

                System.err.println("*** server shut down")
            }
        })
    }

    fun stop() {
        this.server.shutdown()
    }

    @Throws(InterruptedException::class)
    fun blockUntilShutdown() {
        server.awaitTermination()
    }

    internal class PubsServiceImpl : PubServiceGrpc.PubServiceImplBase() {
    }
}


