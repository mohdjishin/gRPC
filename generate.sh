#!/bin/bash

protoc  --go_out=greet/greetpb/ --go-grpc_out=greet/greetpb/  greet/greetpb/greet.proto 