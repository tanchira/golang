package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for fn := 1; fn <= 3000; fn++ {
		file, err := os.Create(fmt.Sprintf("%v.txt", fn))
		if err != nil {
			return
		}

		defer file.Close()
		for name := 1; name <= 1000; name++ {
			file.WriteString("tanchira \nPeawkrathok \n15\n")
		}

	}
	fmt.Println(time.Since(start))
}
