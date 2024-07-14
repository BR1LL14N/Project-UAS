package node

type DataRegisterMahasiswa struct {
	Email    string
	Password string
	Nama     string
	Username string
	CraeteAt string
}

type DataLL struct {
	Member DataRegisterMahasiswa
	Next   *DataLL
}