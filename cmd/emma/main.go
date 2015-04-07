package main

import (
	"fmt"
	"os"

	"github.com/ruedap/go-emma"
)

func main() {
	ret := emma.Find(emma.Src, os.Args[1:])
	str, err := emma.ToJSON(ret)
	if err != nil {
		fmt.Println("Failed to output json.")
	}

	fmt.Println(str)
}
