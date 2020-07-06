package constraint

import "testing"

type constraintChain struct {
}

type assertion struct {
	t *testing.T
}

func Assert(t *testing.T) *assertion {
	return &assertion{t}
}

func (a *assertion) That(obj interface{}) {

}

// func test() {
// 	Assert(t).That(1)
// }
