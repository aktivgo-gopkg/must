# Example

```go
package main

import (
	"errors"
	"github.com/aktivgo-gopkg/unexpected"
	"log"
)

func compare(a, b int) error {
	if a < b {
		return errors.New("a less then b")
	}

	return nil
}

func sub(a, b int) (int, error) {
	if a < b {
		return 0, errors.New("a less then b")
	}

	return a - b, nil
}

func main() {
	unexpected.Panic(compare(1, 2)) // panic

	unexpected.Panic(compare(2, 1)) // ok

	_ = unexpected.Must(sub(1, 2)) // panic

	result := unexpected.Must(sub(2, 1)) // ok
	log.Println(result)                  // 1
}
```