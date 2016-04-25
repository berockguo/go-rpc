#!/bin/bash - 
protoc -I . common.proto --go_out=plugins=grpc:./../pkg/common/proto
protoc -I . ucrehelloworld.proto --go_out=plugins=grpc:./../pkg/ucrehelloworld/proto
