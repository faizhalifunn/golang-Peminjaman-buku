package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Username string
	Nama     string
	NPM      string
	History  []Jumlahpinjam
}

type Book struct {
	Title string
	Stock int
}

type Jumlahpinjam struct {
	Title    string
	Quantity int
}

var users = []User{
	{"Faiz", "Faiz Zhalifun Nafzi", "2406425312", []Jumlahpinjam{}},
}

var books = []Book{
	{"Pemrograman", 10},
	{"Film", 5},
	{"Printing", 20},
	{"Musik", 15},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("===== Program Peminjaman Buku =====")

	fmt.Print("Masukkan Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Masukkan Password (NPM): ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	var currentUser *User
	for i, user := range users {
		if user.Username == username && user.NPM == password {
			currentUser = &users[i]
			break
		}
	}

	if currentUser == nil {
		fmt.Println("Username atau Password salah!")
		return
	}

	for {
		fmt.Println("\n===== Menu =====")
		fmt.Println("1. Lihat Informasi Pengguna Program")
		fmt.Println("2. Lihat Daftar Buku")
		fmt.Println("3. Tambah Daftar Buku")
		fmt.Println("4. Pinjam Buku")
		fmt.Println("5. Histori Peminjaman Buku")
		fmt.Println("6. Keluar")

		fmt.Print("Pilih menu (1-6): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			lihatInformasiPengguna(currentUser)
		case "2":
			lihatDaftarBuku()
		case "3":
			tambahDaftarBuku(reader)
		case "4":
			tambahPeminjamanBuku(currentUser, reader)
		case "5":
			historiPeminjamanBuku(currentUser)
		case "6":
			fmt.Println("Keluar dari Program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		pressEnterToContinue(reader)
	}
}

func lihatInformasiPengguna(user *User) {
	fmt.Println("===== Informasi Pengguna =====")
	fmt.Println("Nama:", user.Nama)
	fmt.Println("Username:", user.Username)
	fmt.Println("NPM:", user.NPM)
}

func lihatDaftarBuku() {
	fmt.Println("===== Daftar Buku =====")
	for i, book := range books {
		status := "Tersedia"
		if book.Stock == 0 {
			status = "Stok Habis"
		}
		fmt.Printf("%d. %s [Stok: %d] [%s]\n", i+1, book.Title, book.Stock, status)
	}
}

func tambahDaftarBuku(reader *bufio.Reader) {
	fmt.Print("Masukkan judul buku baru: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Masukkan jumlah stok buku: ")
	var stock int
	fmt.Scanf("%d", &stock)

	if stock <= 0 {
		fmt.Println("Input Stok Tidak Valid")
	} else {
		books = append(books, Book{Title: title, Stock: stock})
		fmt.Printf("Buku '%s' berhasil ditambahkan dengan stok %d.\n", title, stock)
	}

	reader.ReadString('\n')
}

func tambahPeminjamanBuku(user *User, reader *bufio.Reader) {
	lihatDaftarBuku()

	fmt.Print("Pilih nomor buku yang ingin dipinjam: ")
	var choice int
	fmt.Scanf("%d", &choice)

	reader.ReadString('\n')

	if choice <= 0 || choice > len(books) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	selectedBook := &books[choice-1]

	if selectedBook.Stock > 0 {
		fmt.Printf("Masukkan jumlah buku yang ingin dipinjam (max %d): ", selectedBook.Stock)
		var quantity int
		fmt.Scanf("%d", &quantity)

		reader.ReadString('\n')

		if quantity <= 0 {
			fmt.Println("Input tidak valid")
			return
		}

		if quantity > selectedBook.Stock {
			fmt.Println("Jumlah buku yang diminta melebihi stok yang tersedia.")
		} else {
			selectedBook.Stock -= quantity
			user.History = append(user.History, Jumlahpinjam{Title: selectedBook.Title, Quantity: quantity})
			fmt.Printf("Berhasil meminjam %d buku.\n", quantity)
		}
	} else {
		fmt.Println("Stok buku habis.")
	}
}

func historiPeminjamanBuku(user *User) {
	fmt.Println("===== Histori Peminjaman Buku =====")
	if len(user.History) == 0 {
		fmt.Println("Belum ada buku yang dipinjam.")
	} else {
		for i, record := range user.History {
			fmt.Printf("%d. %s [%d]\n", i+1, record.Title, record.Quantity)
		}
	}
}

func pressEnterToContinue(reader *bufio.Reader) {
	fmt.Print("\nTekan 'Enter' untuk kembali ke menu...")
	reader.ReadString('\n')
}
