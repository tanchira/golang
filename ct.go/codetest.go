package main

import (
	"io/ioutil"
	"os"
)

type file struct {
	name string
	size int64
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
	process.name = f.Name()
	process.size = f.Size()
	process.path = dir + "/" + f.Name()
	*files = append(*files, process)
}

func findFileFromExtention(area map[string]string, dir string, files *[]file) {
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	for _, f := range fs {
		if f.IsDir() {
			path := dir + "/" + f.Name()
			findFileFromExtention(area, path, files)
		}
	}

}
