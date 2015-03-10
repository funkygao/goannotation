package engine

import (
	"fmt"
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
		if Debug {
			log.Printf("[decl] %#v\n", decl)
		}

		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Doc != nil {
			parseStructAnnotations(pkgName, genDecl)
		}

		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			parseFuncAnnotations(pkgName, funcDecl)
		}
	}
}

func parseStructAnnotations(pkgName string, genDecl *ast.GenDecl) {
	for _, comment := range genDecl.Doc.List {
		if Debug {
			log.Printf("[comment] %#v", *comment)
		}

		for _, annotation := range registeredAnnotations {
			if annotation.Type() != ANNOTATION_STRUCT {
				continue
			}

			if !strings.Contains(comment.Text, annotation.Tag()) {
				continue
			}

			// found the annotation
			if Debug {
				log.Println("found struct annotation tag:", annotation.Tag())
			}

			var typeName string
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if typeSpec.Name != nil {
						typeName = typeSpec.Name.Name
						break
					}
				}
			}

			structAnnotation, ok := annotation.(StructAnnotation)
			if !ok {
				panic(fmt.Sprintf("Execute() not found in annotation tag: %s",
					annotation.Tag()))
			}
			structAnnotation.Execute(pkgName, typeName)
		}
	}
}

func parseFuncAnnotations(pkgName string, funcDecl *ast.FuncDecl) {
	if funcDecl.Doc == nil {
		return
	}

	for _, comment := range funcDecl.Doc.List {
		for _, annotation := range registeredAnnotations {
			if annotation.Type() != ANNOTATION_FUNC {
				continue
			}

			if !strings.Contains(comment.Text, annotation.Tag()) {
				continue
			}

			// found the annotation
			if Debug {
				log.Println("found func annotation tag:", annotation.Tag())
			}

			funcAnnotation, ok := annotation.(FuncAnnotation)
			if !ok {
				panic(fmt.Sprintf("Execute() not found in annotation tag: %s",
					annotation.Tag()))
			}

			var funcName string
			if funcDecl.Name != nil {
				funcName = funcDecl.Name.Name
			}

			funcAnnotation.Execute(pkgName, funcName, funcDecl)
		}
	}
}
