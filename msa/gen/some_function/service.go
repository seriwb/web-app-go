// Code generated by goa v3.0.4, DO NOT EDIT.
//
// some-function service
//
// Command:
// $ goa gen github.com/seriwb/front-bff-msa/msa/design

package somefunction

import (
	"context"
)

// The calc service performs operations on numbers.
type Service interface {
	// Add implements add.
	Add(context.Context, *AddPayload) (res int, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "some-function"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"add"}

// AddPayload is the payload type of the some-function service add method.
type AddPayload struct {
	// Left operand
	A int
	// Right operand
	B int
}