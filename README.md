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

	return a / b, nil
}

func main() {
	must.Do1(checkEqual(1, 2)) // panic

	must.Do1(checkEqual(2, 1)) // ok

	_ = must.Do(sub(1, 2)) // or Do2, panic

	result := must.Do(sub(2, 1)) // or Do2, ok
	log.Println(result)          // 1

	// etc. with Do3 and Do4
}
```