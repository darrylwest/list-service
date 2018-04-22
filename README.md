# List Service

```
    __    _      __     _____                 _         
   / /   (_)____/ /_   / ___/___  ______   __(_)_______ 
  / /   / / ___/ __/   \__ \/ _ \/ ___/ | / / / ___/ _ \
 / /___/ (__  ) /_    ___/ /  __/ /   | |/ / / /__/  __/
/_____/_/____/\__/   /____/\___/_/    |___/_/\___/\___/ 
                                                        
```

## Overview

A generic list service used for todo, grocery, menus, etc.  The target design is for the application to run inside a container with a single purpose, i.e., a single list type.  


## Block Diagram

The test controller includes an http REST interface to respond to end point requests.  Requests are first queued then submitted to test runners to run test suites in parallel capped by a maximum limit.  

```
                      ... Docker Container Environment ...

       Edge Proxy
      +------------+          List Service-1
      |            |         +---------------+
      |            |-------->| http/rest     |
      |            |<--------|               |                     +------+
      | http/rest  |         +---------------+  List Service-2     |  db1 |
      |            |                           +--------------+    +------+
      |            |-------------------------->| http/rest    |    +------+
      |            |<--------------------------|              |    | db2  |
      |            |          List Service-3   +--------------+    +------+
      |            |         +--------------+                      +------+
      |            |-------->| http/rest    |                      | db3  |
      |            |<--------|              |                      +------+
      |            |         +--------------+
      +------------+         
```

## Rest API

### Proxy Prefix

Internal requests use the following endpoints but are usually prefixed when exposed to the web. Prefixes are specific to the list type, so a ToDo list may have a prefix of `/todoapi` to distinguesh from a `/shopapi`. The proxy strips this prefix off prior to forwarding the request, so the following API is unchanged across various implementations.

* GET  /           - return the html home page
* GET  /list/query - return zero or more items from the list based on query parameters
* GET  /list/:id   - return the list item by id
* POST /list/      - insert a new list item; list data is posted as a json blob
* PUT  /list/:id   - update the list item; list data is posted as a json blob
* DEL  /list/:id   - remove the list item (or archive it)
* GET  /status
* GET  /logger
* PUT  /logger/:n - set the log level 1..5

## Document Dataset

The list model is quite basic.  Attribuites enable extending the base document model.  Models are serialized to JSON prior to saving to database.

```
list schema
    id string // ulid
    dateCreated time.Time // ISO8601 / RFC3339 
    lastUpdated time.Time // ISO8601 / RFC3339 
    version int64
    owner string // primary ulid of user/owner
    name string      // the primary list name
    category string   // optional category eg., grocery, shopping, todo
    description string   // optional 
    info  map[string]interface{} // adhoc attributes to support various applications
    status string     // open | closed | archived

item schema
    id string // ulid
    dateCreated time.Time // ISO8601 / RFC3339 
    lastUpdated time.Time // ISO8601 / RFC3339 
    version int64

user schema
    id string // ulid
    dateCreated time.Time // ISO8601 / RFC3339 
    lastUpdated time.Time // ISO8601 / RFC3339 
    version int64

```

###### darryl.west | 2018.04.21

