#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2018-04-24 09:20:01
#

PORT=3880
NAME=lister1
NET=listnet

docker run -d --name=$NAME --hostname=$NAME --net=$NET \
    -p 26257:26257 \
    -p "$PORT":8080  \
    -v "${PWD}/cockroach-data/$NAME:/cockroach/cockroach-data" \
    cockroachdb/cockroach:v2.0.0 start --insecure
