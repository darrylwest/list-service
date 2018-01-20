#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2018-01-19 15:14:53
#

host="http://localhost:9040"

data='{"title":"carrots","category":"produce","status":"open"}'

curl -d "$data" -X POST $host/list

