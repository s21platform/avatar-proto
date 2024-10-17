#!/bin/bash

protoc  --go_out=./  \
        --go-grpc_out=./ \
        avatar.proto

protoc --doc_out=. --doc_opt=markdown,README.md ./avatar.proto

protoc  --go_out=./  \
        --go-grpc_out=./ \
        new_avatar_register.proto
