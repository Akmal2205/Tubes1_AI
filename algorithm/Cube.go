package algorithm

import (
	"fmt"
	"math/rand"
)

const MAGIC_VALUE = 315
const MATRIX_N = 5

// Function to initiate random cube
func CreateCube() [][][]int {
	// rand.Seed(time.Now().UnixNano())

	// pola [x] [y] [z]
	var matrix [][][]int = make([][][]int, MATRIX_N)
	// temp := 0
	for i := 0; i < MATRIX_N; i++ {
		matrix[i] = make([][]int, MATRIX_N)
		for j := 0; j < MATRIX_N; j++ {
			matrix[i][j] = make([]int, MATRIX_N)
			for k := 0; k < MATRIX_N; k++ {
				// initialize random array
				matrix[i][j][k] = rand.Intn(125) + 1
				// matrix[i][j][k] = temp
				// temp++
			}
		}
	}

	// fmt.Println(matrix)
	// print("beres\n")

	return matrix
}

func ShowMatrixXZ(matrix [][][]int) {
	for j := 0; j < MATRIX_N; j++ {
		for i := MATRIX_N - 1; i >= 0; i-- {
			for k := 0; k < MATRIX_N; k++ {
				// print matrix
				fmt.Print(matrix[k][j][i])
				fmt.Print(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func ShowMatrixZX(matrix [][][]int) {
	for j := 0; j < MATRIX_N; j++ {
		for i := MATRIX_N - 1; i >= 0; i-- {
			for k := 0; k < MATRIX_N; k++ {
				// print matrix
				fmt.Print(matrix[i][j][k])
				fmt.Print(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func ShowMatrixYZ(matrix [][][]int) {
	for j := 0; j < MATRIX_N; j++ {
		for i := MATRIX_N - 1; i >= 0; i-- {
			for k := MATRIX_N - 1; k >= 0; k-- {
				// print matrix
				fmt.Print(matrix[i][k][j])
				fmt.Print(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func ShowMatrixXY(matrix [][][]int) {
	for j := 0; j < MATRIX_N; j++ {
		for i := MATRIX_N - 1; i >= 0; i-- {
			for k := 0; k < MATRIX_N; k++ {
				// print matrix
				fmt.Print(matrix[k][i][j])
				fmt.Print(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func CheckMagic(x int) bool {
	return x == MAGIC_VALUE
}

// RETURN : banyak X yang sum-nya bukan magic number
// TAMPILAN XY
func EvaluateX(matrix *[][][]int) int {
	var non_magic int = 0

	for j := 0; j < MATRIX_N; j++ {
		for i := MATRIX_N - 1; i >= 0; i-- {
			sum := 0
			for k := 0; k < MATRIX_N; k++ {
				// print matrix
				sum += (*matrix)[k][i][j]
			}
			// fmt.Println(sum)
			if !CheckMagic(sum) {
				non_magic++
			}
		}
	}
	return non_magic
}

// RETURN : banyak Y yang sum-nya bukan magic number.
// TAMPILAN YZ
func EvaluateY(matrix *[][][]int) int {
	var non_magic int = 0

	for j := 0; j < MATRIX_N; j++ {
		for i := MATRIX_N - 1; i >= 0; i-- {
			sum := 0
			for k := MATRIX_N - 1; k >= 0; k-- {
				sum += (*matrix)[j][k][i]
			}
			// fmt.Println(sum)
			if !CheckMagic(sum) {
				non_magic++
			}
		}
	}

	return non_magic
}

// RETURN : banyak Z yang sum-nya bukan magic number
// TAMPILAN ZX (transpose dari XZ)
func EvaluateZ(matrix *[][][]int) int {
	var non_magic int = 0

	for j := 0; j < MATRIX_N; j++ {
		for i := MATRIX_N - 1; i >= 0; i-- {
			sum := 0
			for k := 0; k < MATRIX_N; k++ {
				// print matrix
				sum += (*matrix)[i][j][k]
			}
			// fmt.Println(sum)
			if !CheckMagic(sum) {
				non_magic++
			}
		}
	}

	return non_magic
}

// PROCEDURE : swap values between 2 coordiantes
// ASSUME [x] [y] [z]
func Swap(matrix *[][][]int, x1, y1, z1, x2, y2, z2 int) {
	var temp int = (*matrix)[x1][y1][z1]
	(*matrix)[x1][y1][z1] = (*matrix)[x2][y2][z2]
	(*matrix)[x2][y2][z2] = temp
}
