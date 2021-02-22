package beans

import (
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type InjectorTestSuite struct {
	suite.Suite
	Container *ProviderContainer
}

func (s *InjectorTestSuite) SetupTest() {
	s.Container = NewContainer()
	s.Container.AddProvider(newTest1RetInterface)
	s.Container.AddProviderSingleton(newTest2Interface)
	s.Container.AddProvider(newTest3Ptr)
}

func(s *InjectorTestSuite) TestInjectStruct() {
	dummyStruct := struct {
		Test  test   `bean:"autowired"`
		Test2 *test2 `bean:"autowired"`
	}{}

	err := s.Container.InjectStruct(&dummyStruct)
	s.Nil(err)
	s.NotNil(dummyStruct.Test)
	s.NotNil(dummyStruct.Test2)
}

func(s *InjectorTestSuite) TestInjectVars() {
	var (
		test  test
		test2 *test2
	)

	err := s.Container.InjectVariable(&test,&test2)
	s.Nil(err)
	s.NotNil(test)
	s.NotNil(test2)
}

func(s *InjectorTestSuite) TestInjectVarsDepNotFound() {
	var (
		test21 *test1
	)

	err := s.Container.InjectVariable(&test21)
	s.NotNil(err)
	s.Nil(test21)
}

func(s *InjectorTestSuite) TestInjectVarsNotPtr() {
	var (
		test21 test1
	)

	err := s.Container.InjectVariable(test21)
	s.NotNil(err)
}

func(s *InjectorTestSuite) TestInjectVarsSingleton() {
	var (
		test21 *test2
		test2 *test2
	)

	err := s.Container.InjectVariable(&test21,&test2)
	s.Nil(err)
	s.NotNil(test21)
	s.NotNil(test2)
	s.Equal(test21,test2)
}

func(s *InjectorTestSuite) TestInjectStructWithNoAutoWiredField() {
	dummyStruct := struct {
		Test  test
		Test2 *test2 `bean:"autowired"`
	}{}

	err := s.Container.InjectStruct(&dummyStruct)
	s.Nil(err)
	s.Nil(dummyStruct.Test)
	s.NotNil(dummyStruct.Test2)
}

func(s *InjectorTestSuite) TestInjectStructNil() {
	err := s.Container.InjectStruct(nil)
	s.NotNil(err)
	s.EqualError(err,notNil)
}

func(s *InjectorTestSuite) TestInjectStructNotAPointer() {
	dummyStruct := struct {
		Test  test
		Test2 *test2 `bean:"autowired"`
	}{}

	err := s.Container.InjectStruct(dummyStruct)
	s.NotNil(err)
	s.EqualError(err,mustBePtr)
}

func(s *InjectorTestSuite) TestInjectStructDepNotFoundInStructField() {
	dummyStruct := struct {
		Test1 *test1 `bean:"autowired"`
	}{}

	err := s.Container.InjectStruct(&dummyStruct)
	s.NotNil(err)
	s.EqualError(err, errorDepNotFound(reflect.TypeOf((*test1)(nil))).Error())
	s.Nil(dummyStruct.Test1)
}

func(s *InjectorTestSuite) TestInjectStructWithAValueInside() {
	dummyStruct := struct {
		Test  test
		Test2 *test2 `bean:"autowired"`
	}{
		Test: &test1{
			"test1",
		},
	}

	err := s.Container.InjectStruct(&dummyStruct)
	s.Nil(err)
	s.NotNil(dummyStruct.Test)
	s.NotNil(dummyStruct.Test2)
	s.Equal("test1",dummyStruct.Test.getTest())
}

func(s *InjectorTestSuite) TestInjectStructSingleton() {
	dummyStruct1 := struct {
		Test2 *test2 `bean:"autowired"`
	}{}
	dummyStruct2 := struct {
		Test2 *test2 `bean:"autowired"`
	}{}



	err := s.Container.InjectStruct(&dummyStruct1)
	s.Nil(err)
	s.NotNil(dummyStruct1.Test2)
	err1 := s.Container.InjectStruct(&dummyStruct2)
	s.Nil(err1)
	s.NotNil(dummyStruct2.Test2)
	s.Equal(dummyStruct2.Test2,dummyStruct1.Test2)
}

func TestInjectorTestSuite(t *testing.T) {
	suite.Run(t, new(InjectorTestSuite))
}
