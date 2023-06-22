#!/bin/bash

# this doesn't work:
# protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc ./greet/greetpb/greet.proto --go-grpc_out=.

#protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculator.proto --go-grpc_out=.

protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.