package engine

import (
	"fmt"
	"strings"
)

type AnnotationType uint8

const (
	ANNOTATION_FUNC AnnotationType = iota + 1
	ANNOTATION_TYPE
)

var (
	registeredAnnotations []Annotation = make([]Annotation, 0)
)

type Annotation interface {
	Tag() string
	Type() AnnotationType
}

func RegisterPlugin(a Annotation) {
	var tag = a.Tag()
	if !strings.HasPrefix(tag, "@") {
		panic(fmt.Sprintf("invalid plugin with tag: %s", tag))
	}

	// find dup
	for _, plugin := range registeredAnnotations {
		if plugin.Tag() == tag {
			panic(fmt.Sprintf("found dup plugin tag: %s", tag))
		}
	}

	registeredAnnotations = append(registeredAnnotations, a)
}
