package algorithm

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_ITERATION = 100

var final_objective_value int
var objective_value_list[MAX_ITERATION]int

func StochasticHC() {
	var current_objective_value, neighbor_objective_value, start_x, start_y, start_z, destination_x, destination_y, destination_z int
	var objective_value_list [MAX_ITERATION]int
	magic_cube := CreateCube() // [][][]int

	// saved initial state
	saved_magic_cube := CopyCube(magic_cube)
	var saved_steps []CoordinatePair

	fmt.Println("State awal kubus :")
	ShowMatrixXZ(magic_cube)

	start_time := time.Now()
	for i := 0; i < MAX_ITERATION; i++ {
		// fmt.Printf("Pengulangan ke-%d\n", i+1)
		current_objective_value = EvaluateObjectiveFunction(&magic_cube)

		// indeks pada current magic cube yang ingin di swap
		start_x = rand.Intn(MATRIX_N)
		start_y = rand.Intn(MATRIX_N)
		start_z = rand.Intn(MATRIX_N)

		// indeks tujuan swapping
		destination_x = rand.Intn(MATRIX_N)
		destination_y = rand.Intn(MATRIX_N)
		destination_z = rand.Intn(MATRIX_N)

		// dapetin neighbor dengan swap start dan destination
		Swap(&magic_cube, start_x, start_y, start_z, destination_x, destination_y, destination_z)
		neighbor_objective_value = EvaluateObjectiveFunction(&magic_cube)

		var cek_batal bool = false
		if neighbor_objective_value > current_objective_value {
			// dibalikin ke awal
			cek_batal = true
			Swap(&magic_cube, start_x, start_y, start_z, destination_x, destination_y, destination_z)
		}
        final_objective_value = EvaluateX(&magic_cube)+EvaluateY(&magic_cube)+EvaluateZ(&magic_cube)
        objective_value_list[i] = final_objective_value
		if final_objective_value == 0 {
			fmt.Println("ketemu euy gacor!!")
			break
		}

		final_objective_value = EvaluateObjectiveFunction(&magic_cube)
		objective_value_list[i] = final_objective_value

		if !cek_batal {
			// SAVED DATA FOR VISUALIZATION
			var new_data CoordinatePair

			var kordinat1 Coordinate3D
			kordinat1.X = start_x
			kordinat1.Y = start_y
			kordinat1.Z = start_z

			var kordinat2 Coordinate3D
			kordinat2.X = destination_x
			kordinat2.Y = destination_y
			kordinat2.Z = destination_z

			new_data.N = final_objective_value
			new_data.Point1 = kordinat1
			new_data.Point2 = kordinat2

			saved_steps = append(saved_steps, new_data)
		}
	}

	duration := time.Since(start_time)

	SaveMatrixXZ(saved_magic_cube, saved_steps, "edbert.txt")
	// print
	fmt.Println("State akhir kubus :")
	ShowMatrixXZ(magic_cube)

	fmt.Println("Maksimum iterasi :", MAX_ITERATION)
	fmt.Println("Nilai fungsi objektif terakhir:", final_objective_value)
	fmt.Println("Waktu eksekusi:", duration)

}
