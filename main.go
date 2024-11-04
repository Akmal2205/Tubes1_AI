package main

import (
	"fmt"
	"project-root/algorithm"
)

func main() {
	fmt.Println("Vanson")

	algorithm.GeneticAlgorithm(100, 250)

	// edbert
	algorithm.StochasticHC()

	// fmt.Println(c)
	// algorithm.Ed1()
	// algorithm.Ed2()
	// algorithm.Ed3()

	// algorithm.MainSimulatedAnnealing()
	// fmt.Println(algorithm.Probabilistic(1.0, 2.0))

}
