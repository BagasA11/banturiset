package errorf

import (
	e "errors"
)

var (
	ErrRedundant         = e.New("data redundant")
	ErrDonationStillOpen = e.New("tidak dapat membuat laporan saat pendanaan masih dibuka")
	ErrHaveNotStartEvent = e.New("waktu kegiatan belum dimulai")
	ErrRedundantTahap    = e.New("data tahap redundant")
	ErrRedundantData     = e.New("data redundant")
	ErrNilTahap          = e.New("data tahapan tidak ditemukan")
	ErrInvalidRole       = e.New("role ditolak")
)
