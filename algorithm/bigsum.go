package algorithm

import "fmt"

func Bigsum() {
  var n int
  fmt.Scan(&n)
  nums := make([]int, 0, 0)
  for i := 0; i < n; i++ {
    var input int
    fmt.Scanf("%d", &input)
    nums = append(nums, input)
  }
  res := 0
  for _, num := range(nums) {
    res = res + num
  }
  fmt.Print(res)
}
