/*
Exercício 2 — Canal sem buffer vs. com buffer
Objetivo: observar e comparar o comportamento de canais síncronos e assíncronos.
Crie um programa ex2.go com duas goroutines — uma produtora e uma consumidora — conectadas por um canal. A produtora envia 5 mensagens no canal, com
intervalo de 200ms entre cada envio. A consumidora recebe e imprime cada mensagem com intervalo de 1 segundo. A main() deve esperar a consumidora terminar (use
um canal de sinalização, como no Exercício 1).
Parte A — Canal sem buffer:
a) Use make(chan string) (sem buffer). Execute e observe o comportamento.
b) Responda: quem dita o ritmo da comunicação, o produtor ou o consumidor? Por
quê?
Parte B — Canal com buffer:
c) Altere para make(chan string, 5) (buffer de tamanho 5). Execute e observe.
d) Responda: o comportamento do produtor mudou? Por quê?
e) Experimente com make(chan string, 2) (buffer de tamanho 2). O que acontece
quando o buffer enche?
Requisitos:
• A goroutine produtora deve imprimir uma mensagem antes de cada envio (ex.:
“Enviando mensagem 1…”) para que você visualize quando o envio ocorre ou
bloqueia
• A goroutine consumidora deve imprimir o horário de recebimento de cada mensagem usando time.Now().Format("15:04:05")
• A produtora deve fechar o canal ao terminar de enviar
• A main() deve esperar a consumidora terminar antes de encerrar
*/

package main

import (
	"fmt"
	"time"
)

const bufferSize = 0

func produtora(ch chan<- string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Enviando mensagem %d...\n", i)
		ch <- fmt.Sprintf("Mensagem %d", i)
		if i < 5 {
			time.Sleep(200 * time.Millisecond)
		}
	}
	close(ch)
}

func consumidora(ch <-chan string, done chan<- struct{}) {
	for msg := range ch {
		fmt.Printf("[%s] Recebido: %s\n", time.Now().Format("15:04:05"), msg)
		time.Sleep(1 * time.Second)
	}
	done <- struct{}{}
}

func main() {
	ch := make(chan string, bufferSize)
	done := make(chan struct{})

	go produtora(ch)
	go consumidora(ch, done)

	<-done
	fmt.Println("Fim!")
}
