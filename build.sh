#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o server ./server.go || exit 1
docker build -t matthieujacquot/http-logger:0.1 . || exit 1