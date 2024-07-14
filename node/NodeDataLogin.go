package node

type DataMahasiswa struct {
	Pelengkap            DataRegisterMahasiswa
	NamaLengkap          string
	TanggalLahir         string
	Jurusan              string
	KondisionalMahasiswa string
	TahunAjar            string
	NPM                  string
	CraeteAt             string
}

type DataLoginLL struct {
	Anggota DataMahasiswa
	Next    *DataLoginLL
}