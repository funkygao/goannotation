package engine

type AnnotationType uint8

const (
	TAG_FUNC AnnotationType = iota + 1
	TAG_TYPE
)

var (
	Plugins []Plugin = make([]Plugin, 0)
)

type Plugin interface {
	AnnotationTag() string
	AnnotationType() AnnotationType
}

func RegisterPlugin(p Plugin) {
	Plugins = append(Plugins, p)
}
