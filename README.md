# Pencarian Solusi Diagonal Magic Cube dengan Local Search
> *Repository* ini dibuat untuk memenuhi Tugas Besar 1 Mata kuliah Intelegensi Artifisial IF3170.

## Deskripsi Singkat Program
Program ini mencoba menemukan susunan angka di dalam kubus ajaib (3D matrix) yang memenuhi kondisi magic cube, di mana jumlah angka pada tiap diagonal (baik pada sumbu X, Y, maupun Z) mencapai target jumlah tertentu. Algoritma local search yang digunakan adalah Hill Climbing, Simulated Annealing, dan Genetic Algorithm.

## Requirements
- Docker

## Menjalankan Program
1. Aktifkan Docker yang sudah terpasang di perangkat anda.
2. Pada terminal, lakukan *clone repository* dengan *command* berikut.
```
git clone https://github.com/Akmal2205/Tubes1_AI.git
``` 
3. Masuk ke *directory*.
```
cd Tubes1_AI
```
4. Jalankan program dengan Docker Compose
```
docker-compose up
```
5. Pada Docker Desktop, klik container dengan nama golang_app yang sedang berjalan lalu pilih opsi Exec untuk membuka terminal di dalam container tersebut.
6. Program dapat dijalankan pada terminal tersebut dengan command
```
./run
```
atau
```
go run main.go
```

## Menghentikan Program
Setelah selesai menggunakan program, Anda dapat menghentikan semua container dengan menekan `Ctrl + C` di terminal tempat `docker-compose up` dijalankan, atau dengan perintah:
```
docker-compose down
```
## Pembagian Tugas
| Task                          | NIM                                      |
|-------------------------------|------------------------------------------|
| Stochastic Hill Climbing Algorithm | 13522105                          |
| Simulated Annealing Algorithm  | 13522076                               |
| Genetic Algorithm              | 13522049                               |
| Cube Class                     | 13522039                               |
| Visualization [Bonus]          | 13522039                               |
| Dokumentasi Laporan            | 13522039, 13522049, 13522076, 13522105 |

## Author

1. [Edbert Eddyson Gunawan - 13522039](github.com/WazeAzure)
2. [Vanson Kurnialim - 13522049](github.com/VansonK)
3. [Muhammad Syarafi Akmal - 13522076](github.com/Akmal2205)
4. [Fabian Radenta Bangun - 13522105](github.com/fabianradenta)