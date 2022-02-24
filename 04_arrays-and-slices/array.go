// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices

package arrays

func SumArray(numbers [5]int) (result int) { // An interesting property of arrays is that the size is encoded in its type.
	for i := 0; i < 5; i++ {
		result += numbers[i]
	}
	return
}

func SumSlices(numbers []int) (result int) { // slices: https://go.dev/blog/slices-intro
	for _, number := range numbers { // _ is a 'blank identifier'. https://golang.org/doc/effective_go.html#blank
		result += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int { // variadic function: https://gobyexample.com/variadic-functions
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	// for i, numbers := range numbersToSum {
	// 	sums[i] = SumSlices(numbers)
	// }
	// return sums

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlices(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // Slices can be sliced! The syntax is slice[low:high]. If you omit the value on one of the sides of the : it captures everything to that side of it. In our case, we are saying "take from 1 to the end" with numbers[1:].
			sums = append(sums, SumSlices(tail))
		}
	}

	return sums
}
