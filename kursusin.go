package main

import (
	"fmt"
	"sort"
)

type Peserta struct {
	ID      int
	Nama    string
	Bidang  string
	Tanggal string
	Aktif   bool
}

var data [100]Peserta
var jumlah int

func tambahPeserta() {

	var p Peserta

	fmt.Println()
	fmt.Println("=== TAMBAH PESERTA ===")

	fmt.Print("ID               : ")
	fmt.Scanln(&p.ID)

	fmt.Print("Nama             : ")
	fmt.Scanln(&p.Nama)

	fmt.Print("Bidang Minat     : ")
	fmt.Scanln(&p.Bidang)

	fmt.Print("Tanggal Daftar   : ")
	fmt.Scanln(&p.Tanggal)

	p.Aktif = true

	data[jumlah] = p
	jumlah++

	fmt.Println("Data berhasil ditambahkan")
}

func tampilData() {

	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	fmt.Println()
	fmt.Println("=== DATA PESERTA ===")

	for i := 0; i < jumlah; i++ {

		fmt.Println("Data ke-", i+1)
		fmt.Println("ID       :", data[i].ID)
		fmt.Println("Nama     :", data[i].Nama)
		fmt.Println("Bidang   :", data[i].Bidang)
		fmt.Println("Tanggal  :", data[i].Tanggal)
		fmt.Println("Aktif    :", data[i].Aktif)
		fmt.Println()
	}
}

func ubahPeserta() {

	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	var id int
	var ketemu bool

	fmt.Print("Masukkan ID yang ingin diubah : ")
	fmt.Scanln(&id)

	for i := 0; i < jumlah; i++ {

		if data[i].ID == id {

			fmt.Print("Nama baru       : ")
			fmt.Scanln(&data[i].Nama)

			fmt.Print("Bidang baru     : ")
			fmt.Scanln(&data[i].Bidang)

			fmt.Print("Tanggal baru    : ")
			fmt.Scanln(&data[i].Tanggal)

			fmt.Println("Data berhasil diubah")

			ketemu = true
			break
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusPeserta() {

	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	var id int
	var ketemu bool

	fmt.Print("Masukkan ID yang ingin dihapus : ")
	fmt.Scanln(&id)

	for i := 0; i < jumlah; i++ {

		if data[i].ID == id {

			for j := i; j < jumlah-1; j++ {
				data[j] = data[j+1]
			}

			jumlah--

			fmt.Println("Data berhasil dihapus")

			ketemu = true
			break
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

func sequentialSearchNama() {

	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	var cari string
	var ketemu bool

	fmt.Print("Masukkan nama yang dicari : ")
	fmt.Scanln(&cari)

	for i := 0; i < jumlah; i++ {

		if data[i].Nama == cari {

			fmt.Println()
			fmt.Println("Data ditemukan")
			fmt.Println("ID      :", data[i].ID)
			fmt.Println("Nama    :", data[i].Nama)
			fmt.Println("Bidang  :", data[i].Bidang)
			fmt.Println("Tanggal :", data[i].Tanggal)

			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

func sequentialSearchBidang() {

	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	var cari string
	var ketemu bool

	fmt.Print("Masukkan bidang yang dicari : ")
	fmt.Scanln(&cari)

	for i := 0; i < jumlah; i++ {

		if data[i].Bidang == cari {

			fmt.Println()
			fmt.Println("ID      :", data[i].ID)
			fmt.Println("Nama    :", data[i].Nama)
			fmt.Println("Bidang  :", data[i].Bidang)
			fmt.Println("Tanggal :", data[i].Tanggal)

			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

func selectionSortID() {

	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i := 0; i < jumlah-1; i++ {

		min := i

		for j := i + 1; j < jumlah; j++ {

			if data[j].ID < data[min].ID {
				min = j
			}
		}

		data[i], data[min] = data[min], data[i]
	}

	fmt.Println("Data berhasil diurutkan berdasarkan ID")
}

func insertionSortNama() {

	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i := 1; i < jumlah; i++ {

		temp := data[i]
		j := i - 1

		for j >= 0 && data[j].Nama > temp.Nama {

			data[j+1] = data[j]
			j--
		}

		data[j+1] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan Nama")
}

func binarySearchNama() {

	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	sort.Slice(data[:jumlah], func(i, j int) bool {
		return data[i].Nama < data[j].Nama
	})

	var cari string

	fmt.Print("Masukkan nama yang dicari : ")
	fmt.Scanln(&cari)

	left := 0
	right := jumlah - 1

	for left <= right {

		mid := (left + right) / 2

		if data[mid].Nama == cari {

			fmt.Println()
			fmt.Println("Data ditemukan")
			fmt.Println("ID      :", data[mid].ID)
			fmt.Println("Nama    :", data[mid].Nama)
			fmt.Println("Bidang  :", data[mid].Bidang)
			fmt.Println("Tanggal :", data[mid].Tanggal)

			return

		} else if cari < data[mid].Nama {

			right = mid - 1

		} else {

			left = mid + 1
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func statistik() {

	if jumlah == 0 {
		fmt.Println("Data kosong")
		return
	}

	var aktif int

	fmt.Println()
	fmt.Println("    KursusIn    ")
	fmt.Println()

	for i := 0; i < jumlah; i++ {

		hitung := 1
		sudahAda := false

		for j := 0; j < i; j++ {

			if data[i].Bidang == data[j].Bidang {
				sudahAda = true
			}
		}

		if !sudahAda {

			for j := i + 1; j < jumlah; j++ {

				if data[i].Bidang == data[j].Bidang {
					hitung++
				}
			}

			fmt.Println(data[i].Bidang, ":", hitung)
		}

		if data[i].Aktif {
			aktif++
		}
	}

	fmt.Println()
	fmt.Println("Total Peserta Aktif :", aktif)
}

func main() {

	var pilihan int

	for {

		fmt.Println()
		fmt.Println("    KursusIn    ")
		fmt.Println("1. Tambah Peserta")
		fmt.Println("2. Tampilkan Data")
		fmt.Println("3. Ubah Peserta")
		fmt.Println("4. Hapus Peserta")
		fmt.Println("5. Sequential Search Nama")
		fmt.Println("6. Sequential Search Bidang")
		fmt.Println("7. Selection Sort ID")
		fmt.Println("8. Insertion Sort Nama")
		fmt.Println("9. Binary Search Nama")
		fmt.Println("10. Statistik")
		fmt.Println("11. Keluar")

		fmt.Print("Pilih menu : ")
		fmt.Scanln(&pilihan)

		if pilihan < 1 || pilihan > 11 {
			fmt.Println("Menu tidak tersedia")
			continue
		}

		switch pilihan {

		case 1:
			tambahPeserta()

		case 2:
			tampilData()

		case 3:
			ubahPeserta()

		case 4:
			hapusPeserta()

		case 5:
			sequentialSearchNama()

		case 6:
			sequentialSearchBidang()

		case 7:
			selectionSortID()

		case 8:
			insertionSortNama()

		case 9:
			binarySearchNama()

		case 10:
			statistik()

		case 11:
			fmt.Println("Program selesai")
			return
		}
	}
}
