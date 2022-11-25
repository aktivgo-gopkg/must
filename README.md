# Example

```go
package main

import (
	"errors"
	"github.com/aktivgo-gopkg/unexpected"
)

func main() {
	// panic
	unexpected.Panic(func() error {
		return errors.New("error")
	})
	
	unexpected.Panic(func() error {
		return nil
	})

	// panic
	result := unexpected.Must(func() (any, error) {
		return nil, errors.New("error")
	})
	
	result := unexpected.Must(func() (any, error) {
		return 0, nil
	})
}
```