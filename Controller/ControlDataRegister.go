package controller

import (
	model "ProjekGolangMVC/Model"
	"ProjekGolangMVC/node"
)

func ControlinsertDataRegister(Email string, Pass string, Nama string, Username string) bool {
	if (Email != "" && Pass != "" && Nama != "" && Username != ""){
		model.InsertMahasiswaRegister(Email,Pass,Nama,Username)
		return true
	}
	return false
}

func ControlviewDataRegister() []node.DataRegisterMahasiswa{
	Member := model.ReadAllDataRegister()

	if Member == nil {
		return nil
	}
	return Member
}

func ControlsearchRegister(Username string) *node.DataLL {
	hasil := model.SearchDataRegister(Username)

	if hasil == nil{
		return nil
	}
	return hasil
}

func ControldeleteRegister(Search string) bool {
	// Lakukan pencarian data berdasarkan username
	prevNode, currentNode := model.SearchDataBelakang(Search)

	// Jika currentNode adalah nil, artinya data yang akan dihapus tidak ditemukan
	if currentNode == nil {
		return false
	}

	// Hapus data berdasarkan node yang ditemukan
	return model.DeleteDataRegister(prevNode, currentNode)
}

func ControlupdateRegister( Pass string, Email string,Usernamekondisi string, Nama string) bool {
	Search := model.SearchDataRegister(Usernamekondisi)
	if Search  != nil {
		model.UpdateDataRegister(Email,Pass,Nama,Usernamekondisi, Search)
		return true
	}else{
		return false
	}
}