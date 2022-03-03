package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Articles struct {
  Title string `json:"Title"`
  Desc string `json:"Description"`
  Cont string `json:"content"`
}

var Article []Articles

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests(){
    http.HandleFunc("/", homePage)
    http.HandleFunc("/article", returnArticles)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnArticles(w http.ResponseWriter, r *http.Request){
  fmt.Println("EndPoint reached")
  json.NewEncoder(w).Encode(Article)
}

func EndPoint(){
  var Article = []Articles{
    Articles{Title: "Hello", Desc:"sdjdjdj", Cont:"hdhdhdh"},
  }
  for k, _ := range(Article){
    fmt.Printf(Article[k])
  }
  handleRequests()
}
