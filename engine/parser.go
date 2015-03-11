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
			structAnnotation, ok := annotation.(StructAnnotation)
			if !ok {
				continue
			}

			if !strings.Contains(comment.Text, annotation.Tag()) {
				continue
			}

			// found the annotation
			if Debug {
				log.Println("found struct annotation tag:", annotation.Tag())
			}

			parseAnnotationAttrs(comment.Text, annotation)

			var typeName string
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if typeSpec.Name != nil {
						typeName = typeSpec.Name.Name
						break
					}
				}
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
			funcAnnotation, ok := annotation.(FuncAnnotation)
			if !ok {
				continue
			}

			if !strings.Contains(comment.Text, annotation.Tag()) {
				continue
			}

			// found the annotation
			if Debug {
				log.Println("found func annotation tag:", annotation.Tag())
			}

			parseAnnotationAttrs(comment.Text, annotation)

			var funcName string
			if funcDecl.Name != nil {
				funcName = funcDecl.Name.Name
			}

			funcAnnotation.Execute(pkgName, funcName, funcDecl)
		}
	}
}

func parseAnnotationAttrs(comment string, an Annotation) {
	if Debug {
		log.Printf("%T comment: %s\n", an, comment)
	}

	kv := make(map[string]string)
	match := annotationRE.FindStringSubmatch(comment)
	if match == nil {
		return
	}

	for i, name := range annotationRE.SubexpNames() {
		// ignore the whole regexp match and unnamed groups
		if i == 0 || name == "" {
			continue
		}

		kv[name] = match[i]
	}

	if Debug {
		log.Printf("%T attrs: %+v", an, kv)
	}

	an.SetAttrs(kv)
}
