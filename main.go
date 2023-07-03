package main

import (
	"fmt"
	"os"
	"plugin"
)

func main() {
	_, err := plugin.Open("./plugin.so")
	if err != nil {
		fmt.Fprintf(os.Stderr, "loading plugin: %v\n", err)
	}
}
