Instalação
========

```go
go get github.com/lucasmdomingues/goceps
```

### Como usar:

```go
package main

import (
	"fmt"
	"log"

	"github.com/lucasmdomingues/goceps"
)

func main() {

	endereco, err := goceps.BuscaEndereco("07748415")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(endereco)
}
```
