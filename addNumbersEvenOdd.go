package main

import "fmt"

//function that returns even numbers, odd numbers and sum of all the numbers
func addNumbersEvenOdd(numbers ...int) (sum int, evenNumbers, oddNumbers []int) {
	sum = 0
	evenNumbers = []int{}
	oddNumbers = []int{}

	for _, number := range numbers {
		sum += number
		if number != 0 {
			if number%2 == 0 {
				evenNumbers = append(evenNumbers, number)
			} else {
				oddNumbers = append(oddNumbers, number)
			}
		}
	}
	return sum, evenNumbers, oddNumbers
}

func main() {
	//call the function here
	sum, evenNumbers, oddNumbers := addNumbersEvenOdd(1, 11, 21, 30, 40)
	sum_, evenNumbers_, oddNumbers_ := addNumbersEvenOdd(-1, -11, 21, 30, 40, 200, 800, 18, -20)

	fmt.Println(sum, evenNumbers, oddNumbers)
	fmt.Println(sum_, evenNumbers_, oddNumbers_)
}
