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
                              +------------+           List Service-1
                              |            |         +---------------+
                              |            |-------->| http/rest     |
                              |            |<--------| list.db       |
                              | http/rest  |         +---------------+     List Service-2
                              |            |                             +---------------+
                              |            |---------------------------->| http/rest     |
                              |            |<----------------------------| list.db       |
                              |            |                             |               |
          +-----------+       |            |           List Service-3    +---------------+
          |           |       |            |         +---------------+
          | Sidecar   |<----->|            |-------->| http/rest     |
          | Container |       |            |<--------| list.db       |
          |           |       |            |         |               |
          +-----------+       +------------+         +---------------+
```

## Rest API

### Proxy Prefix

Internal requests use the following endpoints but are usually prefixed when exposed to the web. Prefixes are specific to the list type, so a ToDo list may have a prefix of `/todoapi` to distinguesh from a `/shopapi`. The proxy strips this prefix off prior to forwarding the request, so the following API is unchanged across various implementations.

* GET  /list/query - return zero or more items from the list based on query parameters
* GET  /list/:id   - return the list item by id
* POST /list/      - insert a new list item; list data is posted as a json blob
* PUT  /list/:id   - update the list item; list data is posted as a json blob
* DEL  /list/:id   - remove the list item (or archive it)
* PUT  /db/backup
* GET  /status
* GET  /logger
* PUT  /logger/:n - set the log level 1..5

## Document Dataset

The list model is quite basic.  Attribuites enable extending the base document model.  Models are serialized to JSON prior to saving to database.

```
generic list schema
    id string // ulid
    dateCreated time.Time // ISO8601 / RFC3339 
    lastUpdated time.Time // ISO8601 / RFC3339 
    version int64
    category string   // optional category or list of categories
    title string      // the primary list item title
    attributes map[string]interface{} // adhoc attributes to support various applications
    status string     // open | closed | archived
```

###### darryl.west | 2018.01.18

