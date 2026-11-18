package main
import "fmt"

func main() {
	var jenis_daging, menu_masakan string
	fmt.Print("Daging : ")
	fmt.Scan(&jenis_daging)
	fmt.Print("Menu : ")
	fmt.Scan(&menu_masakan)
	if jenis_daging == "ayam" {
		if menu_masakan == "sop"{
			fmt.Println("Daging : Ayam")
			fmt.Println("Tambahan : Sayur")
		
	 }else if menu_masakan == "katsu" {
		fmt.Println("Daging : Ayam")
		fmt.Println("Tambahan : Tepung")
	}

}
	if jenis_daging == "sapi" {
		if menu_masakan == "teriyaki" {
			fmt.Println("Daging : Sapi")
			fmt.Println("Tambahan : Bombay")
		} else if menu_masakan == "semur" {
			fmt.Println("Daging : Sapi")
			fmt.Println("Tambahan : Kecap")
		}
	}
}
