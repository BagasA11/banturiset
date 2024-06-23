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
	if subj == "verif_project" {
		body = verif_proyek(args[0], args[1], args[2])
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

func verif_proyek(email string, projectID string, title string) string {
	return fmt.Sprintf("kepada yth. %s, kami menginformasikan bahwa proyek penelitian anda dengan id %s dan judul %s telah diverifikasi",
		email, projectID, title)
}
