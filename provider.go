package beans

import "reflect"

func(c *ProviderContainer) AddProvider(fun interface{}) {
	c.createBean(fun,false,false)
}

func(c *ProviderContainer) createBean(data interface{},singleton,obj bool,cleanUpFunc ...CleanUpFunc) {
	var (
		returnType reflect.Type
		beanData *bean
		err error
	)
	if obj {
		returnType,beanData,err = newBeanFromObjectData(data,cleanUpFunc...)
	}else {
		returnType,beanData,err = newBean(data,c,singleton)
	}
	if err != nil {
		panic(err.Error())
	}
	_,ok := c.dependencyInitiator[returnType]
	if ok {
		panic(depAlreadyDefined)
	}
	c.dependencyInitiator[returnType] = beanData
}

func(c *ProviderContainer) AddProviderSingleton(fun interface{}) {
	c.createBean(fun,true,false)
}

func(c *ProviderContainer) AddObjectSingleton(obj interface{},cleanUpFunc ...CleanUpFunc) {
	c.createBean(obj,true,true,cleanUpFunc...)
}

func(c *ProviderContainer) CleanUp() {
	for _,v := range c.dependencyInitiator {
		v.cleanUp()
	}
}


