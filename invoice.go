package latihan1

import "fmt"

type Invoice interface {
	PrintInvoice(keranjang Keranjang)
}

type ConsolePrinter struct{}

func (cp *ConsolePrinter) PrintInvoice(keranjang Keranjang) {
	fmt.Println("Invoice Belanja: ")
	keranjang.LihatKeranjang()
	fmt.Printf("Total Harga: Rp%d\n", keranjang.HitungTotal())
}

type filePrinter struct{}
func (fp *filePrinter) PrintInvoice(keranjang Keranjang, fileName string) error {
	//Implementasi pencetakan file
	return nil
}