#!/bin/bash - 
protoc -I protos protos/*.proto --go_out=plugins=grpc:protos
