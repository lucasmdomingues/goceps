### Instalação

```go
go get github.com/lucasmdomingues/goceps
```

### Exemplo

```go
package main

import (
	"fmt"
	"log"

	"github.com/lucasmdomingues/goceps"
)

func main() {
	
	// Utilize um CEP com ou sem carateres especiais(Ex: 07748415-415)
	endereco, err := goceps.BuscaEndereco("07748415")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(endereco)
}
```
### ReceitaWS
https://www.receitaws.com.br/
