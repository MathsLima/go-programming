/*
construir um pipeline com goroutines conectadas por canais e contém a estrutura do programa
com as assinaturas das funções, mas os corpos estão incompletos.

arefa: complete as três funções do pipeline:
1. gerador — recebe uma lista de números inteiros e envia cada um no canal de
saída. Fecha o canal ao terminar.
2. multiplicador — lê cada valor do canal de entrada, multiplica por 2, e envia no
canal de saída. Fecha o canal de saída ao terminar.
3. impressora — lê cada valor do canal de entrada e imprime na tela. Envia true
no canal done ao terminar.

Requisitos:
a) Use range para iterar sobre os canais de entrada.
b) Cada estágio deve fechar seu canal de saída com close() quando terminar de
enviar.
c) A main() deve esperar o pipeline terminar usando o canal done.
d) Teste com os valores: 1, 2, 3, 4, 5. A saída esperada é: 2, 4, 6, 8, 10
*/
package main

import "fmt"

// gerador envia os números de 'valores' no canal 'out' e fecha o canal.
func gerador(valores []int, out chan<- int) {
	for _, v := range valores {
		out <- v
	}
	close(out)
}

// multiplicador lê do canal 'in', multiplica por 2, e envia no canal 'out'.
// Fecha 'out' ao terminar.
func multiplicador(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * 2
	}
	close(out)
}

// impressora lê do canal 'in' e imprime cada valor.
// Envia true no canal 'done' ao terminar.
func impressora(in <-chan int, done chan<- bool) {
	for v := range in {
		fmt.Println(v)
	}
	done <- true
}

func main() {
	valores := []int{1, 2, 3, 4, 5}

	c1 := make(chan int)
	c2 := make(chan int)
	done := make(chan bool)

	go gerador(valores, c1)
	go multiplicador(c1, c2)
	go impressora(c2, done)

	<-done
	fmt.Println("Pipeline concluído!")
}
