# go shutdown

### example
```
package main

import (
	"fmt"
	"github.com/minms/shutdown"
)

func main() {
	shutdown.WaitTerminationSignal(func() {
		fmt.Println("shutdown")
	})
}

```