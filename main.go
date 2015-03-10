package main

import (
	"github.com/funkygao/goannotation/engine"
	_ "github.com/funkygao/goannotation/plugins"
	"os"
)

func main() {
	for _, fn := range os.Args[1:] {
		engine.ParseFile(fn)
	}

}
