package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func walk(input string, deep bool, callback func(string)) error {
	return filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
		callback(path)
		return nil
	})
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		wd, err := os.Getwd()
		if err != nil {
			os.Exit(1)
		}
		err = walk(wd, false, func(path string) {
			log.Println(path)
		})

		if err != nil {
			os.Exit(1)
		}

	}
	inputArg := flag.Arg(0)
	pattern := flag.Arg(1)
	err := walk(inputArg, true, func(path string) {
		if strings.Contains(path, pattern) {
			log.Println(path)
		}
	})

	if err != nil {
		panic(err)
	}
}
