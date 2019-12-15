package main

func sum(number ...int) int {
	total := 0
	for _, n := range number {
		total = total + n
	}
}

func main() {
	include := sum(1, 3, 5, 7, 9)
}
