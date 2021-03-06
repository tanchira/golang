package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type file struct {
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
		} else {
			pathfile(dir, f, area, files)
		}
	}

}
func createfile(path, size []string) {
	file, err := os.Create("Output.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	for info := range path {
		file.WriteString(size[info] + "\t" + path[info] + "\n")
	}
}
func main() {
	myfiles := []file{}
	areaMap := make(map[string]string)
	drivers := driver()

	for _, drive := range drivers {
		findFileFromExtention(areaMap, drive, &myfiles)

	}
	var pathfile, sizefile []string
	for _, pathtofile := range myfiles {
		pathfile = append(pathfile, pathtofile.path)
		sizefile = append(sizefile, strconv.Itoa(int(pathtofile.size)))

	}
	createfile(pathfile, sizefile)
	fmt.Println("file", len(myfiles))

}
