package algorithm

import (
	"fmt"
	"math/rand"
	"time"
)

const cube_size = 5

func StochasticHC() {
	var magic_cube[cube_size][cube_size][cube_size]int
	var current_objective_value, neighbor_objective_value, start_x,start_y, start_z, destination_x, destination_y, destination_z, max_iteration  int
	var start_index[3]int
	var destination_index[3]int

	//TODO: inisiasi magic_cube
	initializeMagicCube(&magic_cube)


	max_iteration = 10
	for i:=0; i<max_iteration; i++ {
		fmt.Printf("Pengulangan ke-%d\n", i+1)
		current_objective_value = ObjectiveFunction(magic_cube)

		// indeks pada current magic cube yang ingin di swap
		start_x = rand.Intn(cube_size)
		start_y = rand.Intn(cube_size)
		start_z = rand.Intn(cube_size)
	
		start_index[0] = start_x
		start_index[1] = start_y
		start_index[2] = start_z

		// indeks tujuan swapping
		destination_x = rand.Intn(cube_size)
		destination_y = rand.Intn(cube_size)
		destination_z = rand.Intn(cube_size)
	
		destination_index[0] = destination_x
		destination_index[1] = destination_y
		destination_index[2] = destination_z
		
		// dapetin neighbor dengan swap start dan destination
		Swap(&magic_cube, start_index, destination_index)
		neighbor_objective_value = ObjectiveFunction(magic_cube)

		if neighbor_objective_value <= current_objective_value {
			// dibalikin ke awal
			Swap(&magic_cube, destination_index, start_index)
		}
        printMagicCube(magic_cube)
        fmt.Println("objective function value : ", ObjectiveFunction(magic_cube))
	}
}

func Swap(magic_cube *[cube_size][cube_size][cube_size]int, start, destination [3]int) {
    temp := magic_cube[start[0]][start[1]][start[2]]
    magic_cube[start[0]][start[1]][start[2]] = magic_cube[destination[0]][destination[1]][destination[2]]
    magic_cube[destination[0]][destination[1]][destination[2]] = temp
}

func initializeMagicCube(magic_cube *[cube_size][cube_size][cube_size]int) {
    // Membuat daftar angka dari 1 hingga 125
    numbers := make([]int, cube_size*cube_size*cube_size)
    for i := 0; i < cube_size*cube_size*cube_size; i++ {
        numbers[i] = i + 1
    }

    // Mengacak daftar angka untuk mendapatkan urutan acak tanpa pengulangan
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(numbers), func(i, j int) {
        numbers[i], numbers[j] = numbers[j], numbers[i]
    })

    // Mengisi magic_cube dengan angka acak unik
    index := 0
    for i := 0; i < cube_size; i++ {
        for j := 0; j < cube_size; j++ {
            for k := 0; k < cube_size; k++ {
                magic_cube[i][j][k] = numbers[index]
                index++
            }
        }
    }
}

func ObjectiveFunction(magic_cube [cube_size][cube_size][cube_size]int) int {
    // Hitung magic constant untuk ukuran cube_size
    magicConstant := cube_size * (cube_size*cube_size*cube_size + 1) / 2
    totalDeviation := 0

    // Periksa setiap baris, kolom, dan diagonal pada tiga sumbu (x, y, z)
    for i := 0; i < cube_size; i++ {
        // Hitung jumlah pada setiap baris, kolom, dan lapisan
        sumX, sumY, sumZ := 0, 0, 0
        for j := 0; j < cube_size; j++ {
            for k := 0; k < cube_size; k++ {
                // Sum baris sejajar sumbu x, y, dan z
                sumX += magic_cube[i][j][k]
                sumY += magic_cube[j][i][k]
                sumZ += magic_cube[j][k][i]
            }
        }
        // Tambahkan selisih dari magic constant ke total deviasi
        totalDeviation += abs(sumX - magicConstant)
        totalDeviation += abs(sumY - magicConstant)
        totalDeviation += abs(sumZ - magicConstant)
    }

    // Periksa diagonal utama di tiap layer untuk mendekati magic constant
    for i := 0; i < cube_size; i++ {
        diag1XY, diag2XY := 0, 0
        diag1XZ, diag2XZ := 0, 0
        diag1YZ, diag2YZ := 0, 0
        for j := 0; j < cube_size; j++ {
            diag1XY += magic_cube[i][j][j]
            diag2XY += magic_cube[i][j][cube_size-j-1]
            diag1XZ += magic_cube[j][i][j]
            diag2XZ += magic_cube[cube_size-j-1][i][j]
            diag1YZ += magic_cube[j][j][i]
            diag2YZ += magic_cube[cube_size-j-1][j][i]
        }
        // Tambahkan selisih dari magic constant ke total deviasi
        totalDeviation += abs(diag1XY - magicConstant)
        totalDeviation += abs(diag2XY - magicConstant)
        totalDeviation += abs(diag1XZ - magicConstant)
        totalDeviation += abs(diag2XZ - magicConstant)
        totalDeviation += abs(diag1YZ - magicConstant)
        totalDeviation += abs(diag2YZ - magicConstant)
    }

    return totalDeviation
}

// Fungsi helper untuk menghitung nilai absolut
func abs(value int) int {
    if value < 0 {
        return -value
    }
    return value
}

func printMagicCube(magic_cube [cube_size][cube_size][cube_size]int) {
    fmt.Println("Isi magic_cube saat ini:")
    for i := 0; i < cube_size; i++ {
        fmt.Printf("Layer %d:\n", i)
        for j := 0; j < cube_size; j++ {
            for k := 0; k < cube_size; k++ {
                fmt.Printf("%4d ", magic_cube[i][j][k])
            }
            fmt.Println() // Ganti baris setelah setiap baris di layer
        }
        fmt.Println() // Ganti baris setelah setiap layer
    }
    fmt.Println() // Baris kosong untuk pemisah antara iterasi
}