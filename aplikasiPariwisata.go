package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const maxWisata = 100
const MaxFasilitas = 5

type DestinasiWisata struct {
	id              int
	nama            string
	kategori        string
	hargaTiket      int
	jarak           int
	rating          float64
	jumlahRating    int
	fasilitasWisata [MaxFasilitas]string
}

var TempatWisata [maxWisata]DestinasiWisata
var JumlahWisata int

func main() {
	for {
		fmt.Println("======= Menu utama ========== ")
		fmt.Println("1. Masuk sebagai Admin ")
		fmt.Println("2. Masuk sebagain Pengguna ")
		fmt.Println("3. Keluar aplikasi")
		fmt.Print("pilih menu : ")
		var pilihanRole int
		fmt.Scanln(&pilihanRole)

		if pilihanRole == 3 {
			fmt.Println("Terimakasih telah menggunakan aplikasi ini")
			break
		}

		switch pilihanRole {
		case 1:
			for {
				fmt.Println("========== Anda masuk sebagai admin ==============")
				fmt.Println("1. Tambah Tempat Wisata")
				fmt.Println("2. Edit Tempat Wisata")
				fmt.Println("3. Hapus Tempat Wisata")
				fmt.Println("4. Lihat semua Tempat Wisata")
				fmt.Println("0. Logout")
				fmt.Print("Pilih menu: ")

				var pilihanAdmin int
				fmt.Scanln(&pilihanAdmin)

				switch pilihanAdmin {
				case 1:
					tambahDestinasi()
				case 2:
					editDestinasi()
				case 3:
					hapusDestinasi()
				case 4:
					lihatSemuaDestinasi()
				case 0:
					fmt.Println("Anda keluar dari admin")
					break
				default:
					fmt.Println("pilihan tidak valid")
				}
				if pilihanAdmin == 0 {
					break
				}
			}
		case 2:
			for {
				fmt.Println("======= Anda masuk sebagai pengguna ==========")
				fmt.Println("1. Lihat Semua Tempat Wisata")
				fmt.Println("2. Urutkan Tempat Wisata")
				fmt.Println("3. Cari tempat wisata")
				fmt.Println("4. Beri Rating Tempat Wisata")
				fmt.Println("0. Logout")
				fmt.Print("Pilih menu: ")

				var pilihanPengguna int
				fmt.Scanln(&pilihanPengguna)

				switch pilihanPengguna {
				case 1:
					lihatSemuaDestinasi()
				case 2:
					urutkanDestinasi()
				case 3:
					cariDestinasi()
				case 4:
					beriRating()
				case 0:
					fmt.Println("Keluar dari pengguna")
					break
				default:
					fmt.Println("pilihan tidak valid")
				}
				if pilihanPengguna == 0 {
					break
				}
			}
		}
	}
}

// fungsi untuk membaca input dengan spasi
func bacaInput(inputan string) string {
	fmt.Print(inputan)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input) // menghapus spasi dan newline
}

// Untuk Admin
func tambahDestinasi() {
	if JumlahWisata >= maxWisata {
		fmt.Println("Kapasitas penuh, wisata tidak bisa  ditambahkan")
		return
	}

	var Wisata DestinasiWisata
	JumlahWisata++
	Wisata.id = JumlahWisata

	Wisata.nama = bacaInput("Masukan nama Destinasi Wisata : ")

	Wisata.kategori = bacaInput("Masukan kategori destinasi wisata : ")

	fmt.Print("Masukan Harga tiket masuk destinasi wisata : ")
	fmt.Scanln(&Wisata.hargaTiket)

	fmt.Print("Masukan jarak destinasi wisata : ")
	fmt.Scanln(&Wisata.jarak)

	fmt.Println("Masukan fasilitas yang ada di destinasi wisata : ")
	for i := 0; i < MaxFasilitas; i++ {
		fasilitas := bacaInput(fmt.Sprintf("Fasilitas %d: ", i+1))
		if fasilitas == "" {
			break
		}
		Wisata.fasilitasWisata[i] = fasilitas
	}

	Wisata.rating = 0
	Wisata.jumlahRating = 0
	TempatWisata[JumlahWisata-1] = Wisata

	fmt.Println("Tempat wisata berhasil ditambahkan")
}

