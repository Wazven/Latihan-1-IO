package main

import "fmt"

type Ikeranjang interface {
	TambahItem(barang Barang)
	HapusItem(itemName string)
	HitungTotal() int
	LihatKeranjang()
}

type Keranjang struct {
	barangs []Barang
}

func NewKeranjang() *Keranjang {
	return &Keranjang{}
}

func (k *Keranjang) TambahItem(barang Barang) {
	k.barangs = append(k.barangs, barang)
}

func (k *Keranjang) HapusItem(barangNama string) {
	for i, barang := range k.barangs {
		if barang.Nama == barangNama {
			k.barangs = append(k.barangs[:i], k.barangs[i+1:]...)
			break
		}
	}
}

func (k *Keranjang) HitungTotal() int {
	total := 0
	for _, barang := range k.barangs {
		total += barang.Harga
	}
	return total
}

func (k *Keranjang) LihatKeranjang() {
	fmt.Println("Isi Keranjang Belanja:")
	for _ , barang := range k.barangs {
		fmt.Printf("- %s: Rp%d\n",barang.Nama, barang.Harga)
	}
}