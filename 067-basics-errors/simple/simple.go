package main

import (
	"fmt"
)

func SumDigits(nums []int) (int, error) {
	sum := 0
	for _, num := range nums {
		if num >= 10 {
			return 0, fmt.Errorf("Number is too big: %d, i need number <10", num)
			// errors.New creates new `error` instance
			// we can use also:
			// fmt.Errorf("Some Error ID:%d", 102323812389)
		}
		sum += num
	}

	return sum, nil
}

func main() {
	nums, err := SumDigits([]int{1, 2, 3, 4, 8885, 5, 6, 7})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(nums)
}
