package main

import (
	"errors"
	"fmt"
)

func SumDigits(nums []int) (int, error) {
	sum := 0
	for _, num := range nums {
		if num > 10 {
			return 0, errors.New("aaaaaaaaaa")
			// errors.New creates new `error` instance
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
