#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2018-04-24 09:03:35
#

NET=listnet

if docker network inspect $NET > /dev/null
then
    echo "network $NET up..."
else
    docker network create -d bridge $NET
fi
