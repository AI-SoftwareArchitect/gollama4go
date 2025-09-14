# gollama4go
gollama4go is ollama sdk for go


## Quic start
```go
package main

import (
	"fmt"
	"gollama4go/entities"
)

func main() {
	g4g := entities.NewG4GClient(
		"http://127.0.0.1:11434",
		20,
	)
	models, err := g4g.ListOllamaModels()
	if err != nil {
		return
	}
	fmt.Println(models)

}
```
