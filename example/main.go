package main

import (
	"fmt"
	"github.com/hsjsjsj009/go-beans"
)

func main() {
	provider := beans.NewContainer()
	provider.AddProvider(newTest1RetInterface)
	provider.AddProviderSingleton(newTest2Interface)
	provider.AddProvider(newTest3Ptr)

	var (
		test3 *test3
	)
	err := provider.InjectVariable(&test3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test3)
	fmt.Println(test3.Test2)
	fmt.Println(test3.Test1)

	demoStruct := struct {
		Test2 *test2 `bean:"autowired"`
	}{}

	err = provider.InjectStruct(&demoStruct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(demoStruct.Test2)
	fmt.Println(demoStruct.Test2.Test1)

	fmt.Printf("struct Test2 %p - var Test2 %p - equal %v\n",demoStruct.Test2,test3.Test2,demoStruct.Test2 == test3.Test2)

}