package algorithm

import (
	"fmt"
	"math/rand"
	"time"
)

const MAGIC_VALUE = 315
const MATRIX_N = 5

// Function to initiate random cube
func CreateCube() [][][]int {
	rand.Seed(time.Now().UnixNano())

	// Generate numbers 1 to 125
	values := make([]int, 125)
	for i := 0; i < 125; i++ {
		values[i] = i + 1
	}

	// Shuffle the numbers
	rand.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })

	// pola [x] [y] [z]
	var matrix [][][]int = make([][][]int, MATRIX_N)
	index := 0
	// temp := 0
	for i := 0; i < MATRIX_N; i++ {
		matrix[i] = make([][]int, MATRIX_N)
		for j := 0; j < MATRIX_N; j++ {
			matrix[i][j] = make([]int, MATRIX_N)
			for k := 0; k < MATRIX_N; k++ {
				// initialize random array
				// matrix[i][j][k] = values[index]
				matrix[i][j][k] = index
				index++
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
				fmt.Print(matrix[j][k][i])
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

func EvaluateDiagonalBidang(matrix *[][][]int) int {
	var non_magic = 0

	// Diagonal in the XZ plane for each Y level
	for y := 0; y < MATRIX_N; y++ {
		sum1, sum2 := 0, 0
		for x, z := 0, 0; x < MATRIX_N; x, z = x+1, z+1 {
			sum1 += (*matrix)[x][y][z]
			sum2 += (*matrix)[MATRIX_N-1-x][y][z]
		}
		// fmt.Println(sum1, sum2)
		if !CheckMagic(sum1) {
			non_magic++
		}
		if !CheckMagic(sum2) {
			non_magic++
		}
	}

	// Diagonal in the YZ plane for each X level
	for x := 0; x < MATRIX_N; x++ {
		sum1, sum2 := 0, 0
		for y, z := 0, 0; y < MATRIX_N; y, z = y+1, z+1 {
			// fmt.Println(x, y, z)
			// fmt.Println((*matrix)[x][y][z])
			sum1 += (*matrix)[x][y][z]
			sum2 += (*matrix)[x][MATRIX_N-1-y][z]
		}
		// fmt.Println(sum1, sum2)
		if !CheckMagic(sum1) {
			non_magic++
		}
		if !CheckMagic(sum2) {
			non_magic++
		}
	}

	// Diagonal in the XY plane for each Z level
	for z := 0; z < MATRIX_N; z++ {
		sum1, sum2 := 0, 0
		for x, y := 0, 0; x < MATRIX_N; x, y = x+1, y+1 {
			sum1 += (*matrix)[x][y][z]
			sum2 += (*matrix)[MATRIX_N-1-x][y][z]
		}
		// fmt.Println(sum1, sum2)
		if !CheckMagic(sum1) {
			non_magic++
		}
		if !CheckMagic(sum2) {
			non_magic++
		}
	}

	return non_magic
}

func EvaluateDiagonalRuang(matrix *[][][]int) int {
	var non_magic = 0

	// A to G diagonal (0,0,0) to (4,4,4)
	sum := 0
	for i := 0; i < MATRIX_N; i++ {
		sum += (*matrix)[i][i][i]
	}
	fmt.Println("A to G:", sum)
	if !CheckMagic(sum) {
		non_magic++
	}

	// B to H diagonal (0,0,4) to (4,4,0)
	sum = 0
	for i := 0; i < MATRIX_N; i++ {
		sum += (*matrix)[i][i][MATRIX_N-1-i]
	}
	fmt.Println("B to H:", sum)
	if !CheckMagic(sum) {
		non_magic++
	}

	// E to C diagonal (4,0,0) to (0,4,4)
	sum = 0
	for i := 0; i < MATRIX_N; i++ {
		sum += (*matrix)[MATRIX_N-1-i][i][i]
	}
	fmt.Println("E to C:", sum)
	if !CheckMagic(sum) {
		non_magic++
	}

	// F to D diagonal (4,0,4) to (0,4,0)
	sum = 0
	for i := 0; i < MATRIX_N; i++ {
		sum += (*matrix)[MATRIX_N-1-i][i][MATRIX_N-1-i]
	}
	fmt.Println("F to D:", sum)
	if !CheckMagic(sum) {
		non_magic++
	}

	return non_magic
}

// OBJECTIVE FUNCTION = state value.
// number of rows / col / diagonal that is not equal to magic number
func EvaluateObjectiveFunction(matrix *[][][]int) int {
	return EvaluateX(matrix) + EvaluateY(matrix) + EvaluateZ(matrix) + EvaluateDiagonalBidang(matrix) + EvaluateDiagonalRuang(matrix)
}

// PROCEDURE : swap values between 2 coordiantes
// ASSUME [x] [y] [z]
func Swap(matrix *[][][]int, x1, y1, z1, x2, y2, z2 int) {
	var temp int = (*matrix)[x1][y1][z1]
	(*matrix)[x1][y1][z1] = (*matrix)[x2][y2][z2]
	(*matrix)[x2][y2][z2] = temp
}
