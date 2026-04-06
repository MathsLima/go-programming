// Exemplo: Chamadas sequenciais (sem concorrência)
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
	diga("Olá!")  // bloqueia até terminar
	diga("Mundo") // só executa depois
}
