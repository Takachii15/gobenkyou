package main

import (
	// "fmt"
	// "gobenkyou/basic"
	//  "gobenkyou/web"
	"fmt"
	// "gobenkyou/hennge"
)

func main(){
  // basic.TestDefer()
  // var test basic.TestDb
  // test.Name = "Daffa"
  // fmt.Println(test)
  // basic.TestPointer()
  test := new(string)
  *test = "test"
  // web.TestHttp()
  // basic.Run()
  fmt.Println(*test)
}

