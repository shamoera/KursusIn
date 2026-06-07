package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ============================================================
//  KURSUS-IN: Sistem Pendaftaran Kursus Online Terpadu
//  Tugas Besar Algoritma Pemrograman 2
//  Fakultas Informatika – Telkom University
// ============================================================

const MAXDATA = 100

type Peserta struct {
	ID            string
	NamaLengkap   string
	Email         string
	NoTelp        string
	TanggalDaftar string
	BidangMinat   string
	StatusAktif   bool
}

type DataPeserta struct {
	data   [MAXDATA]Peserta
	jumlah int
}

var db DataPeserta
var reader = bufio.NewReader(os.Stdin)

var bidangMinatList = [7]string{
	"Web Development",
	"Mobile Development",
	"Data Science",
	"Cybersecurity",
	"UI/UX Design",
	"Cloud Computing",
	"Artificial Intelligence",
}

func inputStr() string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ReplaceAll(text, "\r", ""))
}

func getTanggal() string {
	return time.Now().Format("02-01-2006")
}

func generateID() string {
	// Cari nomor urut terbesar agar tidak duplikat setelah hapus
	maxNum := db.jumlah
	for i := 0; i < db.jumlah; i++ {
		if len(db.data[i].ID) == 7 {
			num, err := strconv.Atoi(db.data[i].ID[3:])
			if err == nil && num > maxNum {
				maxNum = num
			}
		}
	}
	return fmt.Sprintf("KRS%04d", maxNum+1)
}

func garis(n int) string { return strings.Repeat("─", n) }

func cetakHeader() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════╗")
	fmt.Println("║      KURSUS-IN: Sistem Pendaftaran Kursus        ║")
	fmt.Println("║              Online Terpadu (v1.0)               ║")
	fmt.Println("╚══════════════════════════════════════════════════╝")
}

func pilihBidangMinat(label string) (string, bool) {
	fmt.Printf("  %s:\n", label)
	for i, b := range bidangMinatList {
		fmt.Printf("    %d. %s\n", i+1, b)
	}
	fmt.Print("  Pilih (1-7): ")
	pilStr := inputStr()
	pil, err := strconv.Atoi(pilStr)
	if err != nil || pil < 1 || pil > 7 {
		fmt.Println("  [!] Pilihan tidak valid!")
		return "", false
	}
	return bidangMinatList[pil-1], true
}

func tambahPeserta() {
	if db.jumlah >= MAXDATA {
		fmt.Println("  [!] Data peserta sudah penuh (maks 100)!")
		return
	}
	fmt.Println("\n  ── TAMBAH PESERTA BARU ──")

	var p Peserta
	p.ID = generateID()
	p.TanggalDaftar = getTanggal()
	p.StatusAktif = true

	fmt.Printf("  ID Pendaftaran : %s (otomatis)\n", p.ID)

	fmt.Print("  Nama Lengkap   : ")
	p.NamaLengkap = inputStr()
	if p.NamaLengkap == "" {
		fmt.Println("  [!] Nama tidak boleh kosong!")
		return
	}

	fmt.Print("  Email          : ")
	p.Email = inputStr()

	fmt.Print("  No. Telepon    : ")
	p.NoTelp = inputStr()

	bidang, ok := pilihBidangMinat("Bidang Minat")
	if !ok {
		return
	}
	p.BidangMinat = bidang

	db.data[db.jumlah] = p
	db.jumlah++
	fmt.Printf("\n  [✓] Peserta '%s' berhasil didaftarkan! (ID: %s)\n", p.NamaLengkap, p.ID)
}

