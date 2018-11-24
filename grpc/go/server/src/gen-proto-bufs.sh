#!/bin/bash
protoc server/pubs.proto --go_out=plugins=grpc:.