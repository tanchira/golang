package main

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	deleteIndex := 2
	a = append(a[:deleteIndex], a[deleteIndex+1:]...)
}
