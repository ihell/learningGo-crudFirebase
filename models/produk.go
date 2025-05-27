package models

type Produk struct {
	Nama  string  `json:"nama" firestore:"nama"`
	Harga float64 `json:"harga" firestore:"harga"`
	Stok  int     `json:"stok" firestore:"stok"`
}
