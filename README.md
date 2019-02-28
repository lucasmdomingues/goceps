Instalação
========

Para instalar o gorreios execute o comando abaixo em seu terminal:

```go
go get github.com/lucasmdomingues/goceps
```

### Como usar:

Para chamar a aplicação dentro do seu arquivo **main.go** escreva o código abaixo:

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
### Tests:

```
go test
```
