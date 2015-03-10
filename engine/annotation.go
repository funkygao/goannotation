package engine

import (
	"fmt"
	"go/ast"
	"strings"
)

type AnnotationType uint8

const (
	ANNOTATION_FUNC   AnnotationType = iota + 1 // applied to func declaration
	ANNOTATION_STRUCT                           // applied to struct declaration

	ANNOTATION_PREFIX = "@"
)

var (
	registeredAnnotations []Annotation = make([]Annotation, 0)

	Debug = true
)

type Annotation interface {
	Tag() string
	Type() AnnotationType
}

type StructAnnotation interface {
	Annotation
	Execute(pkgName string, typeName string)
}

type FuncAnnotation interface {
	Annotation
	Execute(pkgName string, funcName string, decl *ast.FuncDecl)
}

func RegisterPlugin(an Annotation) {
	var tag = an.Tag()
	if !strings.HasPrefix(tag, ANNOTATION_PREFIX) {
		panic(fmt.Sprintf("invalid plugin with tag: %s", tag))
	}

	// find dup
	for _, plugin := range registeredAnnotations {
		if plugin.Tag() == tag {
			panic(fmt.Sprintf("found dup plugin tag: %s", tag))
		}
	}

	registeredAnnotations = append(registeredAnnotations, an)
}
