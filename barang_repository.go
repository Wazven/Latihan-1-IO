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
			{"Buku", 12000},
			{"Pulpen", 3500},
			{"Pensil", 2000},
			{"Tipe-x", 4000},
			{"Pulpen Warna", 7000},
		},
	}
}

func (lb *ListBarang) GetAllBarang() ([]Barang, error) {
	return lb.barang, nil
}

// Metode GetAllBarang() dalam ListBarang mengembalikan daftar semua barang yang tersedia.
// Dalam implementasi ini, daftar barang diambil dari variabel barang yang telah diinisialisasi.
