package main

import (
	"flag"
)

var (
	port int = 6017
)

func init() {
	flag.IntVar(&prot, "p", 6017, "")
}

func main() {
	flag.Parse()
}
