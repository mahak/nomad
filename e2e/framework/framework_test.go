package framework

import (
	"fmt"
	"testing"
)

type myCase struct {
	TC
	foo string
}

func TestFramework(t *testing.T) {

	c := &myCase{
		TestCase: &internalTestCase{},
		foo:      "bar",
	}

	fmt.Println(c.Steps())
}
