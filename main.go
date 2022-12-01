package main

import (
	"fmt"
)

func main(){
	type items struct{
		id int
		name string
		price int
	}

	type cart struct{
		item_id int
		amount int
		subtotal int
	}

	var allItems = []items{
		{id: 1, name: "Buku", price: 4000},
		{id: 2, name: "Kacang", price: 5000},
		{id: 3, name: "Wafer", price: 3000},
		{id: 4, name: "Permen", price: 3000},
	}

	var carts = []cart{}
	total := 0

	for {
		fmt.Println("Halo, selamat datang. Silahkan pilih barang yang ingin Anda beli")
		
		for _, item := range allItems{
			fmt.Println(item.id, ". ", item.name, " - ", item.price)
		}

		var pick int
		var amount int

		fmt.Scan(&pick)
		fmt.Print("Anda akan membeli ", allItems[pick-1].name, ", dengan harga ", allItems[pick-1].price, ", masukkan jumlah: ")
		fmt.Scanln(&amount)
		fmt.Println("Pembelian sukses, subtotal: ", allItems[pick-1].price * amount)
		var subtotal int = allItems[pick-1].price * amount

		carts = append(carts, cart{allItems[pick-1].id, amount, subtotal})
		total+=subtotal

		var option string
		fmt.Print("Mau beli lagi? (Y/N)")
		fmt.Scan(&option)
		if option == "N" || option == "n" {
			break
		}
	}

	fmt.Println("Total belanja Anda ", total)
	var cash int
	fmt.Println("Masukkan total bayar")
	fmt.Scanln(&cash)
	check := cash - total
	if check != 0 {
		fmt.Println("Pembayaran sukses, uang kembalian", check)
	}else{
		fmt.Println("Pembayaran sukses")
	}
}