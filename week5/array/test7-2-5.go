package main

import "fmt"

func main() {
	number := [2][5][2]int{
		{
			{1, 2},
			{10, 20},
			{100, 200},
			{1000, 2000},
			{10000, 2000},
		},
		{
			{1, 2},
			{10, 20},
			{100, 200},
		},
	}
	fmt.Println(number)
}
