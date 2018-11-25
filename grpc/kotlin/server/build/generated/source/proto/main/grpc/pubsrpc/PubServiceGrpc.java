package pubsrpc;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 * <pre>
 * The greeting service definition.
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.16.1)",
    comments = "Source: pubs.proto")
public final class PubServiceGrpc {

  private PubServiceGrpc() {}

  public static final String SERVICE_NAME = "pubsrpc.PubService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<pubsrpc.PubsRpcProto.PubRequest,
      pubsrpc.PubsRpcProto.PubsResponse> getGetPubMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPub",
      requestType = pubsrpc.PubsRpcProto.PubRequest.class,
      responseType = pubsrpc.PubsRpcProto.PubsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<pubsrpc.PubsRpcProto.PubRequest,
      pubsrpc.PubsRpcProto.PubsResponse> getGetPubMethod() {
    io.grpc.MethodDescriptor<pubsrpc.PubsRpcProto.PubRequest, pubsrpc.PubsRpcProto.PubsResponse> getGetPubMethod;
    if ((getGetPubMethod = PubServiceGrpc.getGetPubMethod) == null) {
      synchronized (PubServiceGrpc.class) {
        if ((getGetPubMethod = PubServiceGrpc.getGetPubMethod) == null) {
          PubServiceGrpc.getGetPubMethod = getGetPubMethod = 
              io.grpc.MethodDescriptor.<pubsrpc.PubsRpcProto.PubRequest, pubsrpc.PubsRpcProto.PubsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "pubsrpc.PubService", "GetPub"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  pubsrpc.PubsRpcProto.PubRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  pubsrpc.PubsRpcProto.PubsResponse.getDefaultInstance()))
                  .setSchemaDescriptor(new PubServiceMethodDescriptorSupplier("GetPub"))
                  .build();
          }
        }
     }
     return getGetPubMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static PubServiceStub newStub(io.grpc.Channel channel) {
    return new PubServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static PubServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new PubServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static PubServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new PubServiceFutureStub(channel);
  }

  /**
   * <pre>
   * The greeting service definition.
   * </pre>
   */
  public static abstract class PubServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * Sends a greeting
     * </pre>
     */
    public void getPub(pubsrpc.PubsRpcProto.PubRequest request,
        io.grpc.stub.StreamObserver<pubsrpc.PubsRpcProto.PubsResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getGetPubMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetPubMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                pubsrpc.PubsRpcProto.PubRequest,
                pubsrpc.PubsRpcProto.PubsResponse>(
                  this, METHODID_GET_PUB)))
          .build();
    }
  }

  /**
   * <pre>
   * The greeting service definition.
   * </pre>
   */
  public static final class PubServiceStub extends io.grpc.stub.AbstractStub<PubServiceStub> {
    private PubServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private PubServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PubServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new PubServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Sends a greeting
     * </pre>
     */
    public void getPub(pubsrpc.PubsRpcProto.PubRequest request,
        io.grpc.stub.StreamObserver<pubsrpc.PubsRpcProto.PubsResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getGetPubMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * The greeting service definition.
   * </pre>
   */
  public static final class PubServiceBlockingStub extends io.grpc.stub.AbstractStub<PubServiceBlockingStub> {
    private PubServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private PubServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PubServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new PubServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Sends a greeting
     * </pre>
     */
    public pubsrpc.PubsRpcProto.PubsResponse getPub(pubsrpc.PubsRpcProto.PubRequest request) {
      return blockingUnaryCall(
          getChannel(), getGetPubMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * The greeting service definition.
   * </pre>
   */
  public static final class PubServiceFutureStub extends io.grpc.stub.AbstractStub<PubServiceFutureStub> {
    private PubServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private PubServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PubServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new PubServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Sends a greeting
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<pubsrpc.PubsRpcProto.PubsResponse> getPub(
        pubsrpc.PubsRpcProto.PubRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getGetPubMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_PUB = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final PubServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(PubServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_PUB:
          serviceImpl.getPub((pubsrpc.PubsRpcProto.PubRequest) request,
              (io.grpc.stub.StreamObserver<pubsrpc.PubsRpcProto.PubsResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class PubServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    PubServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return pubsrpc.PubsRpcProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("PubService");
    }
  }

  private static final class PubServiceFileDescriptorSupplier
      extends PubServiceBaseDescriptorSupplier {
    PubServiceFileDescriptorSupplier() {}
  }

  private static final class PubServiceMethodDescriptorSupplier
      extends PubServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    PubServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (PubServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new PubServiceFileDescriptorSupplier())
              .addMethod(getGetPubMethod())
              .build();
        }
      }
    }
    return result;
  }
}
