package main

func main() {
	human := []string{"bulma", "chi-chi", "videl"}
	name := []string{"goku", "gohan"}
	name = append(human, name...)
}
