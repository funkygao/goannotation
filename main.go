package main

import (
	_ "github.com/funkygao/goannotation/plugins"
	"os"
)

func main() {
	for _, fn := range os.Args[1:] {
		ParseFile(fn)
	}

}
