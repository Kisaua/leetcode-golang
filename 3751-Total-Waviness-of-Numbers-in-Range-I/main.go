package main

func totalWaviness(num1 int, num2 int) int {
	count := 0

	for i := num1; i <= num2; i++ {
		count += waveniess(i)
	}
	return count
}

func waveniess(num int) int {
	count := 0
	for num/100 > 0 {
		right := num % 10
		center := (num % 100) / 10
		left := (num % 1000) / 100

		num = num / 10

		if (right < center && left < center) || (right > center && left > center) {
			count++
		}
	}

	return count
}
