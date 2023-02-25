#!/bin/bash

protoc  --go_out=greet/. --go-grpc_out=greet/.  greet/greetpb/greet.proto 
# protoc  --go_out=calculator/. --go-grpc_out=calculator/.  calculator/calc/calc.proto 