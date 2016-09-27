package annotations

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/funkygao/goannotation/engine"
	"github.com/funkygao/golib/color"
)

var _ engine.FuncAnnotation = &Rest{}

// Rest is an annotation implementation applied on a REST controller function.
//
// Format of @rest:
// {method} {uri}
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
		fmt.Printf("%s\n", funcName)

		parts := strings.SplitN(doc, " ", 2)
		if len(parts) == 1 {
			fmt.Printf("    %s\n", color.Yellow("%8s", parts[0]))
		} else {
			fmt.Printf("    %s %s\n", color.Green("%8s", parts[0]), parts[1])

		}
	}
}

func init() {
	engine.RegisterPlugin(&Rest{})
}
