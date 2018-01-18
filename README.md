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

## Rest API

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

