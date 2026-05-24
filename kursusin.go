package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ============================================================
// KONSTANTA DAN TIPE DATA
// ============================================================

const nMax int = 1000
const namaFile string = "peserta.csv"

type Peserta struct {
	id          string
	nama        string
	tanggal     string
	bidangMinat string
}

type Kursus struct {
	idKursus   string
	namaKursus string
	bidang     string
}

type BidangMinat struct {
	idBidang   string
	namaBidang string
	jumlah     int
}

type DaftarPeserta [1000]Peserta
type DaftarKursus [100]Kursus
type DaftarBidang [20]BidangMinat

// ============================================================
// VARIABEL GLOBAL
// ============================================================

var (
	daftarPeserta DaftarPeserta
	nPeserta      int
	daftarKursus  DaftarKursus
	nKursus       int
	daftarBidang  DaftarBidang
	nBidang       int
	reader        = bufio.NewReader(os.Stdin)
)

// ============================================================
// UTILITAS INPUT
// ============================================================

func bacaInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func cetakGaris() {
	fmt.Println("============================================================")
}

func cetakGarisTipis() {
	fmt.Println("------------------------------------------------------------")
}

// ============================================================
// SIMPAN & LOAD CSV
// ============================================================

func simpanCSV() {
	file, err := os.Create(namaFile)
	if err != nil {
		fmt.Println("Gagal menyimpan data:", err)
		return
	}
	defer file.Close()

	// Tulis header
	fmt.Fprintln(file, "ID,Nama,Tanggal,BidangMinat")

	// Tulis setiap peserta
	i := 0
	for i < nPeserta {
		fmt.Fprintf(file, "%s,%s,%s,%s\n",
			daftarPeserta[i].id,
			daftarPeserta[i].nama,
			daftarPeserta[i].tanggal,
			daftarPeserta[i].bidangMinat)
		i = i + 1
	}
}

