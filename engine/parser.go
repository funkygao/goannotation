package engine

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func ParseFile(fn string) {
	f, err := parser.ParseFile(token.NewFileSet(), fn, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	pkgName := f.Name.Name
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Doc == nil {
			// skip it
			continue
		}

		for _, comment := range genDecl.Doc.List {
			for _, plugin := range registeredAnnotations {
				if strings.Contains(comment.Text, plugin.Tag()) {
					// found the annotation
					log.Println("found:", plugin)

					for _, spec := range genDecl.Specs {
						switch plugin.Type() {
						case ANNOTATION_TYPE:
							parseTypeAnnotation(pkgName, spec)
						case ANNOTATION_FUNC:
							parseFuncAnnotation(pkgName, spec)
						}
					}
				}

			}
		}

	}

	return

}

func parseTypeAnnotation(pkgName string, spec ast.Spec) {

}

func parseFuncAnnotation(pkgName string, spec ast.Spec) {

}
