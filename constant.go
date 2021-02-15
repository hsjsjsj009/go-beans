package beans

import (
	"fmt"
	"reflect"
)

const (
	beanAutoWired = "autowired"
)

const (
	mustBeFunc        = "provider must be a function"
	outputRestriction = "provider function must have 1 output or 2 output include error in the second output"
	mustBePtr         = "must be pointer"
	notNil            = "not nil"
	depAlreadyDefined = "dependency already defined"
)

func errorDepNotFound(depType reflect.Type) error {
	return fmt.Errorf("dependency %s not found",depType.String())
}

func errorDepReturnError(depType reflect.Type,err error) error {
	return fmt.Errorf("dependency %s error : %s",depType.String(),err.Error())
}
