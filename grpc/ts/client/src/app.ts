import * as grpc from 'grpc'
import * as messages from '../proto/pubs_pb'
import * as services from '../proto/pubs_grpc_pb'

var client = new services.PubServiceClient('localhost:50051', 
            grpc.credentials.createInsecure());

var msg = new messages.PubRequest()
msg.setPubid(3)

client.getPub(msg, (err, pub) => { 
    console.log('pub',pub)
    console.log('err', err)
 });
