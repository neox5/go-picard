package goproject

var mainGoTmpl = `package main

import "fmt"

func main() {
	fmt.Println("{{.}}")
}
`

var makefileTmpl = `build:
	@go build -o=./bin/ ./cmd/...
`

var gitignoreTmpl = `bin`
