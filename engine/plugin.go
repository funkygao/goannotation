package engine

import (
	"fmt"
	"strings"
)

type AnnotationType uint8

const (
	TAG_FUNC AnnotationType = iota + 1
	TAG_TYPE
)

var (
	registeredPlugins []Plugin = make([]Plugin, 0)
)

type Plugin interface {
	AnnotationTag() string
	AnnotationType() AnnotationType
}

func RegisterPlugin(p Plugin) {
	var tag = p.AnnotationTag()
	if !strings.HasPrefix(tag, "@") {
		panic(fmt.Sprintf("invalid plugin with tag: %s", tag))
	}

	// find dup
	for _, plugin := range registeredPlugins {
		if plugin.AnnotationTag() == tag {
			panic(fmt.Sprintf("found dup plugin tag: %s", tag))
		}
	}

	registeredPlugins = append(registeredPlugins, p)
}
