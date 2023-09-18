package main

import (
	"fmt"
	"os"
)

type Invoice interface {
	PrintInvoice(keranjang Keranjang)
}

type ConsolePrinter struct{}

func (cp *ConsolePrinter) PrintInvoice(keranjang Keranjang) {
	fmt.Println("Invoice Belanja: ")
	keranjang.LihatKeranjang()
	fmt.Printf("Total Harga: Rp%d\n", keranjang.HitungTotal())
}

type FilePrinter struct{}
func (fp *FilePrinter) PrintInvoice(keranjang Keranjang, fileName string) error {
	//Implementasi pencetakan file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close() //pastikan file ditutup setelah selesai

	//Menuliskan rincian invoice ke file
	konten := fmt.Sprintf("Invoice Belanja:\n")
	for _, barang := range keranjang.barangs {
		konten += fmt.Sprintf("- %s: Rp%d\n", barang.Nama, barang.Harga)
	}
	konten += fmt.Sprintf("Total Harga: Rp%d\n", keranjang.HitungTotal())
	_, err = file.WriteString(konten)
	if err != nil {
		return err
	}
	return nil
}