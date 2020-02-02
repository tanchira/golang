package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var bigfile, smallfile []string

func getDrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			d := string(drive) + ":/"
			r = append(r, d)
			f.Close()
		}
	}
	return
}
func FindFileFromExtension(extension []string, dir string, files *[]string) {
	fs, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, f := range fs {
			for _, ex := range extension {
				if strings.HasSuffix(f.Name(), ex) {
					if (time.Now().Unix() - f.ModTime().Unix()) < 2592000 {
						*files = append(*files, dir+"/"+f.Name())
						if f.Size() > 10240 {
							bigfile = append(bigfile, dir+"/"+f.Name())
						} else {
							smallfile = append(smallfile, dir+"/"+f.Name())
						}

					}
				}
			}
			if f.IsDir() {
				path := dir + "/" + f.Name()
				FindFileFromExtension(extension, path, files)
			}

		}
	}
}

func main() {
	start := time.Now()
	drives := getDrives()
	files := []string{}
	for _, d := range drives {
		FindFileFromExtension([]string{".jpg", ".gif", ".png", ".bmp"}, d, &files)
	}
	fmt.Println("file : ", len(files))
	fmt.Println("bigfile : ", len(bigfile))
	fmt.Println("smallfile : ", len(smallfile))
	fmt.Println("time :", time.Since(start))

}
