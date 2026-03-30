// Exemplo: Múltiplas goroutines com contagem intercalada
package main

import (
	"fmt"
	"sync"
	"time"
)

func contagem(nome string, inicio, fim, passo int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := inicio; i != fim+passo; i += passo {
		fmt.Printf("[%s] %d\n", nome, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go contagem("Crescente", 1, 5, 1, &wg)
	go contagem("Decrescente", 5, 1, -1, &wg)
	wg.Wait()
	fmt.Println("Fim!")
}
