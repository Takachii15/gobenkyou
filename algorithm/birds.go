package algorithm

import "fmt"

func Birds() {
  var n int
  fmt.Scanf("%d", &n)
  addr := make([]int, 0, 0)
  for i := 0; i < n; i++ {
    var input int
    fmt.Scan(&input)
    addr = append(addr, input)
  }
  db := make(map[int]int)
  for i, id := range(addr) {
    for j := i + 1; j < len(addr); j++ {
      if id == addr[j] {
        db[id]++
      }
    }
  }
  fmt.Printf("%v", db)
}
