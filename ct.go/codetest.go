package main

import "os"

type file struct {
	name string
	size int
	path string
}

func driver() (read []string) {
	for _, drive := range "D:" {
		f, err := os.Open(string(drive) + ":\\")
	}

}
