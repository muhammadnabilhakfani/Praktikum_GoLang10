package main
import "fmt"

func main() {
	var a, b, durasi int
	fmt.Scan(&a, &b)

	if (a >= 0 && a <=23) && (b >=0 && b <= 23) {
	if b < a {
		durasi = (24 - a) + b
	} else {
		durasi = b - a
	}
	}

	fmt.Println(durasi)
}