func ubahPeserta() {
	if db.jumlah == 0 {
		fmt.Println("  [!] Belum ada data peserta!")
		return
	}
	fmt.Println("\n  ── UBAH DATA PESERTA ──")
	fmt.Print("  Masukkan ID peserta yang ingin diubah: ")
	id := strings.ToUpper(inputStr())

	idx := -1
	for i := 0; i < db.jumlah; i++ {
		if db.data[i].ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("  [!] Peserta tidak ditemukan!")
		return
	}

	p := &db.data[idx]
	fmt.Printf("\n  Data ditemukan: [%s] %s | %s\n", p.ID, p.NamaLengkap, p.BidangMinat)
	fmt.Println("  (Tekan Enter untuk melewati / tidak mengubah)")

	fmt.Printf("  Nama Lengkap baru [%s]: ", p.NamaLengkap)
	if nama := inputStr(); nama != "" {
		p.NamaLengkap = nama
	}

	fmt.Printf("  Email baru [%s]: ", p.Email)
	if email := inputStr(); email != "" {
		p.Email = email
	}

	fmt.Printf("  No. Telepon baru [%s]: ", p.NoTelp)
	if telp := inputStr(); telp != "" {
		p.NoTelp = telp
	}

	fmt.Printf("  Ubah Bidang Minat [%s]? (y/n): ", p.BidangMinat)
	if strings.ToLower(inputStr()) == "y" {
		if bidang, ok := pilihBidangMinat("Bidang Minat Baru"); ok {
			p.BidangMinat = bidang
		}
	}

	fmt.Printf("  Status Aktif saat ini: ")
	if p.StatusAktif {
		fmt.Print("Aktif")
	} else {
		fmt.Print("Tidak Aktif")
	}
	fmt.Print(" – ubah? (y/n): ")
	if strings.ToLower(inputStr()) == "y" {
		p.StatusAktif = !p.StatusAktif
	}

	fmt.Println("  [✓] Data peserta berhasil diperbarui!")
}

func hapusPeserta() {
	if db.jumlah == 0 {
		fmt.Println("  [!] Belum ada data peserta!")
		return
	}
	fmt.Println("\n  ── HAPUS DATA PESERTA ──")
	fmt.Print("  Masukkan ID peserta yang ingin dihapus: ")
	id := strings.ToUpper(inputStr())

	idx := -1
	for i := 0; i < db.jumlah; i++ {
		if db.data[i].ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("  [!] Peserta tidak ditemukan!")
		return
	}

	fmt.Printf("  Yakin hapus '%s' (%s)? (y/n): ", db.data[idx].NamaLengkap, db.data[idx].ID)
	if strings.ToLower(inputStr()) != "y" {
		fmt.Println("  Penghapusan dibatalkan.")
		return
	}

	// Geser elemen ke kiri
	for i := idx; i < db.jumlah-1; i++ {
		db.data[i] = db.data[i+1]
	}
	db.data[db.jumlah-1] = Peserta{} // kosongkan slot terakhir
	db.jumlah--
	fmt.Println("  [✓] Peserta berhasil dihapus!")
}

func cetakBarisPeserta(p Peserta) {
	status := "Aktif"
	if !p.StatusAktif {
		status = "Nonaktif"
	}
	fmt.Printf("│ %-8s │ %-22s │ %-26s │ %-22s │ %-11s │ %-8s │\n",
		p.ID, p.NamaLengkap, p.Email, p.BidangMinat, p.TanggalDaftar, status)
}

func tampilkanSemuaPeserta() {
	fmt.Println("\n  ── DAFTAR SELURUH PESERTA KURSUS ──")
	if db.jumlah == 0 {
		fmt.Println("  [!] Belum ada data peserta.")
		return
	}
	batas := garis(114)
	fmt.Println(batas)
	fmt.Printf("│ %-8s │ %-22s │ %-26s │ %-22s │ %-11s │ %-8s │\n",
		"ID", "Nama Lengkap", "Email", "Bidang Minat", "Tgl Daftar", "Status")
	fmt.Println(batas)
	for i := 0; i < db.jumlah; i++ {
		cetakBarisPeserta(db.data[i])
	}
	fmt.Println(batas)
	fmt.Printf("  Total: %d peserta\n", db.jumlah)
}

