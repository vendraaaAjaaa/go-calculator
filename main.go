package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	// 1. InstallPeta (0: Kosong, 1: Peta, 2: Makanan)
	maze := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 2, 2, 2, 1, 2, 2, 2, 2, 1},
		{1, 2, 1, 2, 1, 2, 1, 1, 2, 1},
		{1, 2, 1, 2, 2, 2, 2, 2, 2, 1},
		{1, 2, 1, 1, 1, 1, 1, 2, 1, 1},
		{1, 2, 2, 2, 2, 2, 2, 2, 2, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	// Posisi awal si Pacman (Baris, Kolom)
	playerRow, playerCol := 1, 1
	score := 0
	gameRunning := true

	for gameRunning {
		// bersihim layar terminal
		clearScreen()

		fmt.Println("=== PACMAN GO (CLI VERSION) ===")
		fmt.Println("Skor:", score)
		fmt.Println("Kontrol: w (atas), s (bawah), a (kiri), d (kanan), q (keluar)")

		// 2. Perulangan Bersarang (Nested Loop) untuk ngegambar petanya
		for r := 0; r < len(maze); r++ {
			for c := 0; c < len(maze[r]); c++ {
				// Percabangan buat nentiin apa yang digambar
				if r == playerRow && c == playerCol {
					fmt.Print("C ") // Pacman
				} else if maze[r][c] == 1 {
					fmt.Print("# ") // tembok
				} else if maze[r][c] == 2 {
					fmt.Print(". ") // Makanan
				} else {
					fmt.Print("  ") // Jalan kosong
				}
			}
			fmt.Println()
		}

		// Input dari user
		var move string
		fmt.Print("Gerak: ")
		fmt.Scan(&move)

		// Simpan posisi sebelumnya
		newRow, newCol := playerRow, playerCol

		// 3. Percabangan buat Logika Input
		if move == "w" {
			newRow--
		} else if move == "s" {
			newRow++
		} else if move == "a" {
			newCol--
		} else if move == "d" {
			newCol++
		} else if move == "q" {
			gameRunning = false
			fmt.Println("Game Berhenti. Skor Akhir:", score)
			break
		}

		// 4. Percabangan buat Cek Tabrakan Dinding
		if maze[newRow][newCol] != 1 {
			playerRow = newRow
			playerCol = newCol
		} else {
			fmt.Println("Ups! Ada dinding!")
		}

		// 5. Percabangan buat Cek Makan Titik
		if maze[playerRow][playerCol] == 2 {
			score += 10
			maze[playerRow][playerCol] = 0 // Ubah jadi kosong pas udah dimakan
		}

		// Cek Menang (kalo skor sampe jumlah tertentu)
		if score == 280 { // Angka ini tergantung jumlah titik di peta
			fmt.Println("Selamat! Kamu Menang!")
			gameRunning = false
		}
	}
}

// Fungsi pembantu buat bersihin terminal
func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
