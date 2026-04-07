/*
substituir sync.WaitGroup por channels como mecanismo de sinalização.
No Código atual tem duas goroutines —uma contando de 1 a 10 e outra de 10 a 1 — onde a main() esperava ambas terminarem usando sync.WaitGroup.

Tarefa: modifique o programa para que as goroutines sinalizem seu término via canal em vez de WaitGroup. A main() deve receber dessas sinalizações para saber quando ambas terminaram.

Requisitos:
a) Crie um canal adequado para receber a sinalização de término das goroutines.
b) Cada goroutine deve enviar um valor no canal ao terminar sua contagem.
c) A main() deve receber do canal duas vezes (uma para cada goroutine) antes de
imprimir “Fim!”.
d) O programa não deve usar sync.WaitGroup.

Dicas:
• Um canal do tipo chan bool ou chan struct{} serve bem para sinalização pura
• Uma operação de recebimento (<-ch) bloqueia até que um valor esteja disponível
• Releia o exemplo exemplos/01-canal-sincrono.go para relembrar o comportamento de rendez-vous
Salve sua solução como ex1.go
*/
package main

import (
	"fmt"
	"time"
)

func crescente(ch chan struct{}) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("[Crescente] %d\n", i)
		time.Sleep(1 * time.Second)
	}
	ch <- struct{}{}
}

func decrescente(ch chan struct{}) {
	for i := 10; i >= 1; i-- {
		fmt.Printf("[Decrescente] %d\n", i)
		time.Sleep(1 * time.Second)
	}
	ch <- struct{}{}
}

func main() {
	ch := make(chan struct{})

	go crescente(ch)
	go decrescente(ch)

	<-ch
	<-ch
	fmt.Println("Fim!")
}
