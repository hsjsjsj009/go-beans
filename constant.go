package bean_DI

import (
	"fmt"
	"reflect"
)

const (
	beanAutoWired = "autowired"
)

const (
	mustBeFunc        = "provider must be a function"
	onlyOneOutput     = "provider function must have 1 output"
	mustBePtr         = "must be pointer"
	notNil            = "not nil"
	depAlreadyDefined = "dependency already defined"
)

func errorDepNotFound(depType reflect.Type) error {
	return fmt.Errorf("dependency %s not found",depType.String())
}
