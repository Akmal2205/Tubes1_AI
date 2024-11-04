package algorithm

import (
	"fmt"
	"math/rand"
	"time"
)

func GeneticAlgorithm(populasi int, iterasi int) {
	start := time.Now()

	highestFitness := 0
	var bestCube [][][]int
	var stateAwal [][][]int
	fitnessArrayAvg := make([]float64, iterasi)
	fitnessArrayBest := make([]int, iterasi)
	kubusMagik := false
	indeksKubusMagik := -1

	ganjil := populasi%2 == 1

	cubesArray := make([][][][]int, populasi)
	for i := 0; i < populasi; i++ {
		cubesArray[i] = CreateCube()
		if i == 0 {
			stateAwal = cubesArray[i]
		}
	}

	for i := 0; i < iterasi; i++ {
		println("Iterasi ", i+1)
		if kubusMagik {
			break
		}
		// Ngisi fitness value untuk tiap cube
		totalFitness := 0
		fitnessValue := make([]int, populasi)
		fitnessAvg := 0.0
		fitnessLocalBest := 0
		for temp := 0; temp < populasi; temp++ {
			fitnessValue[temp] = 109 - EvaluateObjectiveFunction(&cubesArray[temp])
			fitnessAvg += float64(fitnessValue[temp])
			// cek fitness value terbagus global
			if fitnessValue[temp] > highestFitness {
				highestFitness = fitnessValue[temp]
				bestCube = cubesArray[temp]
			}
			// cek fitness value terbagus lokal
			if fitnessValue[temp] > fitnessLocalBest {
				fitnessLocalBest = fitnessValue[temp]
			}
			if fitnessValue[temp] == MAGIC_VALUE {
				kubusMagik = true
				indeksKubusMagik = temp
				break
			}
			totalFitness += fitnessValue[temp]
		}

		fitnessArrayBest[i] = fitnessLocalBest
		fitnessAvg = fitnessAvg / float64(populasi)
		fitnessArrayAvg[i] = fitnessAvg

		//bikin roulette wheel selection
		roulette := make([]float64, populasi)
		totalPersen := 0.0
		for temp := 0; temp < populasi; temp++ {
			bagi := float64(fitnessValue[temp]) / float64(totalFitness)
			totalPersen += float64(bagi)
			roulette[temp] = totalPersen
		}

		// menentukan parent
		parentCubes := make([][][][]int, populasi)
		for temp := 0; temp < populasi; temp++ {
			r := rand.Float64()
			for temp2 := 0; temp2 < populasi; temp2++ {
				if r <= float64(roulette[temp2]) {
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

				var parent2 []int
				if ganjil && j == populasi-1 { // Kalo ganjil, populasi terakhir nge-parent sama elemen sebelumnya (kedua terakhir)
					parent2 = StraightCube(parentCubes[j-1])

				} else {
					parent2 = StraightCube(parentCubes[j+1])
				}

				child1, child2 := pmxCrossover(parent1, parent2)

				SwapStraightRandom(&child1)
				SwapStraightRandom(&child2)

				// balikin jadi cube kotak
				cubesArray[j] = CubedCube(child1)
				if !ganjil { // jika ganjil ga ush tambahin anak ke-2
					cubesArray[j+1] = CubedCube(child2)
				}
			}
		}

	}
	if kubusMagik {
		print("ditemukan kubus magik\n")
		print(indeksKubusMagik)
		fmt.Println(bestCube)
	} else {
		print("tidak ditemukan kubus magik\n")
		fmt.Println("STATE AWAL")
		fmt.Println(stateAwal)
		fmt.Println("BEST CUBE")
		fmt.Println(bestCube)
	}
	fmt.Println("Highest Fitness : ", highestFitness)
	duration := time.Since(start)
	fmt.Println("Time Taken : ", duration)
	fmt.Println("Number of Iteration : ", iterasi)
	fmt.Println("Number of Population : ", populasi)
	fmt.Println("Array of maximum fitness per iteration : ")
	fmt.Println(fitnessArrayBest)
	fmt.Println("Array of average fitness per iteration : ")
	fmt.Println(fitnessArrayAvg)

	var step []CoordinatePair
	SaveMatrixXZ(stateAwal, step, "GA_StateAwal_23.txt")
	SaveMatrixXZ(bestCube, step, "GA_BestCube_23.txt")
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

	// Randomly milih 2 titik buat rentang crossover, diset point1 lebih kecil
	point1 := rand.Intn(length)
	point2 := rand.Intn(length)
	if point1 > point2 {
		point1, point2 = point2, point1
	}

	// Masang rentang crossover ke child masing2
	for i := point1; i <= point2; i++ {
		child1[i] = parent2[i]
		child2[i] = parent1[i]
	}

	// ngisi sisa elemen di child yang masih kosong
	for i := 0; i < length; i++ {
		if i >= point1 && i <= point2 {
			continue
		}
		// println(i)
		fillValue(child1, i, parent1)
		fillValue(child2, i, parent2)
	}

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
