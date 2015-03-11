package annotations

import (
	"github.com/funkygao/goannotation/engine"
	"log"
)

// @joiner(a=b,c=d)
type Joiner struct {
}

// @splitter
func (this *Joiner) Type() engine.AnnotationType {
	return engine.ANNOTATION_STRUCT
}
