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
	outputRestriction = `provider function pattern:
	a. 1 dependency output 
	b. 2 dependency output + error
	c. 2 dependency output + cleanUp Function (singleton only)
	d. 3 dependency output + error + cleanUp Function (singleton only)

cleanUp Func -> 0 input 0 output`
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
