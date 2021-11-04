package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	argsWithProg := os.Args[1:]

	if len(argsWithProg) < 2 {
		log.Fatalf("please provide a file path and a package name. format is wowclassgen file.go mypackage. file location must be relative to the current directory")
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	renderFile(path, argsWithProg)
}

func renderFile(path string, args []string) {
	code, err := generateClass(args[1])
	if err != nil {
		log.Fatalf("err : %v", err)
		return
	}

	file, err := os.Create(fmt.Sprintf("%s/%s", path, args[0]))
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	err = code.Render(file)
	if err != nil {
		log.Fatalln(err)
	}
}
