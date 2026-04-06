/*
Duas goroutines com WaitGroup
Crie um programa em Go ( ex2.go ) com duas goroutines disparadas a partir da main() : 
- Uma goroutine deve imprimir números de 1 a 10, um a cada 1 segundo 
- A outra goroutine deve imprimir números de 10 a 1, um a cada 1 segundo 
- A main() deve esperar até que as duas goroutines terminem
- Utilize sync.WaitGroup para fazer a main() esperar
explique passo a passo a logica, a sintaxe e como funciona as coisas porque eu nao sei nada de go

Explicação do resultado:
As duas goroutines estão rodando ao mesmo tempo, de verdade. Nenhuma espera a outra.
A cada segundo, as duas "acordam" do time.Sleep quase simultaneamente e tentam imprimir,
e o escalonador do Go decide qual vai primeiro naquele momento e muda a cada execução
*/

package main

import (
    "fmt"
    "sync"
    "time"
)

//função conta de 1 a 10, um número por segundo
func contarCrescente(wg *sync.WaitGroup){
	defer wg.Done() //quando a funcao termina decrementa o contador waitgroup

	for i := 1; i <= 10; i++ {
		fmt.Println("Crescente:", i)
		time.Sleep(1 * time.Second)
	}
}

//função conta de 10 a 1, um número por segundo
func contarDecrescente(wg *sync.WaitGroup){
	defer wg.Done() // igual à de cima: avisa "terminei!" ao sair
	
	for i := 10; i >= 1; i-- {
		fmt.Println("Decrescente:", i)
		time.Sleep(1 * time.Second)
	}
}

//funcao main
func main(){
	var wg sync.WaitGroup //declara o WaitGroup
	
	wg.Add(2) // diz: "vou lançar 2 goroutines, espere por 2"

	go contarCrescente(&wg) // lança goroutine 1 (o "go" é o que faz ela rodar em paralelo)
	go contarDecrescente(&wg) // lança goroutine 2

	wg.Wait() // // bloqueia a main() aqui até o contador chegar a 0

	fmt.Println("Ambas as goroutines terminaram")
}