func loadCSV() {
	file, err := os.Open(namaFile)
	if err != nil {
		// File belum ada, mulai kosong
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	barisPertama := true
	nPeserta = 0

	for scanner.Scan() {
		baris := scanner.Text()
		if barisPertama {
			barisPertama = false
			continue // skip header
		}
		if baris == "" {
			continue
		}

		// Pisah berdasarkan koma
		kolom := strings.Split(baris, ",")
		if len(kolom) < 4 {
			continue
		}

		daftarPeserta[nPeserta] = Peserta{
			id:          kolom[0],
			nama:        kolom[1],
			tanggal:     kolom[2],
			bidangMinat: kolom[3],
		}
		updateJumlahBidang(kolom[3], 1)
		nPeserta = nPeserta + 1
	}

	if nPeserta > 0 {
		fmt.Printf("Data berhasil dimuat: %d peserta ditemukan.\n", nPeserta)
	}
}

// ============================================================
// MANAJEMEN BIDANG MINAT
// ============================================================

func inisialisasiBidang() {
	daftarBidang[0] = BidangMinat{"B001", "Web Development", 0}
	daftarBidang[1] = BidangMinat{"B002", "Mobile Development", 0}
	daftarBidang[2] = BidangMinat{"B003", "Data Science", 0}
	daftarBidang[3] = BidangMinat{"B004", "Cybersecurity", 0}
	daftarBidang[4] = BidangMinat{"B005", "UI/UX Design", 0}
	nBidang = 5
}

func inisialisasiKursus() {
	daftarKursus[0] = Kursus{"K001", "HTML & CSS Dasar", "Web Development"}
	daftarKursus[1] = Kursus{"K002", "JavaScript Modern", "Web Development"}
	daftarKursus[2] = Kursus{"K003", "Flutter Pemula", "Mobile Development"}
	daftarKursus[3] = Kursus{"K004", "Android Kotlin", "Mobile Development"}
	daftarKursus[4] = Kursus{"K005", "Python Data Analysis", "Data Science"}
	daftarKursus[5] = Kursus{"K006", "Machine Learning Dasar", "Data Science"}
	daftarKursus[6] = Kursus{"K007", "Ethical Hacking", "Cybersecurity"}
	daftarKursus[7] = Kursus{"K008", "Figma untuk Pemula", "UI/UX Design"}
	nKursus = 8
}

func updateJumlahBidang(bidang string, delta int) {
	i := 0
	for i < nBidang {
		if daftarBidang[i].namaBidang == bidang {
			daftarBidang[i].jumlah = daftarBidang[i].jumlah + delta
		}
		i = i + 1
	}
}

func tampilkanBidang() {
	fmt.Println("Bidang Minat yang Tersedia:")
	i := 0
	for i < nBidang {
		fmt.Printf("  %d. %s\n", i+1, daftarBidang[i].namaBidang)
		i = i + 1
	}
}

func tampilkanKursusByBidang(bidang string) {
	fmt.Printf("Kursus untuk bidang %s:\n", bidang)
	i := 0
	ada := false
	for i < nKursus {
		if daftarKursus[i].bidang == bidang {
			fmt.Printf("  - [%s] %s\n", daftarKursus[i].idKursus, daftarKursus[i].namaKursus)
			ada = true
		}
		i = i + 1
	}
	if !ada {
		fmt.Println("  Tidak ada kursus tersedia.")
	}
}

// ============================================================
// GENERATE ID OTOMATIS
// ============================================================

func generateID() string {
	// Cari ID terbesar lalu tambah 1
	maxNum := nPeserta
	return fmt.Sprintf("P%03d", maxNum+1)
}

// ============================================================
// FITUR A: TAMBAH PESERTA
// ============================================================

func tambahPeserta() {
	cetakGaris()
	fmt.Println("         TAMBAH PESERTA BARU")
	cetakGaris()

	if nPeserta >= nMax {
		fmt.Println("Data peserta sudah penuh!")
		return
	}

	nama := bacaInput("Nama Lengkap     : ")
	tanggal := bacaInput("Tanggal Daftar   : ")

	fmt.Println()
	tampilkanBidang()
	pilihBidang := bacaInput("Pilih nomor bidang minat: ")

	var nomorBidang int
	fmt.Sscan(pilihBidang, &nomorBidang)

	if nomorBidang < 1 || nomorBidang > nBidang {
		fmt.Println("Nomor bidang tidak valid!")
		return
	}

	bidang := daftarBidang[nomorBidang-1].namaBidang
	tampilkanKursusByBidang(bidang)

	id := generateID()
	daftarPeserta[nPeserta] = Peserta{id, nama, tanggal, bidang}
	nPeserta = nPeserta + 1
	updateJumlahBidang(bidang, 1)

	// Auto simpan ke CSV
	simpanCSV()

	fmt.Println()
	fmt.Printf("Peserta berhasil didaftarkan!\n")
	fmt.Printf("ID Peserta: %s\n", id)
	fmt.Printf("Nama      : %s\n", nama)
	fmt.Printf("Tanggal   : %s\n", tanggal)
	fmt.Printf("Bidang    : %s\n", bidang)
}

// ============================================================
// FITUR A: UBAH PESERTA
// ============================================================

func ubahPeserta() {
	cetakGaris()
	fmt.Println("           UBAH DATA PESERTA")
	cetakGaris()

	if nPeserta == 0 {
		fmt.Println("Belum ada data peserta.")
		return
	}

	idCari := bacaInput("Masukkan ID Peserta yang ingin diubah: ")

	idx := -1
	i := 0
	for i < nPeserta {
		if daftarPeserta[i].id == idCari {
			idx = i
		}
		i = i + 1
	}

	if idx == -1 {
		fmt.Println("Peserta dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println()
	fmt.Println("Data saat ini:")
	fmt.Printf("  Nama     : %s\n", daftarPeserta[idx].nama)
	fmt.Printf("  Tanggal  : %s\n", daftarPeserta[idx].tanggal)
	fmt.Printf("  Bidang   : %s\n", daftarPeserta[idx].bidangMinat)
	fmt.Println()

	namaBaru := bacaInput("Nama Baru (Enter jika tidak diubah)   : ")
	tanggalBaru := bacaInput("Tanggal Baru (Enter jika tidak diubah): ")

	fmt.Println()
	tampilkanBidang()
	pilihBidang := bacaInput("Bidang baru (0 jika tidak diubah): ")

	var nomorBidang int
	fmt.Sscan(pilihBidang, &nomorBidang)

	if namaBaru != "" {
		daftarPeserta[idx].nama = namaBaru
	}
	if tanggalBaru != "" {
		daftarPeserta[idx].tanggal = tanggalBaru
	}
	if nomorBidang >= 1 && nomorBidang <= nBidang {
		updateJumlahBidang(daftarPeserta[idx].bidangMinat, -1)
		daftarPeserta[idx].bidangMinat = daftarBidang[nomorBidang-1].namaBidang
		updateJumlahBidang(daftarPeserta[idx].bidangMinat, 1)
	}

	// Auto simpan ke CSV
	simpanCSV()
	fmt.Println("Data peserta berhasil diubah dan disimpan!")
}

// ============================================================
// FITUR A: HAPUS PESERTA
// ============================================================

func hapusPeserta() {
	cetakGaris()
	fmt.Println("           HAPUS DATA PESERTA")
	cetakGaris()

	if nPeserta == 0 {
		fmt.Println("Belum ada data peserta.")
		return
	}

	idCari := bacaInput("Masukkan ID Peserta yang ingin dihapus: ")

	idx := -1
	i := 0
	for i < nPeserta {
		if daftarPeserta[i].id == idCari {
			idx = i
		}
		i = i + 1
	}

	if idx == -1 {
		fmt.Println("Peserta dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Printf("Hapus peserta '%s' (%s)? (y/n): ", daftarPeserta[idx].nama, daftarPeserta[idx].id)
	konfirmasi := bacaInput("")

	if strings.ToLower(konfirmasi) == "y" {
		updateJumlahBidang(daftarPeserta[idx].bidangMinat, -1)
		j := idx
		for j < nPeserta-1 {
			daftarPeserta[j] = daftarPeserta[j+1]
			j = j + 1
		}
		nPeserta = nPeserta - 1

		// Auto simpan ke CSV
		simpanCSV()
		fmt.Println("Peserta berhasil dihapus dan data disimpan!")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}

// ============================================================
// TAMPILKAN SEMUA PESERTA
// ============================================================

func tampilkanSemuaPeserta() {
	if nPeserta == 0 {
		fmt.Println("Belum ada data peserta.")
		return
	}
	fmt.Printf("%-6s %-25s %-12s %-20s\n", "ID", "Nama", "Tanggal", "Bidang Minat")
	cetakGarisTipis()
	i := 0
	for i < nPeserta {
		fmt.Printf("%-6s %-25s %-12s %-20s\n",
			daftarPeserta[i].id,
			daftarPeserta[i].nama,
			daftarPeserta[i].tanggal,
			daftarPeserta[i].bidangMinat)
		i = i + 1
	}
	fmt.Printf("\nTotal peserta: %d\n", nPeserta)
}

// ============================================================
// FITUR C: PENCARIAN - SEQUENTIAL SEARCH
// ============================================================

func sequentialSearchNama(nama string) {
	fmt.Printf("\nHasil Sequential Search untuk nama '%s':\n", nama)
	cetakGarisTipis()
	ketemu := false
	i := 0
	for i < nPeserta {
		if strings.Contains(strings.ToLower(daftarPeserta[i].nama), strings.ToLower(nama)) {
			fmt.Printf("%-6s %-25s %-12s %-20s\n",
				daftarPeserta[i].id,
				daftarPeserta[i].nama,
				daftarPeserta[i].tanggal,
				daftarPeserta[i].bidangMinat)
			ketemu = true
		}
		i = i + 1
	}
	if !ketemu {
		fmt.Println("Data tidak ditemukan.")
	}
}

func sequentialSearchBidang(bidang string) {
	fmt.Printf("\nHasil Sequential Search untuk bidang '%s':\n", bidang)
	cetakGarisTipis()
	ketemu := false
	i := 0
	for i < nPeserta {
		if strings.ToLower(daftarPeserta[i].bidangMinat) == strings.ToLower(bidang) {
			fmt.Printf("%-6s %-25s %-12s %-20s\n",
				daftarPeserta[i].id,
				daftarPeserta[i].nama,
				daftarPeserta[i].tanggal,
				daftarPeserta[i].bidangMinat)
			ketemu = true
		}
		i = i + 1
	}
	if !ketemu {
		fmt.Println("Data tidak ditemukan.")
	}
}

// ============================================================
// FITUR C: PENCARIAN - BINARY SEARCH
// ============================================================

func insertionSortByNama(arr *DaftarPeserta, n int) {
	var temp Peserta
	var i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = arr[j]
		for j > 0 && strings.ToLower(temp.nama) < strings.ToLower(arr[j-1].nama) {
			arr[j] = arr[j-1]
			j = j - 1
		}
		arr[j] = temp
		i = i + 1
	}
}

func binarySearchNama(arr DaftarPeserta, n int, nama string) {
	var arrSalin DaftarPeserta
	i := 0
	for i < n {
		arrSalin[i] = arr[i]
		i = i + 1
	}
	insertionSortByNama(&arrSalin, n)

	fmt.Printf("\nHasil Binary Search untuk nama '%s':\n", nama)
	cetakGarisTipis()

	kiri := 0
	kanan := n - 1
	ketemu := false
	namaLower := strings.ToLower(nama)

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		namaTengah := strings.ToLower(arrSalin[tengah].nama)
		if namaTengah == namaLower {
			fmt.Printf("%-6s %-25s %-12s %-20s\n",
				arrSalin[tengah].id,
				arrSalin[tengah].nama,
				arrSalin[tengah].tanggal,
				arrSalin[tengah].bidangMinat)
			ketemu = true
			break
		} else if namaTengah < namaLower {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	if !ketemu {
		fmt.Println("Data tidak ditemukan.")
	}
}

func menuCari() {
	cetakGaris()
	fmt.Println("           PENCARIAN DATA PESERTA")
	cetakGaris()

	if nPeserta == 0 {
		fmt.Println("Belum ada data peserta.")
		return
	}

	fmt.Println("Metode Pencarian:")
	fmt.Println("  1. Sequential Search berdasarkan Nama")
	fmt.Println("  2. Sequential Search berdasarkan Bidang Minat")
	fmt.Println("  3. Binary Search berdasarkan Nama")
	pilihan := bacaInput("Pilih metode: ")

	switch pilihan {
	case "1":
		kata := bacaInput("Masukkan nama yang dicari: ")
		sequentialSearchNama(kata)
	case "2":
		fmt.Println()
		tampilkanBidang()
		kata := bacaInput("Masukkan nama bidang minat: ")
		sequentialSearchBidang(kata)
	case "3":
		kata := bacaInput("Masukkan nama yang dicari (harus tepat): ")
		binarySearchNama(daftarPeserta, nPeserta, kata)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// ============================================================
// FITUR D: PENGURUTAN
// ============================================================

func selectionSortByID(arr *DaftarPeserta, n int) {
	var i, j, idxMin int
	var temp Peserta
	i = 0
	for i <= n-1 {
		idxMin = i
		j = i + 1
		for j < n {
			if arr[j].id < arr[idxMin].id {
				idxMin = j
			}
			j = j + 1
		}
		temp = arr[idxMin]
		arr[idxMin] = arr[i]
		arr[i] = temp
		i = i + 1
	}
}

func insertionSortNamaAsc(arr *DaftarPeserta, n int) {
	var temp Peserta
	var i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = arr[j]
		for j > 0 && strings.ToLower(temp.nama) < strings.ToLower(arr[j-1].nama) {
			arr[j] = arr[j-1]
			j = j - 1
		}
		arr[j] = temp
		i = i + 1
	}
}

func menuUrut() {
	cetakGaris()
	fmt.Println("           PENGURUTAN DATA PESERTA")
	cetakGaris()

	if nPeserta == 0 {
		fmt.Println("Belum ada data peserta.")
		return
	}

	fmt.Println("Metode Pengurutan:")
	fmt.Println("  1. Selection Sort berdasarkan ID Pendaftaran")
	fmt.Println("  2. Insertion Sort berdasarkan Nama (Alfabetis)")
	pilihan := bacaInput("Pilih metode: ")

	switch pilihan {
	case "1":
		selectionSortByID(&daftarPeserta, nPeserta)
		simpanCSV()
		fmt.Println("\nData berhasil diurutkan berdasarkan ID (Selection Sort)!")
		fmt.Println()
		tampilkanSemuaPeserta()
	case "2":
		insertionSortNamaAsc(&daftarPeserta, nPeserta)
		simpanCSV()
		fmt.Println("\nData berhasil diurutkan berdasarkan Nama A-Z (Insertion Sort)!")
		fmt.Println()
		tampilkanSemuaPeserta()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// ============================================================
// FITUR E: STATISTIK
// ============================================================

func tampilkanStatistik() {
	cetakGaris()
	fmt.Println("        STATISTIK PENDAFTAR KURSUS")
	cetakGaris()

	fmt.Printf("%-25s %s\n", "Bidang Minat", "Jumlah Pendaftar")
	cetakGarisTipis()

	i := 0
	for i < nBidang {
		bar := ""
		j := 0
		for j < daftarBidang[i].jumlah {
			bar = bar + "█"
			j = j + 1
		}
		fmt.Printf("%-25s %d  %s\n",
			daftarBidang[i].namaBidang,
			daftarBidang[i].jumlah,
			bar)
		i = i + 1
	}

	cetakGarisTipis()
	fmt.Printf("Total Peserta Aktif: %d\n", nPeserta)
}

// ============================================================
// MENU MANAJEMEN PESERTA
// ============================================================

func menuManajemenPeserta() {
	for {
		cetakGaris()
		fmt.Println("        MANAJEMEN DATA PESERTA")
		cetakGaris()
		fmt.Println("  1. Tambah Peserta")
		fmt.Println("  2. Ubah Data Peserta")
		fmt.Println("  3. Hapus Peserta")
		fmt.Println("  4. Lihat Semua Peserta")
		fmt.Println("  0. Kembali ke Menu Utama")
		cetakGarisTipis()
		pilihan := bacaInput("Pilihan: ")

		switch pilihan {
		case "1":
			tambahPeserta()
		case "2":
			ubahPeserta()
		case "3":
			hapusPeserta()
		case "4":
			cetakGaris()
			fmt.Println("         DAFTAR SEMUA PESERTA")
			cetakGaris()
			tampilkanSemuaPeserta()
		case "0":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		fmt.Println()
	}
}

// ============================================================
// MAIN
// ============================================================

func tampilkanHeader() {
	cetakGaris()
	fmt.Println("   KURSUS IN - Sistem Pendaftaran Kursus Online")
	fmt.Println("          Algoritma dan Pemrograman 2")
	cetakGaris()
}

func main() {
	inisialisasiBidang()
	inisialisasiKursus()

	// Load data dari CSV saat program dibuka
	loadCSV()

	for {
		fmt.Println()
		tampilkanHeader()
		fmt.Println("  1. Manajemen Data Peserta (Tambah/Ubah/Hapus)")
		fmt.Println("  2. Cari Data Peserta")
		fmt.Println("  3. Urutkan Data Peserta")
		fmt.Println("  4. Statistik Pendaftar")
		fmt.Println("  5. Lihat Katalog Kursus")
		fmt.Println("  0. Keluar")
		cetakGarisTipis()
		pilihan := bacaInput("Pilihan: ")

		switch pilihan {
		case "1":
			menuManajemenPeserta()
		case "2":
			menuCari()
		case "3":
			menuUrut()
		case "4":
			tampilkanStatistik()
		case "5":
			cetakGaris()
			fmt.Println("           KATALOG KURSUS")
			cetakGaris()
			i := 0
			for i < nKursus {
				fmt.Printf("  [%s] %-25s | Bidang: %s\n",
					daftarKursus[i].idKursus,
					daftarKursus[i].namaKursus,
					daftarKursus[i].bidang)
				i = i + 1
			}
		case "0":
			fmt.Println("Terima kasih telah menggunakan KursusIn!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
