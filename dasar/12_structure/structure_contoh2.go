package main

import "fmt"

type Buku struct {
	judul   string
	penulis string
	id_buku int
}

func main() {
	var Buku1 Buku
	var Buku2 Buku

	Buku1.judul = "harry potter - the goblet of fire"
	Buku1.penulis = "jk rowling"
	Buku1.id_buku = 15241

	Buku2.judul = "harry potter - the prisoner of azkaban"
	Buku2.penulis = "jk rowling"
	Buku2.id_buku = 15242

	tampilBuku(Buku1)
	tampilBuku(Buku2)

}

func tampilBuku(buku Buku) {
	fmt.Printf("judul buku %s\n", buku.judul)
	fmt.Printf("penulis %s\n", buku.penulis)
	fmt.Printf("id buku %d\n", buku.id_buku)
}
