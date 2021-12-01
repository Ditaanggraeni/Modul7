package main

import (
	"fmt"
)

type listCust struct {
	//id            int
	Nama          string
	jmlJam, Biaya float64
}

var namaCust string
var jam, biaya float64

var ID int = 0

var cust = []listCust{}

//var list [][]string

// func validate(judul string) (bool, error) {
// 	if strings.TrimSpace(judul) == "" {
// 		return true, errors.New("Judul cannot be empty")
// 	}
// 	return true, nil
// }

func addCust() {
	ID++
	fmt.Print("\nMasukkan Nama : ")
	fmt.Scan(&namaCust)
	fmt.Print("Masukkan Jumlah Jam : ")
	fmt.Scan(&jam)

	biaya = jam * 60 * 1000

	cust = append(cust, listCust{
		Nama:   namaCust,
		jmlJam: jam,
		Biaya:  biaya,
	})
	// else {
	// 	fmt.Println(err.Error())
	// }
}

func deleteCust() {

	if len(cust) == 0 {
		fmt.Println("Tidak ada data yang bisa dihapus")
	}
	var before = cust[:ID-1]
	var after = cust[ID:]
	cust = append(before, after...)

}

func avgJam() float64 {
	var jumlah float64 = 0
	var count float64 = 0
	for i := range cust {
		lists := cust[i]
		jumlah += lists.jmlJam
		count++
	}

	avg := jumlah / count
	return avg
}

// func viewTop3() {
// 	var top int = 0
// 	var i float64 = 5.0

// 	for i >= 0 {
// 		for j := range cust {
// 			lists := cust[j]
// 			if lists.jmlVote <= i {
// 				fmt.Println(top+1, "\t", lists.Judul, "\t\t\t", lists.penyanyi, "\t\t\t", lists.jmlVote)
// 				top++
// 			}
// 			if top == 3 {
// 				break
// 			}
// 		}
// 		if top <= 3 {
// 			break
// 		}
// 	}
// }

// func searchData() {
// 	for j := range list {
// 		lists := list[j]
// 		var isPrefix = strings.HasPrefix(lists.penyanyi, "a")
// 		if isPrefix == true {
// 			fmt.Println(j+1, "\t", lists.Judul, "\t\t\t", lists.penyanyi, "\t\t\t", lists.jmlVote)
// 		} else {
// 			fmt.Println("\n\t\t\tDATA TIDAK DITEMUKAN!")
// 		}
// 	}
// }

func main() {
	var pilih int
	var x = true

	for x {
		fmt.Println("\n\t\t\t====Daftar Pilihan==== ")
		fmt.Println("1. Input Data Customer")
		fmt.Println("2. Hapus data customer berdasarkan id ")
		fmt.Println("3. Tampilkan seluruh data customer beserta jumlah data yang tersimpan dalam list")
		fmt.Println("4. Menampilkan rata-rata jumlah penggunaan jam dari customer ")
		fmt.Println("5. Menampilkan 3 buah data yang memiliki jam penggunaan paling sedikit.")
		fmt.Println("6. Menampilkan seluruh data customer yang menyewa komputer kurang dari rata – rata")
		fmt.Println("0. Keluar")
		fmt.Println("=================================================================== ")
		fmt.Print("Masukkan Pilihan : ")
		fmt.Scan(&pilih)
		fmt.Println("==================================================================== ")

		switch pilih {
		case 1:
			addCust()

		case 2:
			fmt.Print("\nMasukkan ID yang akan dihapus: ")
			fmt.Scan(&ID)

			deleteCust()
		case 3:

			fmt.Println("ID \tNama \t\tJumlah Jam \t\tBiaya")

			for i := range cust {
				lists := cust[i]
				fmt.Println(i+1, "\t", lists.Nama, "\t\t\t", lists.jmlJam, "\t\t\t", lists.Biaya)
			}
			fmt.Println("==================================================================== ")
			fmt.Println("Jumlah Data Dalam List : ", len(cust), "Item")

		case 4:
			fmt.Print("Jumlah Keseluruhan Jumlah Vote : ", avgJam())

		case 5:
			var top int = 0
			fmt.Println("      3  DATA YANG MEMILIKI JUMLAH PENGGUNAAN JAM PALING SEDIKIT     ")
			fmt.Println("==================================================================== ")
			for i := range cust {
				lists := cust[i]
				if lists.jmlJam < avgJam() {
					fmt.Println(top+1, "\t", lists.Nama, "\t\t\t", lists.jmlJam, "\t\t\t", lists.Biaya)
					top++
				}
				if top == 3 {
					break
				}
			}
			if top <= 3 {
				break
			}
		case 6:
			fmt.Println("      DATA CUSTOMER YANG MENYEWA KOMPUTER KURANG DARI RATA-RATA      ")
			fmt.Println("==================================================================== ")
			for i := range cust {
				lists := cust[i]
				if lists.jmlJam < avgJam() {
					fmt.Println(i+1, "\t", lists.Nama, "\t\t\t", lists.jmlJam, "\t\t\t", lists.Biaya)
				}
			}
		case 0:
			x = false
		}
	}
}
