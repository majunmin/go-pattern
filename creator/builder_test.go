// Created By: junmin.ma
// Description: <description>
// Date: 2022-01-29 06:54
package creator

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	b := NewConcreteBuilder()
	director := NewDirector(b)
	director.Construct()

	product := b.GetResults()

	fmt.Printf("%v", product)
}
