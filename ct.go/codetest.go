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
		if err == nil {
			dir := string(drive) + ":/"
			read = append(read, dir)
			f.Close()
		}
	}
	return
}

func pathfile(dir string, f os.FileInfo, area map[string]string, files *[]file) {
	var process file
}