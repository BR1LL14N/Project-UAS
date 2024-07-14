package controller

import (
	model "ProjekGolangMVC/Model"
	"ProjekGolangMVC/node"
)

func ControlGetMahasiswaByEmail(email string) *node.DataMahasiswa {
	return model.GetMahasiswaByEmail(email)
}

func ControlinsertDataLogin(email string, tanggalLahir string, jurusan string, kondisionalMahasiswa string, tahunAjar string, transfer bool) bool {
	if email != "" && tanggalLahir != "" && jurusan != "" && kondisionalMahasiswa != "" && tahunAjar != "" {
		model.InsertDataLoginMahasiwa(email, tanggalLahir, jurusan, kondisionalMahasiswa, tahunAjar, transfer)
		return true
	}
	return false
}

// Nilai khusus yang menunjukkan kesalahan
const (
	ErrJurusanTidakTersedia = "Jurusan tidak tersedia"
)

func ControlCreateNPM(jurusan string, tahunAjar string, transfer string) (string, string) {
	// Mendapatkan kode jurusan menggunakan fungsi dari package model
	kodeJurusan := model.GetKodeJurusan(jurusan)

	// Periksa apakah kode jurusan kosong
	if kodeJurusan == "" {
		// Jurusan yang dimasukkan tidak tersedia, kembalikan pesan kesalahan
		return "", ErrJurusanTidakTersedia
	}

	// Lakukan sesuatu dengan kode jurusan yang didapat
	// Misalnya, gabungkan dengan tahun ajar dan keterangan transfer untuk membentuk NPM

	return kodeJurusan, ""
}

func ControlDataLogin(Email,Pass string) *node.DataLL{
	hasil := model.DataLogin(Email, Pass)

	if hasil == nil {
		return nil
	}
	return hasil
}

func ControlReadAllDataLogin() []node.DataMahasiswa{
	Member := model.ReadAllDataLogin()

	if Member == nil{
		return nil
	}
	return Member
}