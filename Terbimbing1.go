package main
import "fmt"

func main() {
	var s, p int
	fmt.Scan(&s, &p)
	if (s >= 0 && s <= 1000000) && (p >= 0 && p <= 1000000) {
		if s > p {
			fmt.Print("PEMBAYARAN BERHASIL")
	} else if s == p {
		 fmt.Print("SALDO SANGAT PAS")
	} else {
		fmt.Print("SALDO TIDAK CUKUP")
	}
}
}
