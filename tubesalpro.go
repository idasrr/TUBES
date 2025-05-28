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

type tabBahan [NMAX]bahan

func main() {
	var data tabBahan
	var p, n int

	reader := bufio.NewReader(os.Stdin)

	dumyData(&data, &n) // isi data contoh

	for {
		menu()
		fmt.Scan(&p)
		reader.ReadString('\n') // buang newline

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
			deleteData(&data, &n)
		case 7:
			fmt.Println("Keluar dari program")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func menu() {
	fmt.Println("\n===== MENU =====")
	fmt.Println("1. INFO")
	fmt.Println("2. Input")
	fmt.Println("3. Update")
	fmt.Println("4. Cari")
	fmt.Println("5. Sorting")
	fmt.Println("6. Delete")
	fmt.Println("7. Keluar")
	fmt.Print("Masukkan Pilihan: ")
}

func input(A *tabBahan, n *int) {
	var jumlah, kadaluarsa int
	var nama string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan Nama Bahan (ketik 'none' untuk berhenti): ")
	input, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(input)

	for nama != "none" && *n < NMAX {
		fmt.Print("Masukan Jumlah Bahan dan Kadaluarsa (hari): ")
		fmt.Scan(&jumlah, &kadaluarsa)
		reader.ReadString('\n')

		A[*n].nama = nama
		A[*n].jumlah = jumlah
		A[*n].kadaluarsa = kadaluarsa
		A[*n].inDate = time.Now()
		A[*n].tglKadaluarsa = A[*n].inDate.AddDate(0, 0, kadaluarsa)
		A[*n].status = cekStatus(A[*n])

		(*n)++

		fmt.Print("Masukan Nama Bahan (ketik 'none' untuk berhenti): ")
		input, _ = reader.ReadString('\n')
		nama = strings.TrimSpace(input)
	}
}

func show(A tabBahan, n int) {
	fmt.Println("\nTanggal Sekarang:", time.Now().Format("02-01-2006"))
	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("%-4s | %-20s | %-8s | %-15s | %-20s\n", "NO", "Nama", "Jumlah", "Kadaluarsa", "Status")
	fmt.Println(strings.Repeat("-", 70))

	for i := 0; i < n; i++ {
		if A[i].nama != "" {
			fmt.Printf("%-4d | %-20s | %-8d | %-15s | %-20s\n",
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

func update(A *tabBahan, n int) {
	var p, jumlahBaru, kadaluarsaBaru int
	var namaBaru string

	if n <= 0 {
		fmt.Println("Data masih kosong")
		return
	}

	show(*A, n)

	reader := bufio.NewReader(os.Stdin)
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
	reader.ReadString('\n')

	A[p-1].nama = namaBaru
	A[p-1].jumlah = jumlahBaru
	A[p-1].kadaluarsa = kadaluarsaBaru
	A[p-1].inDate = time.Now()
	A[p-1].tglKadaluarsa = A[p-1].inDate.AddDate(0, 0, kadaluarsaBaru)
	A[p-1].status = cekStatus(A[p-1])
}

func deleteData(A *tabBahan, n *int) {
	if *n <= 0 {
		fmt.Println("Data masih kosong")
		return
	}
	show(*A, *n)

	var p int
	fmt.Print("Hapus Data ke? ")
	fmt.Scan(&p)

	if p < 1 || p > *n {
		fmt.Println("Data tidak ditemukan")
		return
	}

	for i := p - 1; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n--
	fmt.Println("Data berhasil di hapus")
}

func cekStatus(B bahan) string {
	sisa := int(B.tglKadaluarsa.Sub(time.Now()).Hours() / 24)
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
	keyword = strings.ToLower(keyword)
	for i := 0; i < n; i++ {
		if strings.Contains(strings.ToLower(A[i].nama), keyword) {
			return true
		}
	}
	return false
}

// Binary search pertama untuk status (harus data sudah urut status A-Z)
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

func search(A *tabBahan, n int) {
	reader := bufio.NewReader(os.Stdin)

	if n == 0 {
		fmt.Println("Data masih kosong, tidak bisa mencari.")
		return
	}

	fmt.Println("Pilih metode pencarian:")
	fmt.Println("1. Cari berdasarkan nama (sequential search)")
	fmt.Println("3. Cari berdasarkan status (binary search, harus urutkan dulu berdasarkan status)")
	fmt.Print("Masukkan pilihan: ")

	var pilihan int
	fmt.Scan(&pilihan)
	reader.ReadString('\n')

	switch pilihan {
	case 1:
		fmt.Print("Masukkan nama bahan yang dicari: ")
		input, _ := reader.ReadString('\n')
		keyword := strings.TrimSpace(input)

		if sequentialSearch(*A, n, keyword) {
			fmt.Println("Bahan ditemukan.")
			fmt.Println("Detail bahan: ")

			fmt.Println(strings.Repeat("=", 70))
			fmt.Printf("%-4s | %-20s | %-8s | %-15s | %-20s\n", "NO", "Nama", "Jumlah", "Kadaluarsa", "Status")
			fmt.Println(strings.Repeat("-", 70))

			hitung := 0
			for i := 0; i < n; i++ {
				if strings.Contains(strings.ToLower((*A)[i].nama), strings.ToLower(keyword)) {
					fmt.Printf(
						"%-4d | %-20s | %-8d | %-15s | %-20s\n",
						i+1,
						A[i].nama,
						A[i].jumlah,
						A[i].tglKadaluarsa.Format("02-01-2006"),
						A[i].status,
					)
					hitung++
				}
			}

			fmt.Println(strings.Repeat("=", 70))
			fmt.Printf("Total ditemukan: %d item.\n", hitung)
			fmt.Println(strings.Repeat("=", 70))
		} else {
			fmt.Println("Bahan tidak ditemukan.")
		}

	case 2:
		fmt.Println("Pastikan data sudah terurut berdasarkan status terlebih dahulu (menu Sorting pilih 7)")
		fmt.Print("Masukkan status yang dicari (Aman, Akan Kadaluarsa, Segera Kadaluarsa, Sudah Kadaluarsa): ")
		input, _ := reader.ReadString('\n')
		target := strings.TrimSpace(input)

		first := BinarySearchFirstStatus(*A, n, target)
		last := BinarySearchLastStatus(*A, n, target)

		if first == -1 || last == -1 {
			fmt.Println("Tidak ditemukan bahan dengan status tersebut.")
			return
		}

		fmt.Printf("Bahan dengan status \"%s\" ditemukan:\n", target)
		fmt.Printf("%-4s | %-20s | %-8s | %-15s | %-20s\n", "NO", "Nama", "Jumlah", "Kadaluarsa", "Status")
		fmt.Println(strings.Repeat("-", 70))

		hitung := 0
		for i := first; i <= last; i++ {
			fmt.Printf(
				"%-4d | %-20s | %-8d | %-15s | %-20s\n",
				i+1,
				A[i].nama,
				A[i].jumlah,
				A[i].tglKadaluarsa.Format("02-01-2006"),
				A[i].status,
			)
			hitung++
		}

		fmt.Println(strings.Repeat("=", 70))
		fmt.Printf("Total ditemukan: %d item.\n", hitung)
		fmt.Println(strings.Repeat("=", 70))

	default:
		fmt.Println("Pilihan tidak valid")
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
	fmt.Println("7. Berdasarkan Status (A-Z)")
	fmt.Print("Masukkan Pilihan: ")
}

func sorting(A *tabBahan, n int) {
	var i, idx, pass int
	var temp bahan

	menuSorting()
	var p int
	fmt.Scan(&p)

	switch p {
	case 1:
		pass = 1
		for pass < n {
			idx = pass - 1
			i = pass
			for i < n {
				if strings.ToLower(A[i].nama) < strings.ToLower(A[idx].nama) {
					idx = i
				}
				i++
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass++
		}
	case 2:
		pass = 1
		for pass < n {
			idx = pass - 1
			i = pass
			for i < n {
				if strings.ToLower(A[i].nama) > strings.ToLower(A[idx].nama) {
					idx = i
				}
				i++
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass++
		}
	case 3:
		pass = 1
		for pass < n {
			idx = pass - 1
			i = pass
			for i < n {
				if A[i].jumlah > A[idx].jumlah {
					idx = i
				}
				i++
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass++
		}
	case 4:
		pass = 1
		for pass < n {
			idx = pass - 1
			i = pass
			for i < n {
				if A[i].jumlah < A[idx].jumlah {
					idx = i
				}
				i++
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass++
		}
	case 5:
		pass = 1
		for pass < n {
			idx = pass - 1
			i = pass
			for i < n {
				if A[i].kadaluarsa > A[idx].kadaluarsa {
					idx = i
				}
				i++
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass++
		}
	case 6:
		pass = 1
		for pass < n {
			idx = pass - 1
			i = pass
			for i < n {
				if A[i].kadaluarsa < A[idx].kadaluarsa {
					idx = i
				}
				i++
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass++
		}
	case 7:
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
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

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
