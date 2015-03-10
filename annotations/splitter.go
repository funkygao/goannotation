package annotations

import (
	"github.com/funkygao/goannotation/engine"
	"go/ast"
	"log"
)

type Splitter struct {
}

func (this *Splitter) Tag() string {
	return "@splitter"
}

func (this *Splitter) Type() engine.AnnotationType {
	return engine.ANNOTATION_STRUCT
}

func (this *Splitter) Execute(pkgName string, funcName string, decl *ast.FuncDecl) {
	// we can auto generate src code here
	log.Println(pkgName, funcName)
	log.Printf("%#v\n", *decl)
}

func init() {
	engine.RegisterPlugin(&Splitter{})
}
