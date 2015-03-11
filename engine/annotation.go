package engine

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"
)

const (
	ANNOTATION_PREFIX = "@"
)

var (
	registeredAnnotations []Annotation = make([]Annotation, 0)
	annotationRE                       = regexp.MustCompile(``)
	Debug                              = true
)

type Annotation interface {
	Tag() string
	SetAttrs(kv map[string]string)
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
