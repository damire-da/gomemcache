# Go MemCache

MemCache is a simple in-memory cache for storing and managing data.

# Install

```bash
go get github.com/damire-da/gomemcache
```

# Example

```go
package main

import (
	"fmt"
	"github.com/damire-da/gomemcache"
)

func main() {
	memcache := memcache.New()

	memcache.Set("1", 1)
	fmt.Println(memcache.Get("1"))

	memcache.Delete("1")

	fmt.Println(memcache.Get("1")) // nil
}
```
