package bean_DI

func(c *ProviderContainer) AddProvider(fun interface{}) {
	returnType,beanData,err := newBean(fun,c)
	if err != nil {
		panic(err.Error())
	}
	_,ok := c.dependencyInitiator[returnType]
	if ok {
		panic(depAlreadyDefined)
	}
	c.dependencyInitiator[returnType] = beanData
}
