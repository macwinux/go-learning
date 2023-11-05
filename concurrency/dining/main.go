package main

import (
	"fmt"
	"sync"
	"time"
)

// Problema de los filosofos

// Philosopher is a struc which stores information about a philosopher

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// philosofers es una lista con todos los filosofos

var philosophers = []Philosopher{
	{name: "Platon", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristoteles", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// define unas pocas variables
var hunger = 3                  //cuantas veces un filosofo come
var eatTime = 1 * time.Second   //cuanto tarda en comer
var thinkTime = 3 * time.Second // cuanto tarda un filosofo en pensar
var sleepTime = 1 * time.Second // cuanto hay que esperar para imprimir por pantalla

func main() {

	fmt.Println("Problema de la cena de los filosofos")
	fmt.Println("------------------------------------")
	fmt.Println("The table is empty.")

	dine()

	fmt.Println("La mesa esta vacia")

}

func dine() {

	// wg es el WaitGroup que mantiene la trazabilidad de cuantos filosofos hay todavia en la mesa
	// cuando llega a cero, todo el mundo a terminado de comer y se ha ido. Añadimos 5 a est
	// wait group.

	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// queremos a todos sentados antes de empezar a comer, creamos este grupo para eso
	// y lo inicializamos a 5

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks es un map con los 5 tenedores. Forks son asignados usando las variables leftFork y rightFork
	// de un filosofo.
	// type. Cada tenedor, entonces, puede ser encontrado usando el indice (int), cada tenedor tiene un unico mutex
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// comenzamos la cena iterando sobre nuestros filosofos

	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()

}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// Sentamos a los filosofos a la mesa
	fmt.Printf("%s esta sentado en la mesa. \n", philosopher.name)
	seated.Done()

	// Esperamos hasta que todos los filosofos esten sentados
	seated.Wait()
	// comer tres veces
	for i := hunger; i > 0; i-- {

		// Para lidiar con una condicion de carrera logica en la que todos los filosofos cogen los tenedores a la vez
		// debemos crear un bucle if que modifique el orden en el que se cogen los tenedores si:
		if philosopher.leftFork > philosopher.rightFork {
			// bloqueamos los tenedores
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s coge el tenedor derecho\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s coge el tenedor izquierdo\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s coge el tenedor izquierdo\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s coge el tenedor derecho\n", philosopher.name)
		}

		fmt.Printf("\t%s tiene ambos tenedores y esta comiendo.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s esta pensando", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s deja los tenedores.\n", philosopher.name)

	}

	fmt.Println(philosopher.name, " esta satisfecho.")
	fmt.Println(philosopher.name, " dejo la mesa.")
}
