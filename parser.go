package main

import (
	"github.com/funkygao/goannotation/engine"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func ParseFile(fn string) (pkgName string) {
	f, err := parser.ParseFile(token.NewFileSet(), fn, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	pkgName = f.Name.Name
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Doc == nil {
			// skip it
			continue
		}

		for _, comment := range genDecl.Doc.List {
			for _, plugin := range engine.Plugins {
				if strings.Contains(comment.Text, plugin.AnnotationTag()) {
					// found the annotation
					log.Println("found:", plugin)

					for _, spec := range genDecl.Specs {
						switch plugin.AnnotationType() {
						case engine.TAG_TYPE:
							parseTypeAnnotation(spec)
						case engine.TAG_FUNC:
							parseFuncAnnotation(spec)
						}
					}
				}

			}
		}

	}

	return

}

func parseTypeAnnotation(spec ast.Spec) {

}

func parseFuncAnnotation(spec ast.Spec) {

}
