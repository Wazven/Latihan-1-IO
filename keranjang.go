package main

import "fmt"

type Ikeranjang interface { //interface untuk mendefinisikan operasi-operasi yang dapat dilakukan pada keranjang belanja
	TambahItem(barang Barang) // Tambah barang ke keranjang
	HapusItem(itemID int) //Hapus barang dari keranjang
	HitungTotal() int //Hitung total harga
	LihatKeranjang() //Lihat isi Keranjang
	JumlahTotalBarang() int
	TambahBarangJumlah(barang Barang, jumlah int)
}

type Keranjang struct { //Mewakili keranjang dan mengimplementasikan interface Ikeranjang
	barang []Barang
	totalJumlahBarang int
}

func NewKeranjang() *Keranjang { //construct yang digunakan untuk membuat dan menginisialisasi intance baru dari Keranjang.
	return &Keranjang{}
}

func (k *Keranjang) TambahItem(barang Barang) { //Tambah barang ke dalam keranjang
	k.barang = append(k.barang, barang)
	k.totalJumlahBarang++
}

func (k *Keranjang) HapusItem(itemID int) { //Hapus barang dari keranjang
	//perulangan untuk mencari barang yang akan dihapus
	for i, barangs := range k.barang {
		if barangs.ID == itemID {
			//append untuk menghapus barang dari slice
			k.barang = append(k.barang[:i], k.barang[i+1:]...)
			k.totalJumlahBarang--
			break
		}
	}
}

func (k *Keranjang) HitungTotal() int { //Hitung total harga belanja
	total := 0
	//perulangan untuk mengakses dan menjumlahkan harga setiap barang dlm keranjang
	for _, barangs := range k.barang {
		total += barangs.Harga
	}
	return total
}

func (k *Keranjang) LihatKeranjang() { //menampilkan isi keranjang
	fmt.Println("Isi Keranjang Belanja:")
	//perulangan untuk mencetak setiap barang dalam keranjang
	for _ , barang := range k.barang {
		//pointer k untuk mengakses atribut nama dan harga dari barang
		fmt.Printf("- %s: Rp%d\n",barang.Nama, barang.Harga)
	}
}

func (k *Keranjang) JumlahTotalBarang() int{
	return k.totalJumlahBarang
}

func (k *Keranjang) TambahBarangJumlah(barang Barang, jumlah int){
	for i := 0; i<jumlah; i++{
		k.TambahItem(barang)
	}
}