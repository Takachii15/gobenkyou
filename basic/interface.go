package basic

import(
  "fmt"
)

func test(){
  fmt.Println("こんにちわ")
}

func addition(x float64, n int) float64 {
  for i:= 0; i < n; i++{
    x = x + 1
  }
  return x
}

