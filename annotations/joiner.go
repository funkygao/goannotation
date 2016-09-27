package annotations

import (
	"log"

	"github.com/funkygao/goannotation/engine"
)

var _ engine.StructAnnotation = &Joiner{}

// Jointer is a sample annotation implementation applied on struct.
type Joiner struct {
	attrs map[string]string
}

func (this *Joiner) Tag() string {
	return "@joiner"
}

func (this *Joiner) SetAttrs(kv map[string]string) {
	this.attrs = kv
}

func (this *Joiner) Execute(pkgName string, typeName string) {
	// we can auto generate src code here
	log.Println(pkgName, typeName)
}

func init() {
	engine.RegisterPlugin(&Joiner{})
}
