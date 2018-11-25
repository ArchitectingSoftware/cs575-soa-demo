#!/bin/bash
protoc pubs-proto/pubs.proto --go_out=plugins=grpc:.
