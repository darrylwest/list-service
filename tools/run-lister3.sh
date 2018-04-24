#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2018-04-24 10:01:36
#

PORT=3882
NAME=lister3
NET=listnet

docker run -d --name=$NAME --hostname=$NAME --net=$NET \
    -p 26257:26257 \
    -p "$PORT":8080  \
    -v "${PWD}/cockroach-data/$NAME:/cockroach/cockroach-data" \
    cockroachdb/cockroach:v2.0.0 start --insecure --join=lister1

