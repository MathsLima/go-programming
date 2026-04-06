/*
Objetivo: identificar uma condição de corrida e usar o detector de corrida do Go.

Tarefas:
1 - Execute o programa 5 vezes com go run ex3.go . O resultado é sempre 200000? Anote os valores
obtidos.

2 - Execute com o detector de corrida: go run -race ex3.go . O que a saída indica?

3 - Corrija o programa adicionando um sync.Mutex para proteger o acesso ao contador. Execute
novamente (com e sem -race ) e verifique que o resultado é sempre 200000.

4 - Compare o tempo de execução com e sem mutex:
	time go run ex3.go # sem mutex (resultado incorreto)
	time go run ex3_mutex.go # com mutex (resultado correto)
	- O mutex adiciona overhead? Por quê?

Resultados: 
1 - Contador: 101316,Contador: 117684, Contador: 119855, Contador: 109247, Contador: 140876
2 - Found 2 data race(s)
3 - Contador: 200000 (esperado: 200000) - sim :  
	sem mutex, as goroutines correm livres sem esperar ninguém. 
	Com mutex, quando uma está dentro do Lock(), a outra fica parada esperando a vez — isso é overhead real.
4 - 
*/

package main

import (
	"fmt"
	"sync"
)

var contador int

func incrementar(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
	contador++
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go incrementar(100000, &wg)
	go incrementar(100000, &wg)
	wg.Wait()
	fmt.Printf("Contador: %d (esperado: 200000)\n", contador)
}

/*
Com Mutex:

package main

import (
    "fmt"
    "sync"
)

var contador int
var mu sync.Mutex // o cadeado

func incrementar(n int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < n; i++ {
        mu.Lock()   // tranca: só eu posso mexer agora
        contador++
        mu.Unlock() // destrava: próxima pode entrar
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go incrementar(100000, &wg)
    go incrementar(100000, &wg)
    wg.Wait()
    fmt.Printf("Contador: %d (esperado: 200000)\n", contador)
}
*/