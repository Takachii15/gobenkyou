package basic

import "fmt"

func TestPointer(){
  i,j := 42, 4071

  p := &i
  fmt.Println(*p)
  *p = 21
  fmt.Println(i)

  p = &j
  *p = *p / 37
  fmt.Println(j)
}

type node struct {
  adress int
}
