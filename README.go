package main

import (
	"fmt"
)

type Customer struct {
	Name     string
	Email    string
	Username string
	Password string
	Bookings []*Booking
}

type Booking struct {
	RoomType string
	Duration int
}

type Room struct {
	RoomType  string
	Available int
}

func main() {
	customers := []*Customer{}
	rooms := []*Room{
		{RoomType: "Standard", Available: 5},
		{RoomType: "Deluxe", Available: 3},
		{RoomType: "Suite", Available: 2},
	}

	for {
		fmt.Println("================================================")
		fmt.Println("|   Selamat Datang di Sistem Pemesanan Hotel   |")
		fmt.Println("================================================")
		fmt.Println("1. Login Customer")
		fmt.Println("2. Login Pegawai")
		fmt.Println("3. Daftar Customer")
		fmt.Println("4. Melihat Pemesanan")
		fmt.Println("5. Memesan Kamar")
		fmt.Println("6. Keluar")
		var choice int
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("=============")
			fmt.Println("|   Login   |")
			fmt.Println("=============")
			fmt.Print("Username: ")
			var username string
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			var password string
			fmt.Scanln(&password)

			customer := findCustomerByUsername(customers, username)
			if customer == nil || customer.Password != password {
				fmt.Println("Login gagal. Mohon periksa kembali username dan password Anda.")
			} else {
				fmt.Printf("Selamat datang, %s!\n", customer.Name)
				showBooking(customer)
			}
		case 2:
			fmt.Println("=====================")
			fmt.Println("|   Login Pegawai   |")
			fmt.Println("=====================")
			fmt.Print("Username: ")
			var username string
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			var password string
			fmt.Scanln(&password)

			customer := findCustomerByUsername(customers, username)
			if customer == nil || customer.Password != password {
				fmt.Println("Login gagal. Mohon periksa kembali username dan password Anda.")
			} else {
				fmt.Printf("Selamat datang, %s!\n", customer.Name)
				showBooking(customer)
			}
		case 3:
			fmt.Println("==============")
			fmt.Println("|   Daftar   |")
			fmt.Println("==============")
			fmt.Print("Nama: ")
			var name string
			fmt.Scanln(&name)
			fmt.Print("Email: ")
			var email string
			fmt.Scanln(&email)
			fmt.Print("Username: ")
			var username string
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			var password string
			fmt.Scanln(&password)

			customer := &Customer{Name: name, Email: email, Username: username, Password: password}
			customers = append(customers, customer)
			fmt.Println("Registrasi berhasil. Silakan login.")
		case 4:
			fmt.Println("=========================")
			fmt.Println("|   Melihat Pemesanan   |")
			fmt.Println("=========================")
			fmt.Print("Username: ")
			var username string
			fmt.Scanln(&username)

			customer := findCustomerByUsername(customers, username)
			if customer == nil {
				fmt.Println("Pelanggan tidak ditemukan.")
			} else {
				showBooking(customer)
			}
		case 5:
			fmt.Println("=====================")
			fmt.Println("|   Memesan Kamar   |")
			fmt.Println("=====================")
			fmt.Print("Username: ")
			var username string
			fmt.Scanln(&username)

			customer := findCustomerByUsername(customers, username)
			if customer == nil {
				fmt.Println("Pelanggan tidak ditemukan.")
			} else {
				fmt.Println("Pilihan Kamar:")
				for i, room := range rooms {
					fmt.Printf("%d. %s (%d tersedia)\n", i+1, room.RoomType, room.Available)
				}
				fmt.Print("Pilihan Anda: ")
				var roomChoice int
				fmt.Scanln(&roomChoice)
				if roomChoice < 1 || roomChoice > len(rooms) {
					fmt.Println("Pilihan kamar tidak valid.")
				} else {
					room := rooms[roomChoice-1]
					if room.Available == 0 {
						fmt.Println("Kamar yang dipilih sudah habis.")
					} else {
						fmt.Print("Durasi menginap (hari): ")
						var duration int
						fmt.Scanln(&duration)

						booking := &Booking{RoomType: room.RoomType, Duration: duration}
						customer.Bookings = append(customer.Bookings, booking)
						room.Available--
						fmt.Println("Pemesanan kamar berhasil.")
					}
				}
			}
		case 6:
			fmt.Println("Terima kasih telah menggunakan sistem pemesanan hotel. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}

// customer
func findCustomerByUsername(customers []*Customer, username string) *Customer {
	for _, customer := range customers {
		if customer.Username == username {
			return customer
		}
	}
	return nil
}

// pemesanan
func showBooking(customer *Customer) {
	fmt.Printf("Pelanggan: %s\n", customer.Name)
	if len(customer.Bookings) > 0 {
		fmt.Println("Daftar Pemesanan:")
		for i, booking := range customer.Bookings {
			fmt.Printf("%d. Kamar %s, Durasi %d hari\n", i+1, booking.RoomType, booking.Duration)
		}
	} else {
		fmt.Println("Belum ada pemesanan.")
	}
}
