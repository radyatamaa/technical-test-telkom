package models

type CartProduk struct {
	KodeProduk string    `gorm:"type:varchar(50);column:kode_produk"`
	Kuantitas     int            `gorm:"column:kuantitas"`
}
