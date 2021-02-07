package bean_DI

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
	s.Container.AddProvider(newTest2Interface)
	s.Container.AddProvider(newTest3Ptr)
}

func(s *InjectorTestSuite) TestInjectStruct() {
	dummyStruct := struct {
		Test Test `bean:"autowired"`
		Test2 *Test2 `bean:"autowired"`
	}{}

	err := s.Container.InjectStruct(&dummyStruct)
	s.Nil(err)
	s.NotNil(dummyStruct.Test)
	s.NotNil(dummyStruct.Test2)
}

func(s *InjectorTestSuite) TestInjectVars() {
	var (
		test Test
		test2 *Test2
	)

	err := s.Container.InjectVariable(&test,&test2)
	s.Nil(err)
	s.NotNil(test)
	s.NotNil(test2)
}

func(s *InjectorTestSuite) TestInjectStructWithNoAutoWiredField() {
	dummyStruct := struct {
		Test Test
		Test2 *Test2 `bean:"autowired"`
	}{}

	err := s.Container.InjectStruct(&dummyStruct)
	s.Nil(err)
	s.Nil(dummyStruct.Test)
	s.NotNil(dummyStruct.Test2)
}

func(s *InjectorTestSuite) TestInjectStructDepNotFoundInStructField() {
	dummyStruct := struct {
		Test1 *Test1 `bean:"autowired"`
	}{}

	err := s.Container.InjectStruct(&dummyStruct)
	s.NotNil(err)
	s.EqualError(err,ErrorDepNotFound(reflect.TypeOf((*Test1)(nil))).Error())
	s.Nil(dummyStruct.Test1)
}

func(s *InjectorTestSuite) TestInjectStructWithAValueInside() {
	dummyStruct := struct {
		Test Test
		Test2 *Test2 `bean:"autowired"`
	}{
		Test: &Test1{
			"test1",
		},
	}

	err := s.Container.InjectStruct(&dummyStruct)
	s.Nil(err)
	s.NotNil(dummyStruct.Test)
	s.NotNil(dummyStruct.Test2)
	s.Equal("test1",dummyStruct.Test.GetTest())
}

func TestInjectorTestSuite(t *testing.T) {
	suite.Run(t, new(InjectorTestSuite))
}
