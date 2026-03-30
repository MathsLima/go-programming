// Exemplo: Condição de corrida (data race)
// Execute com: go run -race 06-corrida.go
package main

import (
	"fmt"
	"sync"
)

var contador int = 0

func incrementar(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		contador++
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go incrementar(&wg)
	go incrementar(&wg)
	wg.Wait()
	fmt.Println("Contador:", contador)
}
