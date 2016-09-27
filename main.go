package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/funkygao/goannotation/annotations"
	"github.com/funkygao/goannotation/engine"
)

func init() {
	flag.BoolVar(&engine.Debug, "d", false, "debug")
	flag.Parse()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	for _, fn := range os.Args[1:] {
		engine.ParseFile(fn)
	}

}
