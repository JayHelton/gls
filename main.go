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
		if info != nil && info.IsDir() {
			callback("\033[35m" + path + "\033[0m")

			if input != path && !deep {
				return filepath.SkipDir
			}

			return nil
		}
		callback(path)

		return nil
	})
}

func main() {
	deepPtr := flag.Bool("deep", false, "check")

	flag.Parse()

	if flag.NArg() == 0 {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		err = walk(wd, *deepPtr, func(path string) {
			log.Println(path)
		})

		if err != nil {
			panic(err)
		}

	}

	inputArg := flag.Arg(0)
	pattern := flag.Arg(1)

	err := walk(inputArg, *deepPtr, func(path string) {
		if strings.Contains(path, pattern) {
			log.Println("\033[31m" + path + "\033[0m")
		}
	})

	if err != nil {
		panic(err)
	}
}
