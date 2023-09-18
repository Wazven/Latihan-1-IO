package main

import (
	"fmt"
	"os"
)

func main() {
	//daftar barang dan keranjang kosong
	ListBarang := NewBarangList()
	keranjang := NewKeranjang()

	//inisialisasi printer untuk invoice ke layar/display
	printer := &ConsolePrinter{}

	//inisialisasi display dan operasi sebagai keranjang
	display := keranjang
	operasi := keranjang

	for {
		//Menampilkan menu kasir
		fmt.Println("\nMenu Kasir:")
		fmt.Println("================================\n")
		fmt.Println("1. Tambah Barang ke Keranjang")
        fmt.Println("2. Hapus Barang dari Keranjang")
        fmt.Println("3. Lihat Keranjang")
        fmt.Println("4. Cetak Invoice ke Layar")
        fmt.Println("5. Cetak Invoice ke File")
        fmt.Println("6. Keluar")
        fmt.Println("7. Lihat Daftar Barang Tersedia")

		//menu pilihan
		var pilih int
		fmt.Println("Pilih Opearsi: ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			//menampilkan daftar barang yang tersedia
			barang, _ := ListBarang.GetAllBarang()
			DisplayListBarang(barang)

			var itemName string
			fmt.Print("Nama Barang: ")
			fmt.Scanln(&itemName)

			//mencari barang berdasarkan nama
			item := FindBarangByName(barang, itemName)
			if item == nil {
				fmt.Printf("Barang '%s' tidak ditemukan dalam daftar barang.\n", itemName)
				continue
			}

			//menambahkan barang ke keranjang
			operasi.TambahItem(*item)
			fmt.Printf("%s berhasil ditambahkan ke keranjang.\n", itemName)
		case 2:
			var itemName string
			fmt.Print("Nama Barang yang akan dihapus: ")
			fmt.Scanln(&itemName)

			//menghapus barang dari keranjang
			operasi.HapusItem(itemName)
			fmt.Printf("%s berhasil dihapus dari keranjang.\n", itemName)
		case 3:
			//display isi keranjang
			display.LihatKeranjang()
		case 4:
			//display invoice ke cli
			printer.PrintInvoice(*keranjang)
		case 5:
			//cetak invoice dalam file
			var fileName string
			fmt.Print("Masukan nama file invoice (name.txt): ")
			fmt.Scanln(&fileName)

			//memanggil pointer struct fileprinter dari invoice
			filePrinter := &FilePrinter{}
			err := filePrinter.PrintInvoice(*keranjang, fileName)
			if err != nil {
				fmt.Printf("Gagal Mencetak Invoice: %v\n", err)
			} else {
				fmt.Printf("Invoice berhasil dicetak dalam file %s\n", fileName)
			}
		case 6:
			//keluar
			fmt.Println("Terima Kasih telah berbelanja!")
			os.Exit(0)
		case 7:
			//lihat list barang
			barang, _ := ListBarang.GetAllBarang()
			DisplayListBarang(barang)
		default:
			fmt.Println("Pilihan tidak valid. Silahkan Coba Lagi")
		}
		
	}
}

func DisplayListBarang(barang []Barang){ //fungsi displaylistbarang untuk cli
	fmt.Println("Daftar Barang Tersedia: ")
	for _, barang := range barang{
		fmt.Printf("- %s: Rp%d\n", barang.Nama, barang.Harga)
	}
}

func FindBarangByName(barang []Barang, itemName string) *Barang { //fungsi untuk mencari barang berdasarkan nama dan mencocokannya dengan itemName
	for _, barang := range barang {
		if barang.Nama == itemName {
			return &barang
		}
	}
	return nil
}
