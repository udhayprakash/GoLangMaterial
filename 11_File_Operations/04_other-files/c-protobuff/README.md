# Protobuf (Protocol Buffers)

Protobuf is a binary serialization format developed by Google. It requires a .proto schema file to define the data structure.

## setup

### Install Protobuf Compiler:

    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

### define schema (.proto file)

    syntax = "proto3";

    package example;

    message Person {
    string name = 1;
    int32 age = 2;
    string email = 3;
    }

### Generate Go code 

    protoc --go_out=. person.proto
