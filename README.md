# List Service

```
    __    _      __     _____                 _         
   / /   (_)____/ /_   / ___/___  ______   __(_)_______ 
  / /   / / ___/ __/   \__ \/ _ \/ ___/ | / / / ___/ _ \
 / /___/ (__  ) /_    ___/ /  __/ /   | |/ / / /__/  __/
/_____/_/____/\__/   /____/\___/_/    |___/_/\___/\___/ 
                                                        
```

## Overview

A generic list service used for todo, grocery, menus, etc.

## Document Dataset

```
generic list schema
    key string // ulid
    lastUpdated int64 // time to the millisecond? or unix?
    category string   // optional category or list of categories
    title string      // the primary list item title
    attributes string // adhoc attributes to support various applications
    status string
```

###### darryl.west | 2018.01.16

