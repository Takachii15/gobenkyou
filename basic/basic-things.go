package basic

import (
	"fmt"
)

// Using print functionality
func Test() {
	fmt.Println("こんにちわ")      // return こんにちわ
	fmt.Println(len("こんにちわ")) // return 5
	fmt.Println("こんにちわ"[2])   // return に
}

func Addition(x float64, n int) float64 {
	for i := 0; i < n; i++ {
		x = x + 1
	}
	return x
}

func Factorial(x int) int {
	var z byte = 8
	fmt.Println(z)
	return 0
}

func Returnthings(x, y int, name string) string {
  return fmt.Sprintf(name, x, y, "mas %s, umur %d lahir tanggal %d")
}

type TestDb struct {
	Name, Adress string
	Age          int
}
