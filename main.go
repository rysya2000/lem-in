package main

import (
	"fmt"
	"os"

	. "lemin/functions"
)

func main() {

	text, err := OpenFile(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	GraphConstruct(text)

}
