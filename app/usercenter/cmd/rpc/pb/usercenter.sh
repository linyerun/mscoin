#!/bin/bash

goctl rpc protoc *.proto --go_out=../pb --go-grpc_out=../pb --zrpc_out=.. --style=go_zero