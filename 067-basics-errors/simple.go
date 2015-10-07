package main

import (
	"fmt"
)

type NotDigitError struct {
}

func (error *NotDigitError) Error() string {
	return "You need to pass numbers from 0 to 9"
}

func SumDigits(nums []int) (int, error) {
	sum := 0
	for _, num := range nums {
		if num > 10 {
			return 0, &NotDigitError{}
		}
		sum += num
	}

	return sum, nil
}

func main() {
	nums, err := SumDigits([]int{1, 2, 3, 4, 5, 5, 6, 7})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(nums)
}
