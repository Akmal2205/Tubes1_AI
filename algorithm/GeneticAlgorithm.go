package algorithm

import (
	"fmt"
	"math/rand"
)

func GeneticAlgorithm(populasi int, iterasi int) {
	fmt.Println("dari GA")

	kubusMagik := false
	indeksKubusMagik := -1

	ganjil := populasi%2 == 1

	cubesArray := make([][][][]int, populasi)
	for i := 0; i < populasi; i++ {
		cubesArray[i] = CreateCube()
		// fmt.Println(cubesArray[i])
	}

	for i := 0; i < iterasi; i++ {
		println("Iterasi ", i+1)
		if kubusMagik {
			break
		}
		// Ngisi fitness value untuk tiap cube
		totalFitness := 0
		fitnessValue := make([]int, populasi)
		for temp := 0; temp < populasi; temp++ {
			fitnessValue[temp] = EvaluateX(&cubesArray[temp]) + EvaluateY(&cubesArray[temp]) + EvaluateZ(&cubesArray[temp])
			if fitnessValue[temp] == 315 {
				kubusMagik = true
				indeksKubusMagik = temp
				break
			}
			println(fitnessValue[temp])
			totalFitness += fitnessValue[temp]
		}

		//bikin roulette wheel selection
		roulette := make([]float64, populasi)
		totalPersen := 0.0
		for temp := 0; temp < populasi; temp++ {
			totalPersen += float64(fitnessValue[temp]) / float64(totalFitness)
			roulette[temp] = totalPersen
		}

		// tentuin parent
		// rand.Seed(time.Now().UnixNano())
		parentCubes := make([][][][]int, populasi)
		for temp := 0; temp < populasi; temp++ {
			r := rand.Float64()
			for temp2 := 0; temp2 < populasi; temp2++ {
				if r <= float64(fitnessValue[temp2]) {
					parentCubes[temp] = cubesArray[temp2]
					break
				}
			}
		}

		// PMX Crossover
		for j := 0; j < populasi; j += 2 {
			if j+1 < populasi {
				// Membuat 3D array jadi 1D
				parent1 := StraightCube(parentCubes[j])
				// print("Parent 1")
				// fmt.Println(parent1)
				var parent2 []int
				if ganjil { // Kalo ganjil, populasi terakhir nge-parent sama elemen sebelumnya (kedua terakhir)
					parent2 = StraightCube(parentCubes[j-1])

				} else {
					parent2 = StraightCube(parentCubes[j+1])
				}

				child1, child2 := pmxCrossover(parent1, parent2)

				// fmt.Println("Child1", child1)

				// balikin jadi cube kotak
				cubesArray[j] = CubedCube(child1)
				if !ganjil { // jika ganjil ga ush tambahin anak ke-2
					cubesArray[j+1] = CubedCube(child2)
				}
			}
		}

	}
	if kubusMagik {
		print("ditemukan kubus magik")
		print(indeksKubusMagik)
	} else {
		print("tidak ditemukan kubus magik")
	}

}

// terima straightedCube terus balikin tuple 2 biji child
func pmxCrossover(parent1, parent2 []int) ([]int, []int) {
	length := len(parent1)
	child1 := make([]int, length)
	child2 := make([]int, length)

	// inisiasi array kosong
	for i := 0; i < length; i++ {
		child1[i] = -1
		child2[i] = -1
	}

	fmt.Println("child1", child1)

	// Randomly milih 2 titik buat rentang crossover, diset point1 lebih kecil
	point1 := rand.Intn(length)
	point2 := rand.Intn(length)
	if point1 > point2 {
		point1, point2 = point2, point1
	}

	println("point1, point2", point1, point2)

	// Masang rentang crossover ke child masing2
	for i := point1; i <= point2; i++ {
		child1[i] = parent2[i]
		child2[i] = parent1[i]
	}

	fmt.Println("Child1", child1)

	// ngisi sisa elemen di child yang masih kosong
	println("LEngth", length)
	for i := 0; i < length; i++ {
		if i >= point1 && i <= point2 {
			println("CONTINUE")
			continue
		}
		println(i)
		fillValue(child1, i, parent1)
		fillValue(child2, i, parent2)
	}

	fmt.Println("stlh fill", child1)

	return child1, child2
}

// Helper function to fill value in child for uniqueness
func fillValue(child []int, idx int, parent []int) {
	for _, val := range parent {
		if !contains(child, val) {
			child[idx] = val
			break
		}
	}
}

// Checks if a value exists in a slice
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func StraightCube(cube [][][]int) []int {
	straightedCube := make([]int, 125)
	index := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < 5; z++ {
				straightedCube[index] = cube[x][y][z]
				index++
			}
		}
	}
	return straightedCube
}

func CubedCube(straightCube []int) [][][]int {
	cubedCube := make([][][]int, 5)
	for x := 0; x < 5; x++ {
		cubedCube[x] = make([][]int, 5)
		for y := 0; y < 5; y++ {
			cubedCube[x][y] = make([]int, 5)
		}
	}
	index := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < 5; z++ {
				cubedCube[x][y][z] = straightCube[index]
				index++
			}
		}
	}
	return cubedCube
}
