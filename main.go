package main

import "fmt"

type barang struct {
	namaBarang  string
	hargaBarang int
	stokBarang  int
}

func main() {

	toko := make(map[string]barang)

	toko["Indomaret"] = barang{namaBarang: "Tisu", hargaBarang: 5000, stokBarang: 11}
	toko["Alfamart"] = barang{namaBarang: "Bebelac", hargaBarang: 40000, stokBarang: 5}
	toko["Alfamidi"] = barang{namaBarang: "Mie Sakura", hargaBarang: 5000, stokBarang: 21}
	toko["CircleK"] = barang{namaBarang: "Susu Superman", hargaBarang: 5000, stokBarang: 101}
	toko["Lawson"] = barang{namaBarang: "Chitos", hargaBarang: 5000, stokBarang: 8}

	fmt.Println("Berikut adalah daftar toko dengan barang yang stoknya dibawah 10: ")

	for index, i := range toko {
		if i.stokBarang < 10 {
			fmt.Printf("%s : %s dengan stok %d\n", index, i.namaBarang, i.stokBarang)
		}
	}

}
