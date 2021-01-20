package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "go utils", "帮助信息")
	flag.Parse()
	log.Printf("name: %s", name)

	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar()
}
