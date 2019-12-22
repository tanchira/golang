package main

func handlepanic()  {
	fmt.Println("Hello World")
}
func main() {
	defer()
	panic("Hello panic")
}
