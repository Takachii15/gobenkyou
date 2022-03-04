package algorithm

import "fmt"

func DivSum() {
  var k, n int
  fmt.Scanf("%d %d", &n, &k)
  nums := make([]int, 0, 0)
  for i := 0; i < n; i++ {
    var input int
    fmt.Scanf("%d", &input)
    nums = append(nums, input)
  }
  count := 0
  for i := range(nums) {
    for j := i + 1; j < len(nums); j++ {
      if ((nums[i]+nums[j]) % k) == 0 {
          count++
      }
    }
  }
  fmt.Print(count)
}
