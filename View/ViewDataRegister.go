package View

import (
	controller "ProjekGolangMVC/Controller"
	"bufio"
	"fmt"
	"os"
)

func ViewInsertRegister(){
	var nama, username string
	var email, password string

	bufio.NewScanner(os.Stdin).Scan()
	fmt.Print("Silahkan Masukkan Nama Mahasiswa \t: ")
	scanner1 := bufio.NewScanner(os.Stdin);
	scanner1.Scan()
	nama = scanner1.Text()
	fmt.Print("Silahkan Masukkan Username Mahasiswa \t: ")
	scanner2 := bufio.NewScanner(os.Stdin);
	scanner2.Scan()
	username = scanner2.Text()
	fmt.Print("Silahkan Masukkan Email Mahasiswa \t: ")
	scanner3 := bufio.NewScanner(os.Stdin);
	scanner3.Scan()
	email = scanner3.Text()
	fmt.Print("Silahkan Masukkan Password Mahasiswa \t: ")
	scanner4 := bufio.NewScanner(os.Stdin);
	scanner4.Scan()
	password = scanner4.Text()

	cek := controller.ControlinsertDataRegister(email,password,nama,username)

	if cek {
		fmt.Println("Insert Berhasil")
	} else {
		fmt.Println("Insert Gagal Silahkan Isi Semua Data Yang Dibutuhkan Dengan Baik Dan Benar")
	}
}

func ViewUpdateRegister() {
	var nama, usernameKondisi string
	var email, password string

	bufio.NewScanner(os.Stdin).Scan()
	fmt.Print("Masukkan Username Mahasiswa Yang Ingin Diupdate: ")
	scanner1 := bufio.NewScanner(os.Stdin);
	scanner1.Scan()
	usernameKondisi = scanner1.Text()
	fmt.Print("Masukkan Nama Mahasiswa Yang Baru: ")
	scanner2 := bufio.NewScanner(os.Stdin);
	scanner2.Scan()
	nama = scanner2.Text()
	fmt.Print("Masukkan Email Mahasiswa Yang Baru: ")
	scanner3 := bufio.NewScanner(os.Stdin)
	scanner3.Scan()
	email = scanner3.Text()
	fmt.Print("Masukkan Password Mahasiswa Yang Baru: ")
	scanner4 := bufio.NewScanner(os.Stdin)
	scanner4.Scan()
	password = scanner4.Text()

	cek := controller.ControlupdateRegister(password, email, usernameKondisi, nama)

	if cek {
		fmt.Println("Update Berhasil")
	} else {
		fmt.Println("Update Gagal Silahkan Masukkan Data Dengan Baik Dan Benar :)")
	}
}

func ViewDeleteRegister() {
	var username string

	bufio.NewScanner(os.Stdin).Scan()
	fmt.Print("Masukkan Username Mahasiswa yang Akan Dihapus: ")
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	username = scanner1.Text()

	cek := controller.ControldeleteRegister(username)

	if cek {
		fmt.Println("Data Mahasiswa Berhasil Dihapus")
	} else {
		fmt.Println("Data Mahasiswa Gagal Dihapus atau Tidak Ditemukan")
	}
}

func ViewSearchDataRegister() {
	var username string

	bufio.NewScanner(os.Stdin).Scan()
	fmt.Print("Masukkan Username Mahasiswa yang Akan Dicari: ")
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	username = scanner1.Text()

	result := controller.ControlsearchRegister(username)

	if result != nil {
		fmt.Println("======Data Mahasiswa Ditemukan=====")
		fmt.Println("Nama :", result.Member.Nama)
		fmt.Println("Username :", result.Member.Username)
		fmt.Println("Email :", result.Member.Email)
		fmt.Println("Password :", result.Member.Password)
		fmt.Println("Creat at :", result.Member.CraeteAt)
	} else {
		fmt.Println("Data Mahasiswa Tidak Ditemukan")
	}
}

func ViewReadallDataRegister() {
    members := controller.ControlviewDataRegister()

	if members != nil {
		fmt.Println("=========================")
    	for _, member := range members {
            fmt.Println("Nama Mahasiswa : ", member.Nama)
            fmt.Println("Username Mahasiswa : ", member.Username)
            fmt.Println("Email Mahasiswa : ", member.Email)
            fmt.Println("Password Mahasiswa : ", member.Password)
            fmt.Println("Creat At : ", member.CraeteAt)
            fmt.Println("=========================")
    	}
	}else{
		fmt.Println("Data Untuk Ditampilkan Masih Kosong :)")
	}
    
}