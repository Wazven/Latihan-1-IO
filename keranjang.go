package main

import "fmt"

type Ikeranjang interface { //interface untuk mendefinisikan operasi-operasi yang dapat dilakukan pada keranjang belanja
	TambahItem(barang Barang) // Tambah barang ke keranjang
	HapusItem(itemName string) //Hapus barang dari keranjang
	HitungTotal() int //Hitung total harga
	LihatKeranjang() //Lihat isi Keranjang
}

type Keranjang struct { //Mewakili keranjang dan mengimplementasikan interface Ikeranjang
	barangs []Barang
}

func NewKeranjang() *Keranjang { //construct yang digunakan untuk membuat dan menginisialisasi intance baru dari Keranjang.
	return &Keranjang{}
}

func (k *Keranjang) TambahItem(barang Barang) { //Tambah barang ke dalam keranjang
	k.barangs = append(k.barangs, barang)
}

func (k *Keranjang) HapusItem(barangNama string) { //Hapus barang dari keranjang
	//perulangan untuk mencari barang yang akan dihapus
	for i, barang := range k.barangs {
		if barang.Nama == barangNama {
			//append untuk menghapus barang dari slice
			k.barangs = append(k.barangs[:i], k.barangs[i+1:]...)
			break
		}
	}
}

func (k *Keranjang) HitungTotal() int { //Hitung total harga belanja
	total := 0
	//perulangan untuk mengakses dan menjumlahkan harga setiap barang dlm keranjang
	for _, barang := range k.barangs {
		total += barang.Harga
	}
	return total
}

func (k *Keranjang) LihatKeranjang() { //menampilkan isi keranjang
	fmt.Println("Isi Keranjang Belanja:")
	//perulangan untuk mencetak setiap barang dalam keranjang
	for _ , barang := range k.barangs {
		//pointer k untuk mengakses atribut nama dan harga dari barang
		fmt.Printf("- %s: Rp%d\n",barang.Nama, barang.Harga)
	}
}