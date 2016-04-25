#!/bin/bash - 
protoc -I . ucrehelloworld.proto --go_out=plugins=grpc:./../pkg/ucrehelloworld/protos
