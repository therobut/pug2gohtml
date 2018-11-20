package main

import (
	"fmt"
	"github.com/eknkc/pug"
	"io/ioutil"
	"os"
)

func main() {
	file := os.Args[1]
	lenFile := len(file)
	if file[lenFile-4:lenFile] == ".pug" {
		file = file[0 : lenFile-4]
	}
	fmt.Println(file)

	opt := pug.Options{
		PrettyPrint: true,
	}

	html, err := pug.ParseFile(file+".pug", opt)
	if err != nil {
		panic(err)
	}

	newFile := file + ".html"
	if err := ioutil.WriteFile(newFile, []byte(html), 0644); err != nil {
		panic(err)
	}

	fmt.Println("All done!")
}
