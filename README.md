# Pencarian Solusi Diagonal Magic Cube dengan Local Search
> *Repository* ini dibuat untuk memenuhi Tugas Besar 1 Mata kuliah Intelegensi Artifisial IF3170.

## Deskripsi Singkat Program
Program ini mencoba menemukan susunan angka di dalam kubus ajaib (3D matrix) yang memenuhi kondisi magic cube, di mana jumlah angka pada tiap diagonal (baik pada sumbu X, Y, maupun Z) mencapai target jumlah tertentu. Algoritma local search yang digunakan adalah Hill Climbing, Simulated Annealing, dan Genetic Algorithm.

## Requirements
- Docker

## Cara Menjalankan Program
1. Pastikan Java sudah terpasang di perangkat anda. Status pemasangan dapat diperiksa dengan menjalankan *command* `java --version` pada *command prompt*.
2. *Clone repository* dengan *command* berikut
```
git clone https://github.com/fabianradenta/Tucil3_13522105.git
``` 
3. Masuk ke *directory* src. Pastikan mengganti `/path/to/src` dengan *path* yang benar.
```
cd /path/to/src
```
4. *Compile* dan *run* program dengan *command*
```
Javac -d ../bin/ Main.java
Java -cp ../bin/ Main
```
5. Program akan meminta masukan *start word*, *end word*, dan metode algoritma yang akan digunakan. Pastikan masukan yang diterima program benar.
6. Jika masukan yang diterima sudah benar, program akan melakukan pencarian solusi permainan Word Ladder. Jika solusi ditemukan, program akan menampilkan *path* solusi, jumlah *node* yang telah dikunjungi, dan waktu eksekusi program. Namun, jika solusi tidak ditemukan, program akan menampilkan pesan bahwa solusi tidak ditemukan dan menampilkan waktu eksekusi program.

## Pembagian Tugas
| Task                          | NIM                                      |
|-------------------------------|------------------------------------------|
| Stochastic Hill Climbing Algorithm | 13522105                          |
| Simulated Annealing Algorithm  | 13522076                               |
| Genetic Algorithm              | 13522049                               |
| Cube Class                     | 13522039                               |
| Visualization [Bonus]          | 13522039                               |
| Dokumentasi Laporan            | 13522039, 13522049, 13522076, 13522105 |
