package basic

import(
  "fmt"
)

func Test(){
  fmt.Println("こんにちわ")
}

func Addition(x float64, n int) float64 {
  for i:= 0; i < n; i++{
    x = x + 1
  }
  return x
}

type Fun interface {
  test()
  addition()
}
