#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2018-04-24 08:57:31
#

CONTAINER=lister1
USER=lister

docker exec -it $CONTAINER ./cockroach user set $USER --insecure \
    && docker exec -it $CONTAINER ./cockroach sql --insecure -e 'create database lists' \
    && docker exec -it $CONTAINER ./cockroach sql --insecure -e 'grant all on database lists to lister'
