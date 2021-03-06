# Go-Beans
![Go](https://github.com/hsjsjsj009/go-beans/workflows/Go/badge.svg) ![](https://badgen.net/codecov/c/github/hsjsjsj009/go-beans)
<br>
This repository contains a Golang Dependency Injector that inspired by Spring Boot Bean System

# Features

1. Autowire all dependency that you need
2. Struct Injector
3. Variable Injector
4. Can add Singleton provider
5. Can provide clean up function for singleton provider
6. Can register object to the container with the clean up function

# How to Add to Your Project

```bash
go get -u github.com/hsjsjsj009/go-beans
```

# How to Use in Your Project

## Provider Container

In this dependency injector, I use a container concept as a library for any dependencies that you need in your application. First, create your container using this function

```go
provider := beans.NewContainer()
```

## Provider Functions

After you create the provider container, you can register the provider function that you have on your project, provider function is a function that create an object that you need, a provider function only have one return output only. For the example, you can see the example below.

```go
type test interface {
	getTest() string
}

type test1 struct {
	Test1 string
}

type test2 struct {
	Test1 test
}

type test3 struct {
	Test1 test
	Test2 *test2
}

func (t *test1) getTest() string {
	return t.Test1
}

func newTest1RetInterface() test {
	return &test1{Test1: "asda"}
}

func newTest2Interface(t test) *test2 {
	return &test2{
		t,
	}
}

func newTest3Ptr(t test,t2 *test2) *test3 {
	return &test3{
		t,
		t2,
	}
}
```

After that you must register the provider function to the container.

```go
provider.AddProvider(newTest1RetInterface)
provider.AddProvider(newTest3Ptr)
```

You can add error return value for self-defined provider function error by giving second return value as error interface
```go
func newTest3PtrWithErr(t test,t2 *test2) (*test3,error) {
	if t == nil {
	    return nil,fmt.Errorf("error")
}
	return &test3{
		t,
		t2,
	},nil
}

provider.AddProvider(newTest3PtrWithErr)
```

For singleton provider, you can use `AddProviderSingleton` function, the provider function will be executed eagerly to prevent the problem in asynchronous access to the container. You can give a cleanup function in the last return value of the provider function.
```go
provider.AddProviderSingleton(newTest2Interface)
provider.AddProviderSingleton(func() (*sql.DB,CleanUpFunc){
	db := &sql.DB{}
	return db,func() {
		db.Close()
    }
})
//OR with error value
provider.AddProviderSingleton(func() (*sql.DB,err,CleanUpFunc){
        db := &sql.DB{}
        return db,nil,func() {
        db.Close()
    }
})
```

Or you can directly register singleton object to the container and also you can give a cleanup function.
```go
db := &sql.DB{}
provider.AddObjectSingleton(db,func() {
	db.Close()
})
```


## Inject the Dependency

In this module, I created two functions that inject dependencies to a struct or variable. If you want to inject dependencies to a struct you must give `bean:"autowired"` tag to your field. All field must be exported to allow injection from outside. After that you call `InjectStruct` function and give the function your struct pointer. You can see the example below.

```go
demoStruct := struct {
		Test2 *test2 `bean:"autowired"`
	}{}

err = provider.InjectStruct(&demoStruct)
```

If you use struct composition you must provide the composition provider because in go the composition works as a field without name and also the composition must be exported too.
```go
demoStruct := struct {
		*Test2 `bean:"autowired"`
	}{}

err = provider.InjectStruct(&demoStruct)
```

For injection to a variable, you just create a variable and then give the variable pointer to `InjectVariable` function. You can see the example below.

```go
var (
    test3 *test3
)
err := provider.InjectVariable(&test3)
```

## Dependency Finder

This package is depend to provider function return type because the return type of the function is used as a key to instantiate an object. All provider function parameter will be provided automatically by the container as long as the provider function of the parameter type registered to the container.



