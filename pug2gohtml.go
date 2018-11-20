package main

import (
	"io/ioutil"
	"os"

	"github.com/eknkc/pug"
)

func main() {
	// get file name from command line argument
	file := os.Args[1]

	// check for .pug extension and remove it if present
	lenFile := len(file)
	if file[lenFile-4:lenFile] == ".pug" {
		file = file[0 : lenFile-4]
	}

	// tell pug parser we want pretty html
	opt := pug.Options{
		PrettyPrint: true,
	}

	// parse the pug file
	html, err := pug.ParseFile(file+".pug", opt)
	if err != nil {
		panic(err)
	}

	// write the new HTML file
	newFile := file + ".html"
	if err := ioutil.WriteFile(newFile, []byte(html), 0644); err != nil {
		panic(err)
	}
}
