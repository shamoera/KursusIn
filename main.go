package main

import (
	"fmt"
)

type Peserta struct {
	ID      int
	Nama    string
	Tanggal string
	Bidang  string
	Status  bool
}

var dataPeserta []Peserta

func tambahPeserta() {
	var p Peserta

	fmt.Print("Masukkan ID Pendaftaran: ")
	fmt.Scan(&p.ID)

	fmt.Print("Masukkan Nama Peserta: ")
	fmt.Scan(&p.Nama)

	fmt.Print("Masukkan Tanggal Pendaftaran: ")
	fmt.Scan(&p.Tanggal)

	fmt.Print("Masukkan Bidang Minat: ")
	fmt.Scan(&p.Bidang)

	p.Status = true

	dataPeserta = append(dataPeserta, p)

	fmt.Println("Data peserta berhasil ditambahkan")
}

func tampilPeserta() {
	if len(dataPeserta) == 0 {
		fmt.Println("Belum ada data peserta")
		return
	}

	fmt.Println("\n===== DATA PESERTA =====")

	for i := 0; i < len(dataPeserta); i++ {
		fmt.Println("ID      :", dataPeserta[i].ID)
		fmt.Println("Nama    :", dataPeserta[i].Nama)
		fmt.Println("Tanggal :", dataPeserta[i].Tanggal)
		fmt.Println("Bidang  :", dataPeserta[i].Bidang)
		fmt.Println("--------------------------")
	}
}

func sequentialSearch() {
	var nama string
	found := false

	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&nama)

	for i := 0; i < len(dataPeserta); i++ {
		if dataPeserta[i].Nama == nama {
			fmt.Println("Data ditemukan")
			fmt.Println("ID :", dataPeserta[i].ID)
			fmt.Println("Bidang :", dataPeserta[i].Bidang)
			found = true
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

func selectionSort() {
	n := len(dataPeserta)

	for i := 0; i < n-1; i++ {
		min := i

		for j := i + 1; j < n; j++ {
			if dataPeserta[j].ID < dataPeserta[min].ID {
				min = j
			}
		}

		dataPeserta[i], dataPeserta[min] = dataPeserta[min], dataPeserta[i]
	}

	fmt.Println("Data berhasil diurutkan berdasarkan ID")
}

func statistikPeserta() {
	total := 0

	for i := 0; i < len(dataPeserta); i++ {
		if dataPeserta[i].Status {
			total++
		}
	}

	fmt.Println("Total peserta aktif:", total)
}

func main() {
	var pilihan int

	for {
		fmt.Println("\n===== SISTEM KURSUSIN =====")
		fmt.Println("1. Tambah Peserta")
		fmt.Println("2. Tampilkan Peserta")
		fmt.Println("3. Cari Peserta")
		fmt.Println("4. Urutkan Peserta")
		fmt.Println("5. Statistik Peserta")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			tambahPeserta()

		case 2:
			tampilPeserta()

		case 3:
			sequentialSearch()

		case 4:
			selectionSort()

		case 5:
			statistikPeserta()

		case 0:
			fmt.Println("Program selesai")
			return

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}
