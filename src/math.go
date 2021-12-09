package main

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Distance(from [2]int, to [2]int) [2]int {
	return [2]int{
		Abs(from[0] - to[0]),
		Abs(from[1] - to[1]),
	}
}