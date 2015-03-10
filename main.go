package main

import (
	_ "github.com/funkygao/goannotation/annotations"
	"github.com/funkygao/goannotation/engine"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	engine.Debug = false

	for _, fn := range os.Args[1:] {
		engine.ParseFile(fn)
	}

}
