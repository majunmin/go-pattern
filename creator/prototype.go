// Created By: junmin.ma
// Description: <description>
// Date: 2022-01-29 22:53
package creator

// Prototype interface
type Prototype interface {
	Name() string
	Clone() Prototype
}

// ConcretePrototype instance of Prototype
type ConcretePrototype struct {
	name string
}

func (c *ConcretePrototype) Name() string {
	return c.name
}

// Clone create a cloned new instance of a concretePrototype1
func (c *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{name: c.Name()}
}
