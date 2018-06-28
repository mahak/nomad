package example

import (
	"github.com/hashicorp/nomad/e2e/framework"
)

type SimpleExampleTestCase struct {
	framework.TC
}

func (tc *SimpleExampleTestCase) TestExample() {
	jobs, _, err := tc.Nomad().Jobs().List(nil)
	tc.NoError(err)
	tc.Empty(jobs)
}
