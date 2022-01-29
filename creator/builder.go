// Created By: junmin.ma
// Description: <description>
// Date: 2022-01-29 06:46
package creator

// director is the object which orchetrates the building of a product.
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() {
	d.builder.Build()
}

// Builder is an interface for building.
type Builder interface {
	Build()
}

// ConcreteBuilder is a builder for building a Product
type ConcreteBuilder struct {
	built bool
}

func NewConcreteBuilder() *ConcreteBuilder {
	// fill default value
	return &ConcreteBuilder{built: false}
}

func (b *ConcreteBuilder) Build() {
	b.built = true
}

func (b *ConcreteBuilder) GetResults() Product {
	return Product{built: b.built}
}

type Product struct {
	built bool
}
