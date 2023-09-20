package main

import (
	"fmt"
	"os"
)

type Invoice interface { //interface yang digunakan untuk mencetak invoice
	PrintInvoice(keranjang Keranjang)
}
// ConsolePrinter adalah struct yang mengimplementasikan interface Invoice
// untuk mencetak invoice ke konsol.
type ConsolePrinter struct{}
// PrintInvoice mencetak invoice ke konsol.
func (cp *ConsolePrinter) PrintInvoice(keranjang Keranjang) {
	fmt.Println("Invoice Belanja: ")
	keranjang.LihatKeranjang()
	fmt.Printf("Total Harga: Rp%d\n", keranjang.HitungTotal())
	fmt.Printf("Total Barang: %d\n", keranjang.JumlahTotalBarang())
}

// FilePrinter adalah struct yang mengimplementasikan interface Invoice
// untuk mencetak invoice ke dalam file.
type FilePrinter struct{}
// PrintInvoice mencetak invoice ke dalam file dengan nama yang ditentukan.
func (fp *FilePrinter) PrintInvoice(keranjang Keranjang, fileName string) error {
	//Implementasi pencetakan file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close() //pastikan file ditutup setelah selesai

	//Menuliskan rincian invoice ke file
	konten := fmt.Sprintf("Invoice Belanja:\n")
	for _, barang := range keranjang.barang {
		konten += fmt.Sprintf("- %s: Rp%d\n", barang.Nama, barang.Harga)
	}
	konten += fmt.Sprintf("Total Harga: Rp%d\n", keranjang.HitungTotal())
	konten += fmt.Sprintf("Total Barang: %d\n", keranjang.JumlahTotalBarang())
	_, err = file.WriteString(konten)
	if err != nil {
		return err
	}
	return nil
}