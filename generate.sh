#!/bin/bash

protoc  --go_out=greet/. --go-grpc_out=greet/.  greet/greetpb/greet.proto 
protoc  --go_out=calculator/. --go-grpc_out=calculator/.  calculator/calc/calc.proto 
protoc  --go_out=prime/. --go-grpc_out=prime/. prime/prime.proto 

protoc  --go_out=blog/. --go-grpc_out=blog/.  blog/blogpb/blog.proto 