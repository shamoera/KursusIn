package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

//   KURSUS-IN: Sistem Pendaftaran Kursus Online Terpadu
//   Tugas Besar Algoritma Pemrograman 2
//   Fakultas Informatika - Telkom University

const MAXDATA = 100

type Peserta struct {
	id          int
	nama        string
	email       string
	noTelp      string
	tglDaftar   string
	bidangMinat string
	aktif       bool
}

var dataPeserta [MAXDATA]Peserta
var jumlahPeserta int
var nextID int = 1
var bidangList [7]string = [7]string{
	"Web Development",
	"Mobile Development",
	"Data Science",
	"Cybersecurity",
	"UI/UX Design",
	"Cloud Computing",
	"Artificial Intelligence",
}

var reader = bufio.NewReader(os.Stdin)

func bacaTeks(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func pilihBidang() string {
	fmt.Println("Pilih bidang minat:")
	for i, b := range bidangList {
		fmt.Printf("%d. %s\n", i+1, b)
	}
	pil := bacaTeks("Pilihan (1-7): ")
	var n int
	fmt.Sscanf(pil, "%d", &n)
	if n < 1 || n > 7 {
		fmt.Println("Pilihan tidak valid, diset ke Web Development.")
		return bidangList[0]
	}
	return bidangList[n-1]
}

func cariIndeksID(id int) int {
	for i := 0; i < jumlahPeserta; i++ {
		if dataPeserta[i].id == id {
			return i
		}
	}
	return -1
}

// ----------------------------------------------------------
// 1. TAMBAH PESERTA
// ----------------------------------------------------------

func tambahPeserta() {
	if jumlahPeserta >= MAXDATA {
		fmt.Println("Data penuh! Tidak bisa menambah peserta lagi.")
		return
	}

	var p Peserta
	p.id = nextID
	p.aktif = true
	p.tglDaftar = time.Now().Format("02-01-2006")
	nextID++

	fmt.Printf("ID otomatis: %d\n", p.id)
	p.nama = bacaTeks("Nama lengkap : ")

	// cek duplikat langsung setelah input nama
	for i := 0; i < jumlahPeserta; i++ {
		if strings.ToLower(dataPeserta[i].nama) == strings.ToLower(p.nama) {
			fmt.Println("Peserta dengan nama ini sudah terdaftar!")
			nextID-- // kembalikan ID karena batal
			return
		}
	}

	p.email = bacaTeks("Email        : ")
	p.noTelp = bacaTeks("No. Telepon  : ")
	p.bidangMinat = pilihBidang()

	dataPeserta[jumlahPeserta] = p
	jumlahPeserta++

	fmt.Println("Peserta berhasil ditambahkan!")
}

// ----------------------------------------------------------
// 2. TAMPILKAN SEMUA PESERTA
// ----------------------------------------------------------

func tampilSemua() {
	if jumlahPeserta == 0 {
		fmt.Println("Belum ada data peserta.")
		return
	}

	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("%-4s %-20s %-25s %-12s %s\n", "ID", "Nama", "Bidang Minat", "Tgl Daftar", "Status")
	fmt.Println("--------------------------------------------------------------")

	for i := 0; i < jumlahPeserta; i++ {
		status := "Aktif"
		if !dataPeserta[i].aktif {
			status = "Nonaktif"
		}
		fmt.Printf("%-4d %-20s %-25s %-12s %s\n",
			dataPeserta[i].id,
			dataPeserta[i].nama,
			dataPeserta[i].bidangMinat,
			dataPeserta[i].tglDaftar,
			status)
	}

	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("Total: %d peserta\n", jumlahPeserta)
}

// ----------------------------------------------------------
// 3. UBAH DATA PESERTA
// ----------------------------------------------------------

func ubahPeserta() {
	if jumlahPeserta == 0 {
		fmt.Println("Belum ada data peserta.")
		return
	}

	input := bacaTeks("Masukkan ID yang ingin diubah: ")
	var idCari int
	fmt.Sscanf(input, "%d", &idCari)

	ketemu := cariIndeksID(idCari)
	if ketemu == -1 {
		fmt.Println("ID tidak ditemukan!")
		return
	}

	dataPeserta[ketemu].nama = bacaTeks("Nama baru  : ")
	dataPeserta[ketemu].email = bacaTeks("Email baru : ")
	dataPeserta[ketemu].noTelp = bacaTeks("Telp baru  : ")
	dataPeserta[ketemu].bidangMinat = pilihBidang()

	fmt.Println("Data berhasil diubah!")
}

// ----------------------------------------------------------
// 4. HAPUS PESERTA
// ----------------------------------------------------------

func hapusPeserta() {
	if jumlahPeserta == 0 {
		fmt.Println("Belum ada data peserta.")
		return
	}

	input := bacaTeks("Masukkan ID yang ingin dihapus: ")
	var idCari int
	fmt.Sscanf(input, "%d", &idCari)

	ketemu := cariIndeksID(idCari)
	if ketemu == -1 {
		fmt.Println("ID tidak ditemukan!")
		return
	}

	for i := ketemu; i < jumlahPeserta-1; i++ {
		dataPeserta[i] = dataPeserta[i+1]
	}
	dataPeserta[jumlahPeserta-1] = Peserta{}
	jumlahPeserta--

	// rapikan ID supaya tetap urut
	for i := 0; i < jumlahPeserta; i++ {
		dataPeserta[i].id = i + 1
	}
	nextID = jumlahPeserta + 1

	fmt.Println("Peserta berhasil dihapus!")
}

// ----------------------------------------------------------
// 5. SEQUENTIAL SEARCH (by nama)
// ----------------------------------------------------------

func sequentialSearch(keyword string) {
	ketemu := false

	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("%-4s %-20s %-25s\n", "ID", "Nama", "Bidang Minat")
	fmt.Println("--------------------------------------------------------------")

	for i := 0; i < jumlahPeserta; i++ {
		if strings.Contains(strings.ToLower(dataPeserta[i].nama), strings.ToLower(keyword)) {
			fmt.Printf("%-4d %-20s %-25s\n",
				dataPeserta[i].id,
				dataPeserta[i].nama,
				dataPeserta[i].bidangMinat)
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Peserta tidak ditemukan.")
	}
	fmt.Println("--------------------------------------------------------------")
}

// ----------------------------------------------------------
// 6. BINARY SEARCH (by bidang minat)
// ----------------------------------------------------------

func binarySearch(bidangCari string) {
	var temp [MAXDATA]Peserta
	for i := 0; i < jumlahPeserta; i++ {
		temp[i] = dataPeserta[i]
	}

	// pre-sort by bidang minat (insertion sort)
	for i := 1; i < jumlahPeserta; i++ {
		kunci := temp[i]
		j := i - 1
		for j >= 0 && strings.ToLower(temp[j].bidangMinat) > strings.ToLower(kunci.bidangMinat) {
			temp[j+1] = temp[j]
			j--
		}
		temp[j+1] = kunci
	}

	// binary search
	kiri := 0
	kanan := jumlahPeserta - 1
	ketemu := false

	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("%-4s %-20s %-25s\n", "ID", "Nama", "Bidang Minat")
	fmt.Println("--------------------------------------------------------------")

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		tengahBidang := strings.ToLower(temp[tengah].bidangMinat)
		target := strings.ToLower(bidangCari)

		if tengahBidang == target {
			fmt.Printf("%-4d %-20s %-25s\n",
				temp[tengah].id,
				temp[tengah].nama,
				temp[tengah].bidangMinat)
			ketemu = true
			break
		} else if tengahBidang < target {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if !ketemu {
		fmt.Println("Peserta tidak ditemukan.")
	}
	fmt.Println("--------------------------------------------------------------")
}

// ----------------------------------------------------------
// 7. SELECTION SORT (by ID)
// ----------------------------------------------------------

func selectionSort() {
	for i := 0; i < jumlahPeserta-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahPeserta; j++ {
			if dataPeserta[j].id < dataPeserta[minIdx].id {
				minIdx = j
			}
		}
		dataPeserta[i], dataPeserta[minIdx] = dataPeserta[minIdx], dataPeserta[i]
	}
	fmt.Println("Data berhasil diurutkan berdasarkan ID (Selection Sort).")
}

// ----------------------------------------------------------
// 8. INSERTION SORT (by nama)
// ----------------------------------------------------------

func insertionSort() {
	for i := 1; i < jumlahPeserta; i++ {
		kunci := dataPeserta[i]
		j := i - 1
		for j >= 0 && strings.ToLower(dataPeserta[j].nama) > strings.ToLower(kunci.nama) {
			dataPeserta[j+1] = dataPeserta[j]
			j--
		}
		dataPeserta[j+1] = kunci
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Nama (Insertion Sort).")
}

// ----------------------------------------------------------
// 9. STATISTIK
// ----------------------------------------------------------

func tampilStatistik() {
	fmt.Println("==============================")
	fmt.Println("      STATISTIK PESERTA       ")
	fmt.Println("==============================")
	fmt.Printf("Total peserta: %d\n", jumlahPeserta)
	fmt.Println("------------------------------")

	for b := 0; b < 7; b++ {
		count := 0
		for i := 0; i < jumlahPeserta; i++ {
			if dataPeserta[i].bidangMinat == bidangList[b] {
				count++
			}
		}
		persen := 0.0
		if jumlahPeserta > 0 {
			persen = float64(count) / float64(jumlahPeserta) * 100
		}
		fmt.Printf("%-25s: %d orang (%.1f%%)\n", bidangList[b], count, persen)
	}

	fmt.Println("==============================")
}

// ----------------------------------------------------------
// DATA AWAL
// ----------------------------------------------------------

func isiDataAwal() {
	dataPeserta[0] = Peserta{1, "Andi Pratama", "andi@email.com", "081234567890", "01-04-2025", "Web Development", true}
	dataPeserta[1] = Peserta{2, "Budi Santoso", "budi@email.com", "082345678901", "02-04-2025", "Data Science", true}
	dataPeserta[2] = Peserta{3, "Citra Dewi", "citra@email.com", "083456789012", "05-04-2025", "Mobile Development", true}
	dataPeserta[3] = Peserta{4, "Dian Permata", "dian@email.com", "084567890123", "07-04-2025", "Cybersecurity", true}
	dataPeserta[4] = Peserta{5, "Eko Wahyudi", "eko@email.com", "085678901234", "10-04-2025", "Artificial Intelligence", true}
	dataPeserta[5] = Peserta{6, "Fitri Amalia", "fitri@email.com", "086789012345", "11-04-2025", "UI/UX Design", true}
	dataPeserta[6] = Peserta{7, "Gilang Ramadhan", "gilang@email.com", "087890123456", "13-04-2025", "Cloud Computing", true}
	dataPeserta[7] = Peserta{8, "Hana Safitri", "hana@email.com", "088901234567", "15-04-2025", "Web Development", true}
	dataPeserta[8] = Peserta{9, "Ivan Kurniawan", "ivan@email.com", "089012345678", "18-04-2025", "Data Science", true}
	dataPeserta[9] = Peserta{10, "Jasmine Putri", "jasmine@email.com", "081123456789", "20-04-2025", "Artificial Intelligence", true}
	jumlahPeserta = 10
	nextID = 11
}

// ----------------------------------------------------------
// MAIN
// ----------------------------------------------------------

func main() {
	isiDataAwal()

	for {
		fmt.Println()
		fmt.Println("==============================")
		fmt.Println("   KURSUS-IN: Menu Utama      ")
		fmt.Println("==============================")
		fmt.Println("1. Tambah Peserta")
		fmt.Println("2. Tampilkan Semua")
		fmt.Println("3. Ubah Peserta")
		fmt.Println("4. Hapus Peserta")
		fmt.Println("5. Cari - Sequential Search (by Nama)")
		fmt.Println("6. Cari - Binary Search (by Bidang Minat)")
		fmt.Println("7. Urutkan by ID - Selection Sort")
		fmt.Println("8. Urutkan by Nama - Insertion Sort")
		fmt.Println("9. Statistik")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")

		pilihan := bacaTeks("")
		var n int
		fmt.Sscanf(pilihan, "%d", &n)

		switch n {
		case 1:
			tambahPeserta()
		case 2:
			tampilSemua()
		case 3:
			ubahPeserta()
		case 4:
			hapusPeserta()
		case 5:
			keyword := bacaTeks("Kata kunci nama: ")
			sequentialSearch(keyword)
		case 6:
			bidang := bacaTeks("Bidang minat: ")
			binarySearch(bidang)
		case 7:
			selectionSort()
			tampilSemua()
		case 8:
			insertionSort()
			tampilSemua()
		case 9:
			tampilStatistik()
		case 0:
			fmt.Println("Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
