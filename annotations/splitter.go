package annotations

import (
	"github.com/funkygao/goannotation/engine"
	"go/ast"
	"log"
)

// Splitter is a sample annotation implementation applied on func.
type Splitter struct {
}

func (this *Splitter) Tag() string {
	return "@splitter"
}

func (this *Splitter) Execute(pkgName string, funcName string, decl *ast.FuncDecl) {
	// we can auto generate src code here
	log.Printf("pkg:%s func:%s decl:%#v\n", pkgName, funcName, *decl)
}

func init() {
	engine.RegisterPlugin(&Splitter{})
}
