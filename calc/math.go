package calc

import "errors"

//Add returns...
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		return sum, errors.New("provide more than two numbers")
	} else {

		for _, num := range numbers {
			sum = sum + num
		}

		return sum, nil
	}
}