func editDestinasi() {
	var nama string
	nama = bacaInput("Masukkan nama destinasi wisata yang akan diubah : ")

	var index = -1
	for i := 0; i < JumlahWisata; i++ {
		if strings.EqualFold(TempatWisata[i].nama, nama) {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Tempat wisata tidak ditemukan")
		return
	}

	destinasi := &TempatWisata[index]
	fmt.Printf("Mengedit destinasi : %s\n", destinasi.nama)

	namaBaru := bacaInput("Masukkan nama destinasi wisata yang baru (kosongkan jika tidak akan diubah ) : ")
	if namaBaru != "" {
		destinasi.nama = namaBaru
	}

	kategoriBaru := bacaInput("Masukkan Kategori baru destinasi wisata (kosongkan jika tidak diubah ) : ")
	if kategoriBaru != "" {
		destinasi.kategori = kategoriBaru
	}

	fmt.Print("Masukkan harga tiket baru (0 untuk jika tidak ada yang diubah) : ")
	var tiketBaru int
	fmt.Scanln(&tiketBaru)
	if tiketBaru > 0 {
		destinasi.hargaTiket = tiketBaru
	}

	fmt.Print("Masukkan jarak baru (0 jika tidak ada yang diubah) : ")
	var jarakBaru int
	fmt.Scanln(&jarakBaru)
	if jarakBaru > 0 {
		destinasi.jarak = jarakBaru
	}

	fmt.Print("Masukkan fasilitas baru (kosongkan jika tidak ada yang diubah) :  ")
	fmt.Println()
	for i := 0; i < MaxFasilitas; i++ {
		fasilitasBaru := bacaInput(fmt.Sprintf("Fasilitas %d: ", i+1))
		if fasilitasBaru != "" {
			destinasi.fasilitasWisata[i] = fasilitasBaru
		}
	}

	fmt.Println("Destinasi berhasil diperbarui")
}

func hapusDestinasi() {
	var nama string
	nama = bacaInput("Masukan nama destinasi wisata yang akan dihapus : ")

	var index = -1
	for i := 0; i < JumlahWisata; i++ {
		if strings.EqualFold(TempatWisata[i].nama, nama) {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Tempat wisata tidak ditemukan")
		return
	}

	for i := index; i < JumlahWisata-1; i++ {
		TempatWisata[i] = TempatWisata[i+1]
	}
	JumlahWisata--
	fmt.Println("Tempat wisata berhasil dihapus")
}

// Untuk Pengguna
func beriRating() {
	if JumlahWisata == 0 {
		fmt.Println("Belum ada destinasi yang tersedia untuk diberi rating.")
		return
	}

	nama := bacaInput("Masukan nama destinasi wisata yang akan diberi rating : ")

	var index = -1
	for i := 0; i < JumlahWisata; i++ {
		if strings.EqualFold(TempatWisata[i].nama, nama) {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Tempat wisata tidak ditemukan.")
		return
	}

	destinasi := &TempatWisata[index]
	fmt.Printf("Anda memberikan rating untuk destinasi: %s (Rating saat ini: %.2f)", destinasi.nama, destinasi.rating)

	var ratingBaru float64
	fmt.Print("Masukkan rating baru (1.0 - 5.0): ")
	fmt.Scanln(&ratingBaru)

	if ratingBaru < 1.0 || ratingBaru > 5.0 {
		fmt.Println("Rating tidak valid. Harap masukkan nilai antara 1.0 dan 5.0.")
		return
	}

	destinasi.rating = ((destinasi.rating * float64(destinasi.jumlahRating)) + ratingBaru) / float64(destinasi.jumlahRating+1)
	destinasi.jumlahRating++
	fmt.Println("Terima kasih, rating Anda telah disimpan.")
}

func cariDestinasi() {
	if JumlahWisata == 0 {
		fmt.Println("Belum ada destinasi wisata yang tersedia untuk dicari")
		return
	}
	var pilih int
	fmt.Println("Cari destinasi berdasarkan : ")
	fmt.Println("1. Nama destinasi wisata ")
	fmt.Println("2. Kategori destinasi wisata ")
	fmt.Print("pilih : ")
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		nama := bacaInput("Masukan nama destinasi wisata yang ingin dicari : ")
		fmt.Println("Hasil pencarian destinasi berdasarkan nama:")
		found := false
		for _, destinasi := range TempatWisata[:JumlahWisata] {
			if strings.Contains(strings.ToLower(destinasi.nama), strings.ToLower(nama)) {
				fmt.Printf("%d. %s (Kategori: %s, Harga: Rp%d, Jarak: %d km, Rating: %.2f)\n",
					destinasi.id, destinasi.nama, destinasi.kategori, destinasi.hargaTiket, destinasi.jarak, destinasi.rating)
				fmt.Printf(" Fasilitas: %s\n", strings.Join(destinasi.fasilitasWisata[:], ", "))
				found = true
			}
		}
		if !found {
			fmt.Println("Tidak ada destinasi yang ditemukan")
		}
	case 2:
		kategori := bacaInput("Masukan kategori destinasi wisata yang ingin dicari : ")
		fmt.Println("Hasil pencarian destinasi berdasarkan kategori:")
		found := false
		for _, destinasi := range TempatWisata[:JumlahWisata] {
			if strings.EqualFold(destinasi.kategori, kategori) {
				fmt.Printf("%d. %s (Kategori: %s, Harga: Rp%d, Jarak: %d km, Rating: %.2f)\n",
					destinasi.id, destinasi.nama, destinasi.kategori, destinasi.hargaTiket, destinasi.jarak, destinasi.rating)
				fmt.Printf(" Fasilitas: %s\n", strings.Join(destinasi.fasilitasWisata[:], ", "))
				found = true
			}
		}
		if !found {
			fmt.Println("Tidak ada destinasi yang ditemukan.")
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func urutkanDestinasi() {
	if JumlahWisata == 0 {
		fmt.Println("Belum ada destinasi wisata yang tersedia untuk diurutkan ")
		return
	}
	var pilihan int
	fmt.Println("Urutkan destinasi wisata berdasarkan : ")
	fmt.Println("1. Harga Tiket  ") // ascending
	fmt.Println("2. Jarak  ")       // ascending
	fmt.Println("3. Rating  ")      // descending
	fmt.Print(" pilih : ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		for i := 0; i < JumlahWisata-1; i++ { // mengurutkan berdasarkan harga tiket menggunakan selection sort
			minIndx := i
			for j := i + 1; j < JumlahWisata; j++ {
				if TempatWisata[j].hargaTiket < TempatWisata[minIndx].hargaTiket {
					minIndx = j
				}
			}
			if minIndx != i {
				TempatWisata[i], TempatWisata[minIndx] = TempatWisata[minIndx], TempatWisata[i]
			}
		}
		fmt.Println("Tempat wisata berhasil diurutkan berdasarkan harga tiket (ascending)")
	case 2:

		for i := 0; i < JumlahWisata-1; i++ { // mengurutkan berdasarkan jarak menggunakan selection sort
			minIndx := i
			for j := i + 1; j < JumlahWisata; j++ {
				if TempatWisata[j].jarak < TempatWisata[minIndx].jarak {
					minIndx = j
				}
			}
			if minIndx != i {
				TempatWisata[i], TempatWisata[minIndx] = TempatWisata[minIndx], TempatWisata[i]
			}
		}
		fmt.Println("Tempat wisata berhasil diurutkan berdasarkan jarak (ascending)")

	case 3:
		for i := 1; i < JumlahWisata; i++ { // mengurutkan berdasarkan rating menggunakan insertion sort
			key := TempatWisata[i]
			j := i - 1

			for j >= 0 && TempatWisata[j].rating < key.rating {
				TempatWisata[j+1] = TempatWisata[j]
				j--
			}
			TempatWisata[j+1] = key
		}
		fmt.Println("Tempat wisata berhasil diurutkan berdasarkan rating (descending)")
	default:
		fmt.Println("Pilihan tidak valid")
		return
	}
	lihatSemuaDestinasi()
}

// Untuk Semua

func lihatSemuaDestinasi() {
	if JumlahWisata == 0 {
		fmt.Println("Belum ada destinasi yang tersedia")
		return
	}

	fmt.Println("Daftar tempat wisata: ")
	for i := 0; i < JumlahWisata; i++ {
		destinasi := TempatWisata[i]
		fmt.Printf("ID destinasi wisata : %d\n ", destinasi.id)
		fmt.Printf("Nama destinasi wisata : %s\n ", destinasi.nama)
		fmt.Printf("Kategori destinasi wisata : %s\n ", destinasi.kategori)
		fmt.Printf("Harga tiket destinasi wisata : %d\n ", destinasi.hargaTiket)
		fmt.Printf("Jarak destinasi wisata : %d\n ", destinasi.jarak)
		fmt.Printf("Rating destinasi wisata : %.2f\n ", destinasi.rating)
		fmt.Printf("Fasilitas destinasi wisata : %s\n\n", strings.Join(destinasi.fasilitasWisata[:], ","))
	}
}
