package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const NMAX int = 1000 // batas jumlah array

type bahan struct {
	nama          string
	status        string
	jumlah        int
	inDate        time.Time
	kadaluarsa    int       // dalam jumlah hari sejak inDate (dalam bentuk hari)
	tglKadaluarsa time.Time // tanggal kadaluarsa dalam bentuk tanggal
}

// mendeklarasikan array untuk bahan
type tabBahan [NMAX]bahan

func main() {
	var data tabBahan
	var p, n int

	// deklarasi reader
	reader := bufio.NewReader(os.Stdin)

	dumyData(&data, &n) // untuk bahan test

	for {
		menu() // memanggil fungsi menu
		fmt.Scan(&p)
		reader.ReadString('\n') // untuk membaca newline

		switch p {
		case 1:
			show(data, n)
		case 2:
			input(&data, &n)
		case 3:
			update(&data, n)
		case 4:
			search(&data, n)
		case 5:
			sorting(&data, n)
		case 6:
			delete(&data, &n)
		case 7:
			fmt.Println("Keluar dari program")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}

}

// untuk menampilkan menu
func menu() {
	fmt.Println("\n===== MENU =====")
	fmt.Println("1. INFO")
	fmt.Println("2. Input")
	fmt.Println("3. Update")
	fmt.Println("4. Search")
	fmt.Println("5. Sorting")
	fmt.Println("6. Delete")
	fmt.Println("7. Keluar")
	fmt.Print("Masukkan Pilihan: ")
}

// untuk membuat data baru
func input(A *tabBahan, n *int) {
	var jumlah, kadaluarsa int
	var nama string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan Nama Bahan (ketik 'none' untuk berhenti): ")
	input, _ := reader.ReadString('\n') // ini digunakan agar dapat menginputkan multiple word ke array
	nama = strings.TrimSpace(input)

	for nama != "none" && *n <= NMAX {
		fmt.Print("Masukan Jumlah Bahan dan Kadaluarsa: ")
		fmt.Scan(&jumlah, &kadaluarsa) // kadaluarsa yang dibaca dalam bentuk hari, misalkan sekarang tanggal 21 dan diinputkan 20 maka tanggal kadaluarsa yang masuk adalah 20 hari dari tanggal 21
		reader.ReadString('\n') // digunakan agar tidak terjadi error
		A[*n].nama = nama
		A[*n].jumlah = jumlah
		A[*n].kadaluarsa = kadaluarsa
		A[*n].inDate = time.Now()
		A[*n].tglKadaluarsa = A[*n].inDate.AddDate(0, 0, A[*n].kadaluarsa)
		A[*n].status = cekStatus(A[*n])

		(*n)++

		fmt.Print("Masukan Nama Bahan (ketik 'none' untuk berhenti): ")
		input, _ = reader.ReadString('\n')
		nama = strings.TrimSpace(input)
	}
}

// menampilkan data
func show(A tabBahan, n int) {
	// Tampilkan tanggal sekarang dengan format dd-mm-yyyy
	fmt.Println("\nTanggal Sekarang:", time.Now().Format("02-01-2006"))
	fmt.Println(strings.Repeat("=", 70))

	// Header tabel
	fmt.Printf("%-4s | %-20s | %-8s | %-15s | %-20s\n", "NO", "Nama", "Jumlah", "Kadaluarsa", "Status")
	fmt.Println(strings.Repeat("-", 70))

	// Isi tabel
	for i := 0; i < n; i++ {
		if A[i].nama != "" {
			fmt.Printf(
				"%-4d | %-20s | %-8d | %-15s | %-20s\n",
				i+1,
				A[i].nama,
				A[i].jumlah,
				A[i].tglKadaluarsa.Format("02-01-2006"),
				A[i].status,
			)
		}
	}

	fmt.Println(strings.Repeat("=", 70))
}

// untuk mengedit data yang sudah ada
func update(A *tabBahan, n int) {
	var p, jumlahBaru, kadaluarsaBaru int
	var namaBaru string

	reader := bufio.NewReader(os.Stdin)

	if n < 0 {
		fmt.Println("Data Masih kosong")
	} else {
		show(*A, n)

		fmt.Print("Edit Data ke? ")
		fmt.Scan(&p)
		reader.ReadString('\n')

		if p < 1 || p > n {
			fmt.Println("Data tidak ditemukan")
			return
		}

		fmt.Print("Masukan Nama Bahan: ")
		input, _ := reader.ReadString('\n')
		namaBaru = strings.TrimSpace(input)

		fmt.Print("Masukan Jumlah Bahan dan Kadaluarsa: ")
		fmt.Scan(&jumlahBaru, &kadaluarsaBaru)

		A[p-1].nama = namaBaru
		A[p-1].jumlah = jumlahBaru
		A[p-1].kadaluarsa = kadaluarsaBaru
		A[p-1].inDate = time.Now()
		A[p-1].tglKadaluarsa = A[p-1].inDate.AddDate(0, 0, A[p-1].kadaluarsa)
		A[p-1].status = cekStatus(A[p-1])
	}

}

// untuk menghapus data
func delete(A *tabBahan, n *int) {
	var i, p int

	if *n < 0 {
		fmt.Println("Data masih kosong")
	} else {
		show(*A, *n)

		fmt.Print("Hapus Data ke? ")
		fmt.Scan(&p)

		if p < 1 || p > *n {
			fmt.Println("Data tidak ditemukan")
			return
		}

		for i = p - 1; i < *n-1; i++ {
			A[i] = A[i+1]
		}

		*n--
		fmt.Println("Data berhasil di hapus")
	}

}

// untuk mengecek apakah akan kadaluarsa atau tidak
func cekStatus(B bahan) string {
	var sisa int

	sisa = int(B.tglKadaluarsa.Sub(time.Now()).Hours() / 24)

	if sisa < 0 {
		return "Sudah Kadaluarsa"
	} else if sisa <= 1 {
		return "Segera Kadaluarsa"
	} else if sisa <= 3 {
		return "Akan Kadaluarsa"
	} else {
		return "Aman"
	}
}

func sequentialSearch(A tabBahan, n int, keyword string) bool {
	keyword = strings.ToLower(keyword) // menggunakan strings.ToLower agar tidak menjadi case sensitive
	for i := 0; i < n; i++ {
		if strings.Contains(strings.ToLower(A[i].nama), keyword) {
			return true
		}
	}
	return false
}

// Binary search pertama untuk status (data sudah diurutkan)
func BinarySearchFirstStatus(A tabBahan, n int, target string) int {
	low, high := 0, n-1
	first := -1
	target = strings.ToLower(target)

	for low <= high {
		mid := (low + high) / 2
		statusMid := strings.ToLower(A[mid].status)

		if statusMid == target {
			first = mid
			high = mid - 1
		} else if statusMid < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return first
}

// Binary search terakhir untuk status
func BinarySearchLastStatus(A tabBahan, n int, target string) int {
	low, high := 0, n-1
	last := -1
	target = strings.ToLower(target)

	for low <= high {
		mid := (low + high) / 2
		statusMid := strings.ToLower(A[mid].status)

		if statusMid == target {
			last = mid
			low = mid + 1
		} else if statusMid < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return last
}

func sortingStatus(A *tabBahan, n int) {
	var i, idx, pass int
	var temp bahan

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass

		for i < n {
			if strings.ToLower(A[i].status) < strings.ToLower(A[idx].status) {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}


func search(A *tabBahan, n int) {
	reader := bufio.NewReader(os.Stdin)
	lanjut := true

	for lanjut {
		if n == 0 {
			fmt.Println("Data kosong, tidak bisa mencari.")
		} else {
			fmt.Println("Pilih metode pencarian:")
			fmt.Println("1. Cari berdasarkan nama (sequential search)")
			fmt.Println("2. Cari berdasarkan status (binary search dengan sorting otomatis)")
			fmt.Print("Masukkan pilihan: ")

			var pilih int
			fmt.Scan(&pilih)
			fmt.Scanln() // buang newline

			if pilih == 1 {
				fmt.Print("Masukkan nama bahan yang dicari: ")
				input, _ := reader.ReadString('\n')
				keyword := strings.TrimSpace(input)

				if sequentialSearch(*A, n, keyword) {
					fmt.Println("Bahan ditemukan.")
					fmt.Println("Detail bahan:")

					fmt.Println(strings.Repeat("=", 70))
					fmt.Printf("%-4s | %-20s | %-8s | %-15s | %-20s\n", "NO", "Nama", "Jumlah", "Kadaluarsa", "Status")
					fmt.Println(strings.Repeat("-", 70))

					count := 0
					for i := 0; i < n; i++ {
						if strings.Contains(strings.ToLower((*A)[i].nama), strings.ToLower(keyword)) {
							fmt.Printf("%-4d | %-20s | %-8d | %-15s | %-20s\n", i+1, A[i].nama, A[i].jumlah, A[i].tglKadaluarsa.Format("02-01-2006"), A[i].status,
							)
							count++
						}
					}
					fmt.Println(strings.Repeat("=", 70))
					fmt.Printf("Total ditemukan: %d item.\n", count)
					fmt.Println(strings.Repeat("=", 70))
				} else {
					fmt.Println("Bahan tidak ditemukan.")
				}
			} else if pilih == 2 {
				sortingStatus(A, n) // sorting status

				fmt.Print("Masukkan status yang dicari (Aman, Akan Kadaluarsa, Segera Kadaluarsa, Sudah Kadaluarsa): ")
				input, _ := reader.ReadString('\n')
				target := strings.TrimSpace(input)

				first := BinarySearchFirstStatus(*A, n, target)
				last := BinarySearchLastStatus(*A, n, target)

				if first == -1 || last == -1 || first > last {
					fmt.Println("Status tidak ditemukan.")
				} else {
					fmt.Println(strings.Repeat("=", 70))
					fmt.Printf("%-4s | %-20s | %-8s | %-15s | %-20s\n", "NO", "Nama", "Jumlah", "Kadaluarsa", "Status")
					fmt.Println(strings.Repeat("-", 70))

					for i := first; i <= last; i++ {
						fmt.Printf("%-4d | %-20s | %-8d | %-15s | %-20s\n",
							i+1, A[i].nama, A[i].jumlah, A[i].tglKadaluarsa.Format("02-01-2006"), A[i].status)
					}
					fmt.Println(strings.Repeat("=", 70))
					fmt.Printf("Total ditemukan: %d item.\n", last-first+1)
					fmt.Println(strings.Repeat("=", 70))
				}
			} else {
				fmt.Println("Pilihan tidak valid")
			}
		}

		fmt.Print("Mau cari lagi? (y/n): ")
		jawaban, _ := reader.ReadString('\n')
		jawaban = strings.TrimSpace(jawaban)
		if strings.ToLower(jawaban) != "y" {
			lanjut = false
		}
	}
}

func menuSorting() {
	fmt.Println("\n===== Sorting =====")
	fmt.Println("1. A-Z")
	fmt.Println("2. Z-A")
	fmt.Println("3. Terbanyak")
	fmt.Println("4. Paling Sedikit")
	fmt.Println("5. Paling Lama (Kadaluarsa)")
	fmt.Println("6. Paling Dekat (Kadaluarsa)")
	fmt.Print("Masukkan Pilihan: ")
}

func sorting(A *tabBahan, n int) {
	var i, idx, pass int
	var temp bahan
	var p int

	menuSorting()
	fmt.Scan(&p)

	if p == 1 { // sorting berdasarkan nama dari A
		pass = 1

		for pass < n {
			idx = pass - 1
			i = pass

			for i < n {
				if strings.ToLower(A[i].nama) < strings.ToLower(A[idx].nama) {
					idx = i
				}
				i = i + 1
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass = pass + 1
		}
	} else if p == 2 {
		pass = 1

		for pass < n {
			idx = pass - 1
			i = pass

			for i < n {
				if strings.ToLower(A[i].nama) > strings.ToLower(A[idx].nama) {
					idx = i
				}
				i = i + 1
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass = pass + 1
		}
	} else if p == 3 { // sorting berdasarkan jumlah dari yang paling banyak
		pass = 1

		for pass < n {
			idx = pass - 1
			i = pass

			for i < n {
				if A[i].jumlah > A[idx].jumlah {
					idx = i
				}
				i = i + 1
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass = pass + 1
		}
	} else if p == 4 { // sorting berdasarkan jumlah dari yang paling sedikit
		pass = 1

		for pass < n {
			idx = pass - 1
			i = pass

			for i < n {
				if A[i].jumlah < A[idx].jumlah {
					idx = i
				}
				i = i + 1
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass = pass + 1
		}
	} else if p == 5 { // sorting berdasarkan kadaluarsa paling jauh
		pass = 1

		for pass < n {
			idx = pass - 1
			i = pass

			for i < n {
				if A[i].kadaluarsa > A[idx].kadaluarsa {
					idx = i
				}
				i = i + 1
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass = pass + 1
		}
	} else if p == 6 { // sorting berdasarkan kadaluarsa terdekat
		pass = 1

		for pass < n {
			idx = pass - 1
			i = pass

			for i < n {
				if A[i].kadaluarsa < A[idx].kadaluarsa {
					idx = i
				}
				i = i + 1
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass = pass + 1
		}
	} else {
		fmt.Println("piliha tidak valid")
	}
	show(*A, n)
}

// Dumy data
func dumyData(A *tabBahan, n *int) {
	now := time.Now()

	A[0].nama = "Wortel"
	A[0].jumlah = 2
	A[0].kadaluarsa = 20
	A[0].inDate = now
	A[0].tglKadaluarsa = A[0].inDate.AddDate(0, 0, A[0].kadaluarsa)
	A[0].status = cekStatus(A[0])

	A[1].nama = "Kentang"
	A[1].jumlah = 5
	A[1].kadaluarsa = 25
	A[1].inDate = now
	A[1].tglKadaluarsa = A[1].inDate.AddDate(0, 0, A[1].kadaluarsa)
	A[1].status = cekStatus(A[1])

	A[2].nama = "Tomat"
	A[2].jumlah = 3
	A[2].kadaluarsa = 2 // akan segera kadaluarsa
	A[2].inDate = now
	A[2].tglKadaluarsa = A[2].inDate.AddDate(0, 0, A[2].kadaluarsa)
	A[2].status = cekStatus(A[2])

	A[3].nama = "Daging Ayam"
	A[3].jumlah = 1
	A[3].kadaluarsa = 1
	A[3].inDate = now.AddDate(0, 0, -2) // sudah lewat
	A[3].tglKadaluarsa = A[3].inDate.AddDate(0, 0, A[3].kadaluarsa)
	A[3].status = cekStatus(A[3])

	A[4].nama = "Susu"
	A[4].jumlah = 2
	A[4].kadaluarsa = 3 // akan kadaluarsa
	A[4].inDate = now
	A[4].tglKadaluarsa = A[4].inDate.AddDate(0, 0, A[4].kadaluarsa)
	A[4].status = cekStatus(A[4])

	A[5].nama = "Telur"
	A[5].jumlah = 12
	A[5].kadaluarsa = 10
	A[5].inDate = now
	A[5].tglKadaluarsa = A[5].inDate.AddDate(0, 0, A[5].kadaluarsa)
	A[5].status = cekStatus(A[5])

	A[6].nama = "Keju"
	A[6].jumlah = 1
	A[6].kadaluarsa = 0
	A[6].inDate = now.AddDate(0, 0, -1)
	A[6].tglKadaluarsa = A[6].inDate.AddDate(0, 0, A[6].kadaluarsa)
	A[6].status = cekStatus(A[6])

	A[7].nama = "Brokoli"
	A[7].jumlah = 4
	A[7].kadaluarsa = 4
	A[7].inDate = now
	A[7].tglKadaluarsa = A[7].inDate.AddDate(0, 0, A[7].kadaluarsa)
	A[7].status = cekStatus(A[7])

	A[8].nama = "Ikan Tuna"
	A[8].jumlah = 2
	A[8].kadaluarsa = 7
	A[8].inDate = now
	A[8].tglKadaluarsa = A[8].inDate.AddDate(0, 0, A[8].kadaluarsa)
	A[8].status = cekStatus(A[8])

	A[9].nama = "Roti"
	A[9].jumlah = 5
	A[9].kadaluarsa = 2
	A[9].inDate = now.AddDate(0, 0, -1)
	A[9].tglKadaluarsa = A[9].inDate.AddDate(0, 0, A[9].kadaluarsa)
	A[9].status = cekStatus(A[9])

	A[10].nama = "Ikan Lele"
	A[10].jumlah = 6
	A[10].kadaluarsa = 5
	A[10].inDate = now
	A[10].tglKadaluarsa = A[10].inDate.AddDate(0, 0, A[10].kadaluarsa)
	A[10].status = cekStatus(A[10])

	A[11].nama = "Ikan Mujair"
	A[11].jumlah = 3
	A[11].kadaluarsa = 2
	A[11].inDate = now
	A[11].tglKadaluarsa = A[11].inDate.AddDate(0, 0, A[11].kadaluarsa)
	A[11].status = cekStatus(A[11])

	A[12].nama = "Ikan Salmon"
	A[12].jumlah = 1
	A[12].kadaluarsa = 8
	A[12].inDate = now
	A[12].tglKadaluarsa = A[12].inDate.AddDate(0, 0, A[12].kadaluarsa)
	A[12].status = cekStatus(A[12])

	*n = 13

}
