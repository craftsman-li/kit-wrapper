package di

import (
	"github.com/tinrab/kit/di"
)

var c = di.New()

// inject in container
func Register(name string, dependency interface{}) {
	c.Provide(name, dependency)
}

// GetByName returns a dependency by name.
func GetByName(name string) interface{} {
	return c.GetByName(name)
}

// GetByType returns a dependency by type.
func GetByType(typ interface{}) interface{} {
	return c.GetByType(typ)
}

// Resolve decorates objects with dependencies and initializes them.
func Resolve() error {
	return c.Resolve()
}

// Close closes all dependencies.
func Close() {
	c.Close()
}
