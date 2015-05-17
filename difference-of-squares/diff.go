package diffsquares

func SquareOfSums(i int) int {
	sum := (i * (i + 1)) / 2
	return sum * sum
}

func SumOfSquares(i int) int {
	return (i * (i + 1) * (2*i + 1)) / 6
}

func Difference(i int) int {
	return SquareOfSums(i) - SumOfSquares(i)
}
