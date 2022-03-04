package main

import (
	"core/article3/q2/lib"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeeting object")
}

func main() {
	flag.Parse()

	lib.Hello(name)
}
