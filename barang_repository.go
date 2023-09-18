package latihan1

type BarangRepository interface {
	GetAllBarang() ([]Barang, error)
}

type ListBarang struct {
	barang []Barang
}

func NewBarangList() *ListBarang {
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
