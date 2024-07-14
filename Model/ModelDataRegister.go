package model

import (
	database "ProjekGolangMVC/Database"
	node "ProjekGolangMVC/node"
	"time"
)

func InsertMahasiswaRegister(Email string, Pass string, Name string, Username string){
	var curLL *node.DataLL
	curLL = &database.DBMember

	newMahasiswa := node.DataRegisterMahasiswa{
		Email: Email,
		Password: Pass,
		Nama: Name,
		Username: Username,
		CraeteAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	newLL := node.DataLL{
		Member: newMahasiswa,
	}
	if curLL.Next == nil{
		curLL.Next = &newLL
	} else {
		for curLL.Next != nil {
			curLL = curLL.Next
		}
		curLL.Next = &newLL
	}
}

func SearchDataRegister(nama string) *node.DataLL{
	var curLL *node.DataLL
	curLL = &database.DBMember

	if curLL.Next != nil {
		for curLL.Next != nil{
			if curLL.Next.Member.Username == nama{
				return curLL.Next
			}
			curLL = curLL.Next
		}
	} else {
		return nil
	}
	return nil
}

func ReadAllDataRegister() []node.DataRegisterMahasiswa{
	var curLL *node.DataLL
	curLL = &database.DBMember

	var memberBoard []node.DataRegisterMahasiswa
	
	for curLL.Next != nil {
		curLL = curLL.Next
		memberBoard = append(memberBoard, curLL.Member)
	}
	return memberBoard
}

func SearchDataBelakang(nama string) (*node.DataLL, *node.DataLL) {
	curLL := &database.DBMember

		var prevNode *node.DataLL = nil
		for curLL != nil {
			// Periksa apakah curLL.Next dan curLL.Member tidak nil
			if curLL.Member.Username == nama {
				return prevNode, curLL
			}
			prevNode = curLL
			curLL = curLL.Next
		}
	
	return nil, nil
}

func UpdateDataRegister(Npmbaru string, Passbaru string, Namebaru string, Usernamekondisi string, searchResult *node.DataLL){
	searchResult.Member.Password = Passbaru
	searchResult.Member.Nama = Namebaru
	searchResult.Member.Email = Npmbaru
}

func DeleteDataRegister(prevNode *node.DataLL, currentNode *node.DataLL) bool {
	if currentNode == nil {
		return false // Tidak dapat menghapus jika node kosong
	}

	// Menghapus node yang ditunjuk oleh currentNode
	if prevNode == nil {
		// Jika prevNode nil, artinya currentNode adalah node pertama yang akan dihapus
		database.DBMember = *currentNode.Next
	} else {
		// Jika prevNode tidak nil, ubah pointer next dari prevNode
		prevNode.Next = currentNode.Next
	}

	return true
}

/*
Fungsi Delete menerima dua alamat memori: prevNode yang merupakan node
sebelum node yang akan dihapus, dan currentNode yang merupakan node
yang akan dihapus. Dengan memiliki kedua alamat ini, kita dapat mengubah
pointer Next dari prevNode agar menunjuk langsung ke node setelah currentNode,
sehingga menghapus currentNode dari linked list tanpa kehilangan struktur linked list
*/