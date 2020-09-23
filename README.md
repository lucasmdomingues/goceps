### Installation

```go
go get github.com/lucasmdomingues/goceps
```

### Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/lucasmdomingues/goceps"
)

func main() {
	service := goceps.NewService()

	_, err := service.Search("01001-000")
	if err != nil {
		log.Fatal(err)
	}
}
```
### Via CEP
https://viacep.com.br/

