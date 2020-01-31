package main

import (
	"fmt"
	"os"
)

func walkFn(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
}
