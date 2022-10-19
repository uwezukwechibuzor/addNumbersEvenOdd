package main

import (
	"reflect"
	"testing"
)

func TestAddNumbersEvenOdd(t *testing.T) {
	testCases := []struct {
		Name                string
		Values              []int
		ExpectedSum         int
		ExpectedEvenNumbers []int
		ExpectedOddNumbers  []int
	}{
		{
			"addNumbersEvenOdd() -> 0",
			[]int{},
			0,
			[]int{},
			[]int{},
		},
		{
			"addNumbersEvenOdd([]int{10, 20, 100}) -> The Sum is 130 and all is even numbers",
			[]int{10, 20, 100},
			130,
			[]int{10, 20, 100},
			[]int{},
		},
		{
			"addNumbersEvenOdd([]int{101, 20, 100}) -> The Sum is 221, [20, 100] are even numbers and [101] is odd number",
			[]int{101, 20, 100},
			221,
			[]int{20, 100},
			[]int{101},
		},
		{
			"addNumbersEvenOdd([]int{1, 2, 3}) -> The Sum is 6, [2] is even number and [1, 3] are odd numbers",
			[]int{1, 2, 3},
			6,
			[]int{2},
			[]int{1, 3},
		},
		{
			"addNumbersEvenOdd([]int{-10, -2, -3}) -> testing for Negative Numbers",
			[]int{-10, -2, -3},
			-15,
			[]int{-10, -2},
			[]int{-3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			sum, evenNumbers, oddNumbers := addNumbersEvenOdd(tc.Values...)
			if sum != tc.ExpectedSum || !reflect.DeepEqual(evenNumbers, tc.ExpectedEvenNumbers) || !reflect.DeepEqual(oddNumbers, tc.ExpectedOddNumbers) {
				t.Errorf("%d != %d", sum, tc.ExpectedSum)
				t.Errorf("%d != %d", evenNumbers, tc.ExpectedEvenNumbers)
				t.Errorf("%d != %d", oddNumbers, tc.ExpectedOddNumbers)
			} else {
				t.Logf("%d == %d", sum, tc.ExpectedSum)
				t.Logf("%d == %d", evenNumbers, tc.ExpectedEvenNumbers)
				t.Logf("%d == %d", oddNumbers, tc.ExpectedOddNumbers)
			}
		})
	}
}

func TestRunMain(t *testing.T) {
	main()
}
