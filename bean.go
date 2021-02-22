package beans

import (
	"fmt"
	"reflect"
)

type CleanUpFunc func()

var (
	cleanUpFuncType = reflect.TypeOf(CleanUpFunc(nil))
	emptyListValue  []reflect.Value
)

type bean struct {
	dependency []reflect.Type
	initType reflect.Type
	initFun interface{}
	container *ProviderContainer
	haveError bool
	singletonObjectValue []reflect.Value
	cleanUpFunc *reflect.Value
}

func newBeanFromObjectData(initData interface{},cleanUpFunc ...CleanUpFunc) (reflect.Type,*bean,error) {
	dataValue := reflect.ValueOf(initData)
	dataType := dataValue.Type()
	beanData := &bean{
		singletonObjectValue: []reflect.Value{dataValue},
		initType: dataType,
	}
	if len(cleanUpFunc) > 0 {
		cleanUpFuncValue := reflect.ValueOf(cleanUpFunc[0])
		beanData.cleanUpFunc = &cleanUpFuncValue
	}
	return dataType,beanData,nil
}

func newBean(initFun interface{},container *ProviderContainer,singleton bool) (reflect.Type,*bean,error) {
	var (
		dependencyList []reflect.Type
		cleanUpExist bool
	)
	funType := reflect.TypeOf(initFun)
	kind := funType.Kind()
	if kind != reflect.Func {
		return nil, nil, fmt.Errorf(mustBeFunc)
	}
	numOut := funType.NumOut()
	if numOut < 1 || (numOut > 2 && !singleton) || (numOut > 3 && singleton) {
		return nil, nil, fmt.Errorf(outputRestriction)
	}
	numIn := funType.NumIn()
	for i := 0;i<numIn;i++ {
		inType := funType.In(i)
		dependencyList = append(dependencyList,inType)
	}
	returnType := funType.Out(0)
	outputBean := &bean{
		dependency: dependencyList,
		initFun: initFun,
		container: container,
		initType: returnType,
	}
	if numOut == 2 {
		if secType := funType.Out(1);secType.String() != "error" {
			if !singleton {
				return nil, nil, fmt.Errorf(outputRestriction)
			}
			if secType != cleanUpFuncType {
				return nil, nil, fmt.Errorf(outputRestriction)
			}
			cleanUpExist = true
		} else {
			outputBean.haveError = true
		}
	}

	if singleton {
		if numOut == 3 {
			secType := funType.Out(2)
			if secType != cleanUpFuncType {
				return nil, nil, fmt.Errorf(outputRestriction)
			}
			cleanUpExist = true
		}
		val,err := outputBean.call()
		if err != nil {
			return nil, nil, err
		}
		outputBean.singletonObjectValue = val[:1]
		if cleanUpExist {
			outputBean.cleanUpFunc = &(val[len(val)-1])
		}
	}

	return returnType,outputBean,nil
}

func(b *bean) cleanUp() {
	if b.cleanUpFunc != nil {
		b.cleanUpFunc.Call(emptyListValue)
	}
}

func(b *bean) call() ([]reflect.Value,error) {
	if len(b.singletonObjectValue) > 0 {
		return b.singletonObjectValue,nil
	}
	var listDependency []reflect.Value
	for _,dep := range b.dependency{
		depBean,ok := b.container.dependencyInitiator[dep]
		if !ok {
			return nil, errorDepNotFound(dep)
		}
		depVal,err := depBean.call()
		if err != nil {
			return nil, err
		}
		listDependency = append(listDependency,depVal[0])
	}
	retVal := reflect.ValueOf(b.initFun).Call(listDependency)
	if b.haveError {
		if errData := retVal[1];!retVal[1].IsNil() {
			return nil,errorDepReturnError(b.initType,errData.Interface().(error))
		}
	}
	return retVal,nil
}
