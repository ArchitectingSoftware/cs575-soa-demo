// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var proto_pubs_pb = require('../proto/pubs_pb.js');

function serialize_pubs_PubRequest(arg) {
  if (!(arg instanceof proto_pubs_pb.PubRequest)) {
    throw new Error('Expected argument of type pubs.PubRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_pubs_PubRequest(buffer_arg) {
  return proto_pubs_pb.PubRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pubs_PubResponse(arg) {
  if (!(arg instanceof proto_pubs_pb.PubResponse)) {
    throw new Error('Expected argument of type pubs.PubResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_pubs_PubResponse(buffer_arg) {
  return proto_pubs_pb.PubResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The greeting service definition.
var PubServiceService = exports.PubServiceService = {
  // Sends a greeting
  getPub: {
    path: '/pubs.PubService/GetPub',
    requestStream: false,
    responseStream: false,
    requestType: proto_pubs_pb.PubRequest,
    responseType: proto_pubs_pb.PubResponse,
    requestSerialize: serialize_pubs_PubRequest,
    requestDeserialize: deserialize_pubs_PubRequest,
    responseSerialize: serialize_pubs_PubResponse,
    responseDeserialize: deserialize_pubs_PubResponse,
  },
};

exports.PubServiceClient = grpc.makeGenericClientConstructor(PubServiceService);
