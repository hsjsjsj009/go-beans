package beans

func(c *ProviderContainer) AddProvider(fun interface{}) {
	c.createBean(fun,false)
}

func(c *ProviderContainer) createBean(fun interface{},singleton bool) {
	returnType,beanData,err := newBean(fun,c,singleton)
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
	c.createBean(fun,true)
}


