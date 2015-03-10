package annotations

import (
	"github.com/funkygao/goannotation/engine"
	"log"
)

// Jointer is a sample annotation implementation applied on struct.
type Joiner struct {
}

func (this *Joiner) Tag() string {
	return "@joiner"
}

func (this *Joiner) Execute(pkgName string, typeName string) {
	// we can auto generate src code here
	log.Println(pkgName, typeName)
}

func init() {
	engine.RegisterPlugin(&Joiner{})
}
