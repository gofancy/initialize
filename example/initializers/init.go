package initializers

import (
	"github.com/gofancy/initialize"
)

type initializer struct{}

func init() {
	initialize.AllFrom(initializer{})
}
