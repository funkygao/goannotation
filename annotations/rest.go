package annotations

import (
	"fmt"
	"go/ast"

	"github.com/funkygao/goannotation/engine"
)

var _ engine.FuncAnnotation = &Rest{}

// Rest is an annotation implementation applied on a REST controller function.
type Rest struct {
	attrs map[string]string
}

func (this *Rest) Tag() string {
	return "@rest"
}

func (this *Rest) SetAttrs(kv map[string]string) {
	this.attrs = kv
}

func (this *Rest) Execute(pkgName string, funcName string, decl *ast.FuncDecl) {
	if doc, ok := this.attrs["doc"]; ok {
		fmt.Println(doc)
	}
}

func init() {
	engine.RegisterPlugin(&Rest{})
}
