package engine

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
	registeredPlugins = append(registeredPlugins, p)
}
