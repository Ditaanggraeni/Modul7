package main

import (
	"fmt"
	"math"
)

var jumlah int = 0

func deretPangkat6(i int, n int) int {

	if i <= 1 {
		fmt.Printf("%d + ", n)
		jumlah = jumlah + n
		return n
	}
	y := deretPangkat6(i-1, n)
	pangkat := int(math.Pow(float64(6), float64(i-1)))
	jumlah = n * pangkat
	fmt.Printf("%d + ", jumlah)

	return y + jumlah
}

func deretPangkat10(i int, n int) int {

	if i <= 1 {
		fmt.Printf("%d + ", n)
		jumlah = jumlah + n
		return n
	}
	y := deretPangkat10(i-1, n)
	pangkat := int(math.Pow(float64(10), float64(i-1)))
	jumlah = n * pangkat
	fmt.Printf("%d + ", jumlah)

	return y + jumlah
}

func deretPangkat7(i int, n int) int {

	if i <= 1 {
		fmt.Printf("%d + ", n)
		jumlah = jumlah + n
		return n
	}
	y := deretPangkat7(i-1, n)
	pangkat := int(math.Pow(float64(7), float64(i-1)))
	jumlah = n * pangkat
	fmt.Printf("%d + ", jumlah)

	return y + jumlah
}

func deretD(i int, n int) int {
	if i <= 1 {
		fmt.Printf("%d + ", n)
		jumlah = jumlah + n
		return n
	}
	y := deretD(i-1, n)
	pangkat := int(math.Pow(float64(10), float64(i-1)))
	jumlah = i * pangkat
	fmt.Printf("%d + ", jumlah)

	return y + jumlah
}

func deretE(i int, n int) int {

	if i <= 1 {
		fmt.Printf("%d + ", n)
		jumlah = jumlah + n
		return n
	}
	y := deretE(i-1, n)
	pangkat := int(math.Pow(float64(10), float64(i-1)))
	jumlah = n * pangkat
	fmt.Printf("%d + ", jumlah)

	return y + jumlah
}

func main() {
	var jumlah, pilihan int
	var x = true

	for x {
		fmt.Println("\n================== PILIHAN ===================")
		fmt.Println("1. Deret Pangkat 6")
		fmt.Println("2. Deret Pangkat 10")
		fmt.Println("3. Deret Pangkat 7")
		fmt.Println("4. SUM Deret D")
		fmt.Println("5. SUM Deret E")
		fmt.Println("0. Keluar")
		fmt.Println("==============================================")
		fmt.Print("Masukkan Pilihan : ")
		fmt.Scan(&pilihan)
		fmt.Println("==============================================")

		switch pilihan {
		case 1:
			fmt.Print("Masukan jumlah deret : ")
			fmt.Scan(&jumlah)
			fmt.Print("SUM = ")
			fmt.Println(" =", deretPangkat6(jumlah, 1))
		case 2:
			fmt.Print("Masukan jumlah deret : ")
			fmt.Scan(&jumlah)
			fmt.Print("SUM = ")
			fmt.Println("= ", deretPangkat10(jumlah, 5))
		case 3:
			fmt.Print("Masukan jumlah deret : ")
			fmt.Scan(&jumlah)
			fmt.Print("SUM = ")
			fmt.Println("= ", deretPangkat7(jumlah, 1))
		case 4:
			fmt.Print("Masukan jumlah deret : ")
			fmt.Scan(&jumlah)
			fmt.Print("SUM = ")
			fmt.Println("= ", deretD(jumlah, 1))
		case 5:
			fmt.Print("Masukan jumlah deret : ")
			fmt.Scan(&jumlah)
			fmt.Print("SUM = ")
			fmt.Println("= ", deretPangkat10(jumlah, 5))
		case 0:
			x = false
		}
	}
}
