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

			var itemID int
			fmt.Print("ID Barang: ")
			fmt.Scanln(&itemID)

			//mencari barang berdasarkan nama
			item := FindBarangByID(barang, itemID)
			if item == nil {
				fmt.Printf("Barang '%s' tidak ditemukan dalam daftar barang.\n", itemID)
				continue
			}
			var jumlah int
			fmt.Print("Jumlah Barang: ")
			fmt.Scanln(&jumlah)


			//menambahkan barang ke keranjang
			operasi.TambahBarangJumlah(*item, jumlah)
			fmt.Printf("%s berhasil ditambahkan ke keranjang %d.\n", item.Nama, jumlah)
		case 2:
			keranjang.LihatKeranjang()

			var itemID int
			fmt.Print("ID Barang yang akan dihapus: ")
			fmt.Scanln(&itemID)

			//menghapus barang dari keranjang
			operasi.HapusItem(itemID)
			fmt.Printf("ID Barang %d berhasil dihapus dari keranjang.\n", itemID)
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

func FindBarangByID(barang []Barang, id int) *Barang { //fungsi untuk mencari barang berdasarkan nama dan mencocokannya dengan itemName
	for _, item := range barang {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func DisplayListBarang(barang []Barang){
	fmt.Println("Daftar Barang: ")
	for _, barang := range barang{
		fmt.Printf("- ID: %d, Nama: %s, Harga: Rp%d \n", barang.ID,barang.Nama,barang.Harga)
	}
}