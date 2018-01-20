#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2018-01-20 13:05:09
#

host="http://localhost:8080"

data='{"id":"01C4ANWMRRD3KHDS8Q5STDN93A","dateCreated":"2018-01-20T12:59:18.680736562-08:00","lastUpdated":"2018-01-20T12:59:18.680738118-08:00","version":1,"title":"A new title","category":"produce","status":"open"}'
id='01C4ANWMRRD3KHDS8Q5STDN93A'

curl -d "$data" -X PUT $host/list/$id

