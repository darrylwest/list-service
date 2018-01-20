#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2017-12-29 11:29:48
#

name=list-service
port=9040

docker run -itd --name $name -p $port:80 -v $HOME/database:/data darrylwest/list-service

sleep 2
curl http://localhost:$port/status | jq '.'

