# pug2gohtml

Small program to convert Pug templates to Go HTML templates.

Only tested with Pug written for Go using github.com/eknkc/pug

I've recently become frustrated with the limitations of using Pug in Go, and the library doesn't seem to be getting any updates. Since I was using it on a project for work, and don't have the time to contribute, I just abandoned using it altogether. 

This small program will take your Pug code and convert it into a Go html template. You will still have to go in and tweak some things, such as if statements, but it's much better than having to rewrite the entire thing by hand.

## Install
Make sure your GOPATH is set up properly
- `go get github.com/therobut/pug2gohtml`
- `go install github.com/therobut/pug2gohtml`

## Usage
`pug2gohtml <filename>`

### Example
`pug2gohtml index.pug` (file extension optional)

This will generate index.html. Make sure to check any template code within the produced file. I have found that the `{{...}}` statements can be a little garbled and may not work without some tweaking. This is due to the library used for conversion.