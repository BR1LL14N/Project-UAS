package model

import (
	database "ProjekGolangMVC/Database"
	"ProjekGolangMVC/node"
	"fmt"
	"time"
)

func GetMahasiswaByEmail(email string) *node.DataMahasiswa {
	var curLL *node.DataLoginLL
	curLL = &database.DBMemberLogin

	for curLL != nil {
		if curLL.Anggota.Pelengkap.Email == email {
			return &curLL.Anggota
		}
		curLL = curLL.Next
	}
	return nil
}

func findRegistrationData(email string) *node.DataRegisterMahasiswa {
	var curLL *node.DataLL
	curLL = &database.DBMember

	// Loop melalui linked list untuk menemukan data registrasi yang sesuai dengan email
	for curLL != nil {
		if curLL.Member.Email == email {
			return &curLL.Member // Kembalikan pointer ke data registrasi yang ditemukan
		}
		curLL = curLL.Next
	}
	return nil // Kembalikan nil jika data tidak ditemukan
}

func InsertDataLoginMahasiwa(email string, tanggalLahir string, jurusan string, kondisionalMahasiswa string, tahunAjar string, transfer bool) {
	// Cari data registrasi mahasiswa berdasarkan email
	registrationData := findRegistrationData(email)
	if registrationData == nil {
		// Jika data registrasi tidak ditemukan, proses berhenti
		return
	}

	NPM := CreateNPM(jurusan, tahunAjar, transfer)

	// Membuat entri baru untuk data login mahasiswa
	newMahasiswaLogin := node.DataMahasiswa{
		Pelengkap:            *registrationData, // Menyertakan data registrasi lengkap
		NamaLengkap:          registrationData.Nama, // Menggunakan nama dari data registrasi
		TanggalLahir:         tanggalLahir,
		Jurusan:              jurusan,
		KondisionalMahasiswa: kondisionalMahasiswa,
		TahunAjar:            tahunAjar,
		NPM:                  NPM,
		CraeteAt:             time.Now().Format("2006-01-02 15:04:05"),
	}

	// Membuat linked list node baru
	newLL := node.DataLoginLL{
		Anggota: newMahasiswaLogin,
	}

	var curLL *node.DataLoginLL
	curLL = &database.DBMemberLogin

	if curLL.Next == nil {
		curLL.Next = &newLL
	} else {
		for curLL.Next != nil {
			curLL = curLL.Next
		}
		curLL.Next = &newLL
	}
}

var (
	lastSerialNumber int
	initialized      bool
)

func getNextSerialNumber() string {
	// Inisialisasi serial number jika belum diinisialisasi
	if !initialized {
		lastSerialNumber = 0
		initialized = true
	}
	
	// Tambah nomor urut terakhir
	lastSerialNumber++
	return fmt.Sprintf("%05d", lastSerialNumber)
}

func CreateNPM(jurusan string, tahunAjar string, transfer bool) string {
	// Kode jurusan
	kodeJurusan := GetKodeJurusan(jurusan)
	// Kode tahun ajar
	kodeTahun := tahunAjar
	// Kode status transfer
	kodeTransfer := GetKodeTransfer(transfer)

	serialNumber := getNextSerialNumber()

	// Gabungkan semua kode menjadi NPM
	npm := kodeJurusan + "." + kodeTahun + "." + kodeTransfer + "." + serialNumber

	return npm
}

func GetKodeJurusan(jurusan string) string {
	// Implementasi pemetaan kode jurusan
	// Anda dapat menyesuaikan pemetaan ini sesuai dengan kode jurusan yang digunakan
	kodeJurusan := map[string]string{
		"Teknik Sipil": "01",
		"Teknik Mesin": "02",
		"Teknik Elektro": "03",
		"Arsitektur": "04",
		"Teknik Perkapalan": "05",
		"Teknik Informatika": "06",
		"Teknik Industri": "07",
		"Teknik Kimia": "08",
		"Teknik Lingkungan": "09",
		"Desain Produk": "10",
		"Teknik Pertambangan": "11",
		"Sistem Informasi": "12",
	}

	return kodeJurusan[jurusan]
}

func GetKodeTransfer(transfer bool) string {
	if transfer {
		return "9" // Kode untuk mahasiswa transfer
	}
	return "1" // Kode untuk mahasiswa bukan transfer
}

func ReadAllDataLogin() []node.DataMahasiswa{
	var curLL *node.DataLoginLL
	curLL = &database.DBMemberLogin

	var memberBoard []node.DataMahasiswa

	for curLL.Next != nil {
		curLL = curLL.Next
		memberBoard = append(memberBoard, curLL.Anggota)
	}
	return memberBoard
}

func DataLogin(Email,Password string) *node.DataLL{
	curLL := &database.DBMember

	for curLL != nil {
		if curLL.Member.Email == Email && curLL.Member.Password == Password {
			return curLL
		}
		curLL = curLL.Next
	}
	return nil
}