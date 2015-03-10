package plugins

import (
	"github.com/funkygao/goannotation/engine"
)

// @joiner
type Joiner struct {
}

func (this *Joiner) AnnotationTag() string {
	return "@joiner"
}

func (this *Joiner) AnnotationType() engine.AnnotationType {
	return engine.TAG_TYPE
}

func init() {
	engine.RegisterPlugin(&Joiner{})
}
