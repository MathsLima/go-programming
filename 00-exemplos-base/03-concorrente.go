// Exemplo: Chamada concorrente com goroutine
package main

import (
	"fmt"
	"time"
)

func diga(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go diga("Olá!") // goroutine — executa em paralelo
	diga("Mundo")   // executa na main
}
