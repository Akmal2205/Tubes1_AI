package algorithm

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
Probabilistic calculates a probability value based on a given energy change (delta_E)
and temperature (T), using the formula:

	p = e^(delta_E / T)

Parameters:

	delta_E: Represents the change in energy, which can influence the probability.
	T: The temperature, which affects the spread or likelihood of the probabilistic outcome.

Returns:

	A probability value between 0 and 1.
*/
func Probabilistic(delta_E int, T float64) float64 {
	euler := 2.71828
	power := float64(float64(delta_E) / T)
	return math.Pow(euler, power)
}

/*
A procedure to simulate T decrease as time passes (t gets bigger), using the formula:

	T = T * 0.5, decreases by a factor of 0.5 creating an exponential decreasing process

Parameters:

	cube: Reference of the cube, is a 3 dimensional array of integer

Returns:

	none. (Procedure)
*/
func TemperatureDecrease(T *float64, t int) {
	if *T > 1 {
		*T = *T * 0.99
	} else {
		*T = *T - (5 * math.Pow(10.0, -7.0))
	}
}

/*
The main interface of the simulated annealing function

Parameters:

	cube: Reference of the cube, is a 3 dimensional array of integer with 5 x 5 dimension

Returns:

	none. (Procedure)
*/
func MainSimulatedAnnealing() { // main interface for the algorithm
	T0 := math.Pow(10.0, 200.0)
	t := 1
	magic_cube := CreateCube()

	fmt.Println("State awal kubus: ")
	ShowMatrixXZ(magic_cube)

	start_time := time.Now()
	for T0 > 5e-324 { // minimum value represented in float64
		TemperatureDecrease(&T0, t)

		current_objective_value := EvaluateObjectiveFunction(&magic_cube)

		// indeks pada current magic cube yang ingin di swap
		start_x := rand.Intn(MATRIX_N)
		start_y := rand.Intn(MATRIX_N)
		start_z := rand.Intn(MATRIX_N)

		// indeks tujuan swapping
		destination_x := rand.Intn(MATRIX_N)
		destination_y := rand.Intn(MATRIX_N)
		destination_z := rand.Intn(MATRIX_N)

		// dapetin neighbor dengan swap start dan destination
		Swap(&magic_cube, start_x, start_y, start_z, destination_x, destination_y, destination_z)
		neighbor_objective_value := EvaluateObjectiveFunction(&magic_cube)

		if neighbor_objective_value == 0 {
			fmt.Println("Ketemu cuy")
			break
		}

		// kalkulasi delta_E
		delta_E := current_objective_value - neighbor_objective_value

		if delta_E <= 0 {
			if Probabilistic(delta_E, T0) < 0.3 {
				// kondisi tidak diambil
				Swap(&magic_cube, destination_x, destination_y, destination_z, start_x, start_y, start_z)
			}
		}

		// fmt.Println("T, delta_E:", T0, delta_E)
		fmt.Println("T :", T0)
		fmt.Println("p: ", Probabilistic(delta_E, T0))
		fmt.Println("iteration:", t)
		fmt.Println("Objective Value:", current_objective_value)
		fmt.Println("")
		t++
	}
	duration := time.Since(start_time).Minutes()

	fmt.Println("time taken: ", duration, " minute(s)")

	ShowMatrixXZ(magic_cube)
	
	fmt.Println("iterations: ", t)
}
