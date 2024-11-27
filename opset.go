package gonnx

import (
	"github.com/advancedclimatesystems/gonnx/ops"
	"github.com/advancedclimatesystems/gonnx/ops/opset13"
)

// OpGetter is a function that gets an operator based on a string.
type OpGetter func(string) (ops.Operator, error)

var operatorGetters = map[int64]OpGetter{
	13: opset13.GetOperator,
}

// ResolveOperatorGetter resolves the getter for operators based on the opset version.
func ResolveOperatorGetter(opsetID int64) (OpGetter, error) {
	if getOperator, ok := operatorGetters[opsetID]; ok {
		return getOperator, nil
	}

	// Else find the lowest supported opset that is larger than opsetID
	var nextLargerOpset int64 = -1
	for key := range operatorGetters {
		if key > opsetID && (nextLargerOpset == -1 || key < nextLargerOpset) {
			nextLargerOpset = key
		}
	}

	if nextLargerOpset != -1 {
		return operatorGetters[nextLargerOpset], nil
	}

	return nil, ops.ErrUnsupportedOpsetVersion
}
