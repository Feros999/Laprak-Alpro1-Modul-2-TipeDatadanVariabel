package main

import "fmt"

func main() {
	var (
		nama, nim, kelas, prodi string
	)
	fmt.Print("Masukan Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukan Prodi: ")
	fmt.Scanln(&prodi)
	fmt.Print("Masukan Kelas: ")
	fmt.Scanln(&kelas)
	fmt.Print("Masukan NIM: ")
	fmt.Scanln(&nim)

	fmt.Println("\n=== Resume Mahasiswa ===")
	fmt.Println("Nama  : " + nama)
	fmt.Println("NIM   : " + nim)
	fmt.Println("Kelas : " + kelas)
	fmt.Println("Prodi : " + prodi)
	fmt.Println("========================")
	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi %s dari kelas %s dengan NIM %s.\n", nama, prodi, kelas, nim)
	fmt.Println("Selamat belajar di Tel-U, semoga sukses dalam studi Anda!")
}
