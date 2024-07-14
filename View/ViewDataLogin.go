package View

import (
	controller "ProjekGolangMVC/Controller"
	"bufio"
	"fmt"
	"os"
)

func ViewDataLogin() {
	var email, pass string

	bufio.NewScanner(os.Stdin).Scan()
	fmt.Print("Silahkan Masukkan Email Anda : ")
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	email = scanner1.Text()
	fmt.Print("Silahkan Masukkan Password Anda : ")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	pass = scanner2.Text()

	result := controller.ControlDataLogin(email, pass)

	if result != nil {
		existingData := controller.ControlGetMahasiswaByEmail(email)

		if existingData != nil{
			
			fmt.Println("=====Login Berhasil=====")
			fmt.Println("Data Mahasiswa Anda Dibawah Ini : ")
			fmt.Println("Nama Lengkap Mahasiswa : ", existingData.NamaLengkap)
			fmt.Println("Tanggal Lahir Mahasiswa : ", existingData.TanggalLahir)
			fmt.Println("Jurusan Yang Diampuh Mahasiswa : ", existingData.Jurusan)
			fmt.Println("NPM Mahasiswa : ", existingData.NPM)
			fmt.Println("Create At : ", existingData.CraeteAt)

		}else{
			bufio.NewScanner(os.Stdin).Scan()
			fmt.Println("=====Login Berhasil=====")
			fmt.Println("Silahkan Lengkapi Data Diri Anda Dibawah Ini :)")
			fmt.Print("Silahkan Masukkan Tanggal Lahir Anda \t: ")
			scanner3 := bufio.NewScanner(os.Stdin)
			scanner3.Scan()
			tglLahir := scanner3.Text()
			fmt.Print("Silahkan Masukkan Jurusan Yang Akan Anda Ambil \t: ")
			scanner4 := bufio.NewScanner(os.Stdin)
			scanner4.Scan()
			jurusan := scanner4.Text()
			fmt.Print("Apakah Anda Mahasiswa Transfer (Y/T) \t: ")
			scanner5 := bufio.NewScanner(os.Stdin)
			scanner5.Scan()
			kondisioanalMHS := scanner5.Text()

			// Konversi input kondisi mahasiswa menjadi tipe data boolean
			transfer := false
			if kondisioanalMHS == "Y" || kondisioanalMHS == "y" {
				transfer = true
			}

			fmt.Print("Silahkan Masukkan Tahun Ajar Anda \t: ")
			scanner6 := bufio.NewScanner(os.Stdin)
			scanner6.Scan()
			tahunAjar := scanner6.Text()

			controller.ControlinsertDataLogin(email, tglLahir, jurusan, kondisioanalMHS, tahunAjar, transfer)
		}
		
	} else {
		fmt.Println("Login Gagal Silahkan Periksa Kembali Apakah Email Dan Password Anda Sudah Benar")
	}
}

func ViewReadAllDataLogin(){
	members := controller.ControlReadAllDataLogin()

	if members != nil{
		fmt.Println("=============================")
		for _, member := range members {
			fmt.Println("Nama Lengkap Mahasiswa : ", member.NamaLengkap)
			fmt.Println("Tanggal Lahir Mahasiswa : ", member.TanggalLahir)
			fmt.Println("Jurusan Yang Diampuh Mahasiswa : ",member.Jurusan)
			fmt.Println("NPM Mahasiswa : ",member.NPM)
			fmt.Println("Create At : ",member.CraeteAt)
			fmt.Println("=============================")
		}
	}else{
		fmt.Println("Data Untuk Ditampilkan Masih Kosong :)")
	}
}