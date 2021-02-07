package bean_DI

import (
	"fmt"
	"reflect"
)

type bean struct {
	dependency []reflect.Type
	initFun interface{}
	container *ProviderContainer
}

func newBean(initFun interface{},container *ProviderContainer) (reflect.Type,*bean,error) {
	var dependencyList []reflect.Type
	funType := reflect.TypeOf(initFun)
	kind := funType.Kind()
	if kind != reflect.Func {
		return nil, nil, fmt.Errorf(MustBeFunc)
	}
	numOut := funType.NumOut()
	if numOut != 1 {
		return nil, nil, fmt.Errorf(OnlyOneOutput)
	}
	numIn := funType.NumIn()
	for i := 0;i<numIn;i++ {
		inType := funType.In(i)
		dependencyList = append(dependencyList,inType)
	}
	returnType := funType.Out(0)
	return returnType,&bean{
		dependency: dependencyList,
		initFun: initFun,
		container: container,
	},nil
}

func(b *bean) call() ([]reflect.Value,error) {
	var listDependency []reflect.Value
	for _,dep := range b.dependency{
		depBean,ok := b.container.dependencyInitiator[dep]
		if !ok {
			return nil,ErrorDepNotFound(dep)
		}
		depVal,err := depBean.call()
		if err != nil {
			return nil, err
		}
		listDependency = append(listDependency,depVal[0])
	}
	return reflect.ValueOf(b.initFun).Call(listDependency),nil
}
