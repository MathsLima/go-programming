// Exemplo: Correção da condição de corrida com Mutex
// Execute com: go run -race 07-mutex.go
package main

import (
	"fmt"
	"sync"
)

var (
	contador int = 0
	mutex    sync.Mutex
)

func incrementar(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		contador++
		mutex.Unlock()
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
