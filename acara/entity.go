package acara

import "time"

type Acara struct {
	ID            int
	Nama_acara    string
	Harga         string
	Tanggal_acara time.Time
	Deskripsi     string
	Img_url       string
}
