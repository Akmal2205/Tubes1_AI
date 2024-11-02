package algorithm

import (
	"fmt"
	"math/rand"
)

func StochasticHC() {
	var current_objective_value, neighbor_objective_value, start_x,start_y, start_z, destination_x, destination_y, destination_z, max_iteration  int

    magic_cube := CreateCube()

	max_iteration = 5
    fmt.Println("Maksimal iterasi :", max_iteration)
	for i:=0; i<max_iteration; i++ {
		fmt.Printf("Pengulangan ke-%d\n", i+1)
		current_objective_value = EvaluateX(&magic_cube) + EvaluateY(&magic_cube) + EvaluateZ(&magic_cube)

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
		neighbor_objective_value = EvaluateX(&magic_cube) + EvaluateY(&magic_cube) + EvaluateZ(&magic_cube)

		if neighbor_objective_value <= current_objective_value {
			// dibalikin ke awal
			Swap(&magic_cube, destination_x, destination_y, destination_z, start_x, start_y, start_z)
		}
        fmt.Println("XZ : ")
        ShowMatrixXZ(magic_cube)
        fmt.Println("ZX : ")
        ShowMatrixZX(magic_cube)
        fmt.Println("YZ : ")
        ShowMatrixYZ(magic_cube)
        fmt.Println("XY : ")
        ShowMatrixXY(magic_cube)
	}
}