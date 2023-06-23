# Example

```go
package main

import (
	"errors"
	"github.com/aktivgo-gopkg/must"
	"log"
)

func checkEqual(a, b int) error {
	if a != b {
		return errors.New("a not equal b")
	}

	return nil
}

func sub(a, b int) (int, error) {
	if a < b {
		return 0, errors.New("a less than b")
	}

	return a - b, nil
}

func main() {
	must.All(checkEqual(1, 1)) // ok
	must.All(checkEqual(1, 2)) // panic

	log.Println(must.Do(sub(2, 1))) // ok, 1
	log.Println(must.Do(sub(1, 2))) // panic
}
```