package main

type BarangRepository interface { //BarangRepository mendefinisikan metode GetAllBarang() ([]Barang, error).
	GetAllBarang() ([]Barang, error)
}

type ListBarang struct {
	barang []Barang
}

func NewBarangList() *ListBarang { // NewBarangList membuat dan menginisialisasi instansi baru dari ListBarang.

	return &ListBarang{
		barang: []Barang{
			{ID: 1, Nama: "Buku", Harga: 12000},
			{ID: 2, Nama: "Pulpen", Harga: 3500},
			{ID: 3, Nama: "Pensil", Harga: 2000},
			{ID: 4, Nama: "Tipe-x", Harga: 4000},
			{ID: 5, Nama: "Pulpen_Warna", Harga: 7000},
		},
	}
}

func (lb *ListBarang) GetAllBarang() ([]Barang, error) {
	return lb.barang, nil
}

// Metode GetAllBarang() dalam ListBarang mengembalikan daftar semua barang yang tersedia.
// Dalam implementasi ini, daftar barang diambil dari variabel barang yang telah diinisialisasi.
