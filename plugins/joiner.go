package plugins

import (
	"github.com/funkygao/goannotation/engine"
)

// @joiner
type Joiner struct {
}

func (this *Joiner) Tag() string {
	return "@joiner"
}

func (this *Joiner) Type() engine.AnnotationType {
	return engine.ANNOTATION_TYPE
}

func init() {
	engine.RegisterPlugin(&Joiner{})
}
