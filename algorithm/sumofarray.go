package algorithm

import "fmt"

// Sum of Array Hackerrank
func SumOfArray() {
	var cnt, input int
  num := make([]int, 0, 4)
	fmt.Scanf("%v", &cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scanf("%d", &input)
		num = append(num, input)
	}
  result := 0
  for i := range(num) {
    result = result + num[i]
  }
	fmt.Print(result)
}
