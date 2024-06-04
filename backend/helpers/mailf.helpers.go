package helpers

import (
	"fmt"
	"strings"
)

func BodyF(konten string, args ...string) string {

	body := ""
	subj := strings.ToLower(konten)
	if subj == "verifikasi" {
		body = notif_verifikasi()
	}
	return fmt.Sprintf("Subject:%s\n%s", subj, body)
}

// Verifikasi: Kami memberitahukan bahwa akun pendaftaran anda telah diverifikasi. Mohon segera lengkapi data diri dan informasi pembayaran
// Transaksi: Mohon segera lakukan pembayaran pada link berikut: //https:example.com sebelum dd-mm-yyyyy
// Notifikasi pembayaran: Transaksi pembayaran anda dengan id transaksi: %s telah terbayar. Terimakasih
// Notifikasi pemblokiran: Maaf, akun anda kami blokir karena ditemukan adanya pelanggaran aktifitas akun anda di platform kami

func notif_verifikasi() string {
	return "Kami memberitahukan bahwa akun pendaftaran anda telah diverifikasi. Mohon segera lengkapi data diri dan informasi pembayaran"
}