// Sequential Search – berdasarkan nama (mengandung keyword)
func sequentialSearchNama(keyword string) []int {
	var hasil []int
	key := strings.ToLower(keyword)
	for i := 0; i < db.jumlah; i++ {
		if strings.Contains(strings.ToLower(db.data[i].NamaLengkap), key) {
			hasil = append(hasil, i)
		}
	}
	return hasil
}

// Sequential Search – berdasarkan bidang minat (exact)
func sequentialSearchBidang(bidang string) []int {
	var hasil []int
	for i := 0; i < db.jumlah; i++ {
		if strings.EqualFold(db.data[i].BidangMinat, bidang) {
			hasil = append(hasil, i)
		}
	}
	return hasil
}

// Binary Search – berdasarkan nama lengkap (exact match, case-insensitive)
// Catatan: Binary Search membutuhkan data terurut; fungsi ini
// bekerja pada salinan data yang sudah diurutkan (Insertion Sort).
func binarySearchNama(namaCari string) *Peserta {
	if db.jumlah == 0 {
		return nil
	}

	var sorted [MAXDATA]Peserta
	for i := 0; i < db.jumlah; i++ {
		sorted[i] = db.data[i]
	}
	// Insertion Sort pada salinan
	for i := 1; i < db.jumlah; i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && strings.ToLower(sorted[j].NamaLengkap) > strings.ToLower(key.NamaLengkap) {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}

	// Binary Search
	target := strings.ToLower(namaCari)
	lo, hi := 0, db.jumlah-1
	for lo <= hi {
		mid := (lo + hi) / 2
		midNama := strings.ToLower(sorted[mid].NamaLengkap)
		if midNama == target {
			// Kembalikan salinan (bukan pointer ke db asli agar aman)
			hasil := sorted[mid]
			return &hasil
		} else if midNama < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return nil
}

func menuCariPeserta() {
	for {
		fmt.Println("\n  ── CARI DATA PESERTA ──")
		fmt.Println("  1. Cari Nama (Sequential Search – mengandung kata kunci)")
		fmt.Println("  2. Cari Nama (Binary Search – nama lengkap exact)")
		fmt.Println("  3. Cari Bidang Minat (Sequential Search)")
		fmt.Println("  0. Kembali")
		fmt.Print("  Pilih: ")

		switch inputStr() {
		case "1":
			fmt.Print("  Kata kunci nama: ")
			keyword := inputStr()
			hasil := sequentialSearchNama(keyword)
			if len(hasil) == 0 {
				fmt.Printf("  [!] Tidak ada peserta dengan nama mengandung '%s'.\n", keyword)
			} else {
				fmt.Printf("\n  Ditemukan %d peserta:\n", len(hasil))
				fmt.Println("  " + garis(70))
				for _, idx := range hasil {
					p := db.data[idx]
					fmt.Printf("  [%s] %-22s | %-22s | %s\n",
						p.ID, p.NamaLengkap, p.BidangMinat, p.TanggalDaftar)
				}
				fmt.Println("  " + garis(70))
			}

		case "2":
			fmt.Print("  Nama lengkap (exact): ")
			nama := inputStr()
			p := binarySearchNama(nama)
			if p == nil {
				fmt.Printf("  [!] Peserta dengan nama '%s' tidak ditemukan.\n", nama)
				fmt.Println("  Catatan: Binary Search memerlukan nama yang persis sama (case-insensitive).")
			} else {
				fmt.Println("\n  Peserta ditemukan:")
				fmt.Println("  " + garis(60))
				fmt.Printf("  ID            : %s\n", p.ID)
				fmt.Printf("  Nama Lengkap  : %s\n", p.NamaLengkap)
				fmt.Printf("  Email         : %s\n", p.Email)
				fmt.Printf("  No. Telepon   : %s\n", p.NoTelp)
				fmt.Printf("  Bidang Minat  : %s\n", p.BidangMinat)
				fmt.Printf("  Tgl Daftar    : %s\n", p.TanggalDaftar)
				status := "Aktif"
				if !p.StatusAktif {
					status = "Tidak Aktif"
				}
				fmt.Printf("  Status        : %s\n", status)
				fmt.Println("  " + garis(60))
			}

		case "3":
			bidang, ok := pilihBidangMinat("Pilih Bidang Minat")
			if !ok {
				continue
			}
			hasil := sequentialSearchBidang(bidang)
			if len(hasil) == 0 {
				fmt.Printf("  [!] Tidak ada peserta dengan bidang minat '%s'.\n", bidang)
			} else {
				fmt.Printf("\n  Peserta dengan bidang minat '%s' (%d orang):\n", bidang, len(hasil))
				fmt.Println("  " + garis(60))
				for _, idx := range hasil {
					p := db.data[idx]
					status := "Aktif"
					if !p.StatusAktif {
						status = "Nonaktif"
					}
					fmt.Printf("  [%s] %-22s | %s | %s\n",
						p.ID, p.NamaLengkap, p.TanggalDaftar, status)
				}
				fmt.Println("  " + garis(60))
			}

		case "0":
			return
		default:
			fmt.Println("  [!] Pilihan tidak valid.")
		}
	}
}

//  ALGORITMA

// Selection Sort – diurutkan berdasarkan ID Pendaftaran (ascending)
func selectionSortByID() {
	for i := 0; i < db.jumlah-1; i++ {
		minIdx := i
		for j := i + 1; j < db.jumlah; j++ {
			if db.data[j].ID < db.data[minIdx].ID {
				minIdx = j
			}
		}
		if minIdx != i {
			db.data[i], db.data[minIdx] = db.data[minIdx], db.data[i]
		}
	}
}

// Insertion Sort – diurutkan berdasarkan nama (ascending, alfabetis)
func insertionSortByNama() {
	for i := 1; i < db.jumlah; i++ {
		key := db.data[i]
		j := i - 1
		for j >= 0 && strings.ToLower(db.data[j].NamaLengkap) > strings.ToLower(key.NamaLengkap) {
			db.data[j+1] = db.data[j]
			j--
		}
		db.data[j+1] = key
	}
}

func menuUrutkanPeserta() {
	for {
		fmt.Println("\n  ── URUTKAN DATA PESERTA ──")
		fmt.Println("  1. Urutkan berdasarkan ID Pendaftaran – Selection Sort (A→Z)")
		fmt.Println("  2. Urutkan berdasarkan Nama Alfabetis – Insertion Sort (A→Z)")
		fmt.Println("  0. Kembali")
		fmt.Print("  Pilih: ")

		switch inputStr() {
		case "1":
			selectionSortByID()
			fmt.Println("  [✓] Data diurutkan berdasarkan ID (Selection Sort).")
			tampilkanSemuaPeserta()
		case "2":
			insertionSortByNama()
			fmt.Println("  [✓] Data diurutkan berdasarkan Nama Alfabetis (Insertion Sort).")
			tampilkanSemuaPeserta()
		case "0":
			return
		default:
			fmt.Println("  [!] Pilihan tidak valid.")
		}
	}
}

func tampilkanStatistik() {
	fmt.Println("\n  ══════════════════════════════════════════════")
	fmt.Println("         STATISTIK PENDAFTARAN KURSUS-IN       ")
	fmt.Println("  ══════════════════════════════════════════════")

	totalAktif := 0
	for i := 0; i < db.jumlah; i++ {
		if db.data[i].StatusAktif {
			totalAktif++
		}
	}

	fmt.Printf("  Total Peserta Terdaftar  : %d orang\n", db.jumlah)
	fmt.Printf("  Peserta Aktif            : %d orang\n", totalAktif)
	fmt.Printf("  Peserta Tidak Aktif      : %d orang\n", db.jumlah-totalAktif)
	fmt.Println()
	fmt.Println("  Jumlah Pendaftar per Bidang Minat:")
	fmt.Println("  " + garis(50))

	maxCount := 0
	counts := [7]int{}
	for b, bidang := range bidangMinatList {
		for i := 0; i < db.jumlah; i++ {
			if db.data[i].BidangMinat == bidang {
				counts[b]++
			}
		}
		if counts[b] > maxCount {
			maxCount = counts[b]
		}
	}

	for b, bidang := range bidangMinatList {
		bar := ""
		for k := 0; k < counts[b]; k++ {
			bar += "█"
		}
		fmt.Printf("  %-22s │ %2d │ %s\n", bidang, counts[b], bar)
	}
	fmt.Println("  " + garis(50))
}

//  MENU

func menuKelolaPeserta() {
	for {
		fmt.Println("\n  ── KELOLA DATA PESERTA ──")
		fmt.Println("  1. Tambah Peserta Baru")
		fmt.Println("  2. Ubah Data Peserta")
		fmt.Println("  3. Hapus Peserta")
		fmt.Println("  4. Tampilkan Semua Peserta")
		fmt.Println("  0. Kembali ke Menu Utama")
		fmt.Print("  Pilih: ")

		switch inputStr() {
		case "1":
			tambahPeserta()
		case "2":
			ubahPeserta()
		case "3":
			hapusPeserta()
		case "4":
			tampilkanSemuaPeserta()
		case "0":
			return
		default:
			fmt.Println("  [!] Pilihan tidak valid.")
		}
	}
}

//  DATA AWAL (SAMPLE)

func inisialisasiSampleData() {
	samples := []Peserta{
		{"KRS0001", "Andi Pratama", "andi@email.com", "081234567890", "01-04-2026", "Web Development", true},
		{"KRS0002", "Budi Santoso", "budi@email.com", "082345678901", "02-04-2026", "Data Science", true},
		{"KRS0003", "Citra Dewi", "citra@email.com", "083456789012", "05-04-2026", "Mobile Development", true},
		{"KRS0004", "Dian Permata", "dian@email.com", "084567890123", "07-04-2026", "Cybersecurity", true},
		{"KRS0005", "Eko Wahyudi", "eko@email.com", "085678901234", "10-04-2026", "Artificial Intelligence", true},
		{"KRS0006", "Fitri Amalia", "fitri@email.com", "086789012345", "11-04-2026", "UI/UX Design", true},
		{"KRS0007", "Gilang Ramadhan", "gilang@email.com", "087890123456", "13-04-2026", "Cloud Computing", true},
		{"KRS0008", "Hana Safitri", "hana@email.com", "088901234567", "15-04-2026", "Web Development", true},
		{"KRS0009", "Ivan Kurniawan", "ivan@email.com", "089012345678", "18-04-2026", "Data Science", false},
		{"KRS0010", "Jasmine Putri", "jasmine@email.com", "081123456789", "20-04-2026", "Artificial Intelligence", true},
	}
	for i, p := range samples {
		db.data[i] = p
	}
	db.jumlah = len(samples)
}

func main() {
	inisialisasiSampleData()

	for {
		cetakHeader()
		fmt.Println("  1. Kelola Data Peserta (Tambah / Ubah / Hapus)")
		fmt.Println("  2. Cari Data Peserta (Sequential & Binary Search)")
		fmt.Println("  3. Urutkan Data Peserta (Selection & Insertion Sort)")
		fmt.Println("  4. Statistik Pendaftaran")
		fmt.Println("  5. Tampilkan Semua Peserta")
		fmt.Println("  0. Keluar")
		fmt.Print("  Pilih menu: ")

		pilihan := inputStr()
		switch pilihan {
		case "1":
			menuKelolaPeserta()
		case "2":
			menuCariPeserta()
		case "3":
			menuUrutkanPeserta()
		case "4":
			tampilkanStatistik()
		case "5":
			tampilkanSemuaPeserta()
		case "0":
			fmt.Println("\n  Terima kasih telah menggunakan KursusIn!")
			fmt.Println("  Sampai jumpa! ")
			return
		default:
			fmt.Println("  [!] Pilihan tidak valid. Coba lagi.")
		}
	}
}
