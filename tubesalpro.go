package main

import (
	"fmt"
	"time"
)

const NMAX int = 1000

type bahan struct {
	nama              string
	jumlah            int
	tahunKadaluarsa   int
	bulanKadaluarsa   int
	tanggalKadaluarsa int
}

type tabBahan [NMAX]bahan

func main() {
	var data tabBahan
	var pilihan, n int

	for {
		menu()
		fmt.Scan(&pilihan)
		fmt.Scanf("\n") // bersihkan newline setelah input pilihan

		switch pilihan {
		case 1:
			show(data, n)
			cekKadaluarsa(&data, n)
		case 2:
			input(&data, &n)
		case 3:
			update(&data, &n)
		case 4:
			deleteData(&data, &n)
		case 5:
			search(&data, n)
		case 6:
			cekKadaluarsa(&data, n)
		case 7:
			fmt.Println("Keluar dari program")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func menu() {
	fmt.Println("\n====== MENU ======")
	fmt.Println("1. INFO")
	fmt.Println("2. Input")
	fmt.Println("3. Update")
	fmt.Println("4. Delete")
	fmt.Println("5. Pencarian")
	fmt.Println("6. Peringatan Kadaluarsa")
	fmt.Println("7. Keluar")
	fmt.Print("Masukkan Pilihan: ")
}

func input(A *tabBahan, n *int) {
	var nama string
	var jumlah, tahun, bulan, tanggal int
	lanjutkan := true

	for lanjutkan && *n < NMAX {
		fmt.Println("Masukkan nama bahan (atau ketik 'none' untuk berhenti):")
		fmt.Scanln(&nama)

		if nama == "none" {
			lanjutkan = false
		} else {
			fmt.Println("Masukkan jumlah bahan:")
			fmt.Scan(&jumlah)
			fmt.Scanf("\n")

			fmt.Println("Masukkan tahun kadaluarsa (YYYY):")
			fmt.Scan(&tahun)
			fmt.Scanf("\n")

			fmt.Println("Masukkan bulan kadaluarsa (1-12):")
			fmt.Scan(&bulan)
			fmt.Scanf("\n")

			fmt.Println("Masukkan tanggal kadaluarsa (1-31):")
			fmt.Scan(&tanggal)
			fmt.Scanf("\n")

			(*A)[*n].nama = nama
			(*A)[*n].jumlah = jumlah
			(*A)[*n].tahunKadaluarsa = tahun
			(*A)[*n].bulanKadaluarsa = bulan
			(*A)[*n].tanggalKadaluarsa = tanggal
			*n++

			if *n >= NMAX {
				fmt.Println("Batas maksimum bahan tercapai.")
				lanjutkan = false
			}
		}
	}
}

func show(A tabBahan, n int) {
	fmt.Println("NO    Nama              Jumlah      Kadaluarsa")
	fmt.Println("---------------------------------------------------")
	for i := 0; i < n; i++ {
		if A[i].nama != "" {
			fmt.Printf("%-6d %-18s %-10d %04d-%02d-%02d\n", i+1, A[i].nama, A[i].jumlah,
				A[i].tahunKadaluarsa, A[i].bulanKadaluarsa, A[i].tanggalKadaluarsa)
		}
	}
}

func update(A *tabBahan, n *int) {
	var p, tahunBaru, bulanBaru, tanggalBaru, jumlahBaru int
	var namaBaru string

	if *n == 0 {
		fmt.Println("Data masih kosong, tidak bisa update.")
	} else {
		show(*A, *n)

		fmt.Println("Edit data ke? ")
		fmt.Scan(&p)
		fmt.Scanf("\n")

		if p < 1 || p > *n {
			fmt.Println("Data tidak ditemukan")
		} else {
			fmt.Println("Masukkan nama baru:")
			fmt.Scanln(&namaBaru)

			fmt.Println("Masukkan jumlah baru:")
			fmt.Scan(&jumlahBaru)
			fmt.Scanf("\n")

			fmt.Println("Masukkan tahun kadaluarsa baru (YYYY):")
			fmt.Scan(&tahunBaru)
			fmt.Scanf("\n")

			fmt.Println("Masukkan bulan kadaluarsa baru (1-12):")
			fmt.Scan(&bulanBaru)
			fmt.Scanf("\n")

			fmt.Println("Masukkan tanggal kadaluarsa baru (1-31):")
			fmt.Scan(&tanggalBaru)
			fmt.Scanf("\n")

			(*A)[p-1].nama = namaBaru
			(*A)[p-1].jumlah = jumlahBaru
			(*A)[p-1].tahunKadaluarsa = tahunBaru
			(*A)[p-1].bulanKadaluarsa = bulanBaru
			(*A)[p-1].tanggalKadaluarsa = tanggalBaru
		}
	}
}

func deleteData(A *tabBahan, n *int) {
	var p int

	if *n == 0 {
		fmt.Println("Data masih kosong, tidak bisa hapus.")
	} else {
		show(*A, *n)

		fmt.Println("Hapus data ke? ")
		fmt.Scan(&p)
		fmt.Scanf("\n")

		if p < 1 || p > *n {
			fmt.Println("Data tidak ditemukan")
		} else {
			for i := p - 1; i < *n-1; i++ {
				(*A)[i] = (*A)[i+1]
			}
			*n--
			fmt.Println("Data berhasil dihapus")
		}
	}
}

func cekKadaluarsa(A *tabBahan, n int) {
	current := time.Now()
	yearNow := current.Year()
	monthNow := int(current.Month())
	dayNow := current.Day()

	fmt.Printf("\nTanggal sekarang: %04d-%02d-%02d\n", yearNow, monthNow, dayNow)
	fmt.Println("=== Peringatan Kadaluarsa ===")

	for i := 0; i < n; i++ {
		b := (*A)[i]

		if b.tahunKadaluarsa < yearNow ||
			(b.tahunKadaluarsa == yearNow && b.bulanKadaluarsa < monthNow) ||
			(b.tahunKadaluarsa == yearNow && b.bulanKadaluarsa == monthNow && b.tanggalKadaluarsa < dayNow) {

			fmt.Printf("Peringatan: %v sudah kadaluarsa! Kadaluarsa: %04d-%02d-%02d\n",
				b.nama, b.tahunKadaluarsa, b.bulanKadaluarsa, b.tanggalKadaluarsa)

		} else if b.tahunKadaluarsa == yearNow && b.bulanKadaluarsa == monthNow {

			selisih := b.tanggalKadaluarsa - dayNow
			if selisih == 0 {
				fmt.Printf("Peringatan: %v hari ini kadaluarsa! Kadaluarsa: %04d-%02d-%02d\n",
					b.nama, b.tahunKadaluarsa, b.bulanKadaluarsa, b.tanggalKadaluarsa)
			} else if selisih > 0 && selisih <= 7 {
				fmt.Printf("Peringatan: %v kadaluarsa dalam %d hari. Kadaluarsa: %04d-%02d-%02d\n",
					b.nama, selisih, b.tahunKadaluarsa, b.bulanKadaluarsa, b.tanggalKadaluarsa)
			}
		}
	}
}

func sequentialSearch(A tabBahan, n int, keyword string) bool {
	for i := 0; i < n; i++ {
		if A[i].nama == keyword {
			return true
		}
	}
	return false
}

func search(A *tabBahan, n int) {
	var keyword string
	var hitung int

	if n == 0 {
		fmt.Println("Data masih kosong, tidak bisa mencari.")
	} else {
		fmt.Println("Masukkan nama bahan yang dicari:")
		fmt.Scanln(&keyword)

		if sequentialSearch(*A, n, keyword) {
			fmt.Println("Bahan ditemukan.")
			fmt.Println("Detail bahan:")

			for i := 0; i < n; i++ {
				if (*A)[i].nama == keyword {
					fmt.Printf("- %s | Jumlah: %d | Kadaluarsa: %04d-%02d-%02d\n",
						(*A)[i].nama, (*A)[i].jumlah,
						(*A)[i].tahunKadaluarsa, (*A)[i].bulanKadaluarsa, (*A)[i].tanggalKadaluarsa)
					hitung++
				}
			}
			fmt.Printf("Total ditemukan: %d item.\n", hitung)
		} else {
			fmt.Println("Bahan tidak ditemukan.")
		}
	}
}
