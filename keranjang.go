package latihan1

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

func (c *Keranjang) TambahItem(barang Barang) {
	c.barangs = append(c.barangs, barang)
}

func (c *Keranjang) HapusItem(barangNama string) {
	for i, barang := range c.barangs {
		if barang.Nama == barangNama {
			c.barangs = append(c.barangs[:i], c.barangs[i+1:]...)
			break
		}
	}
}

func (c *Keranjang) HitungTotal() int {
	total := 0
	for _, barang := range c.barangs {
		total += barang.Harga
	}
	return total
}

func (c *Keranjang) LihatKeranjang() {
	fmt.Println("Isi Keranjang Belanja:")
	for _ , barang := range c.barangs {
		fmt.Printf("- %s: Rp%d\n", &barang.Nama, barang.Harga)
	}
}