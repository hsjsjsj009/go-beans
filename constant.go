package bean_DI

import (
	"fmt"
	"reflect"
)

const (
	BeanAutoWired = "autowired"
)

const (
	MustBeFunc = "provider must be a function"
	OnlyOneOutput = "provider function must have 1 output"
	MustBePtr = "must be pointer"
	NotNil = "not nil"
	DepAlreadyDefined = "dependency already defined"
)

func ErrorDepNotFound(depType reflect.Type) error {
	return fmt.Errorf("dependency %s not found",depType.String())
}
