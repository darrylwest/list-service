#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2018-01-19 15:14:53
#

host="http://localhost:8080"

# '{"id":"01C45CNR35NZV2AH14XKJA97Q9","dateCreated":"2018-01-18T11:44:40.000Z","lastUpdated":"2018-01-18T11:44:40.000Z","version":1,"title":"my list item entry","category":"","attributes":{}, "status":"open"}

data='{"title":"my grocery item","category":"produce","status":"open"}'

curl -d "$data" -X POST $host/list

