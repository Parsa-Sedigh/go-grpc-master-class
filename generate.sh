#!/bin/bash

# this doesn't work:
# protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
# protoc ./greet/greetpb/greet.proto --go-grpc_out=.
protoc --proto_path=greet/greetpb --go_out=greet/greetpb --go-grpc_out=. --go_opt=paths=source_relative greet.proto

#protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculator.proto --go-grpc_out=.

#protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.
protoc --proto_path=blog/blogpb --go_out=blog/blogpb --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. --go_opt=paths=source_relative blog.proto