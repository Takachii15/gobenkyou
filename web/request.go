package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Town(){
  fmt.Println("test")
}

func TestHttp(){
  resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
  if err != nil {
    log.Fatalln(err)
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }
  sb := string(body)
  log.Printf(sb)
}
