package beans

import (
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type ProviderTestSuite struct {
	suite.Suite
	Container *ProviderContainer
}

func (s *ProviderTestSuite) SetupTest()  {
	s.Container = NewContainer()
}

func(s *ProviderTestSuite) TestAddProviderNotFunc() {
	provider := "test"
	s.PanicsWithValue(mustBeFunc, func() {
		s.Container.AddProvider(provider)
	})
}

func(s *ProviderTestSuite) TestAddProviderFuncWithNoOutput() {
	provider := func() {}
	s.PanicsWithValue(outputRestriction, func() {
		s.Container.AddProvider(provider)
	})
}

func(s *ProviderTestSuite) TestAddProviderFuncWithTwoOutput() {
	provider := func() (error,struct{}) {
		return nil,struct{}{}
	}
	s.PanicsWithValue(outputRestriction, func() {
		s.Container.AddProvider(provider)
	})
}

func(s *ProviderTestSuite) TestAddProviderFunc() {
	s.NotPanics(func() {
		s.Container.AddProvider(newTest1RetInterface)
	})
	s.Equal(1,len(s.Container.dependencyInitiator))
}
func(s *ProviderTestSuite) TestAddProviderFuncDuplicate() {
	s.NotPanics(func() {
		s.Container.AddProvider(newTest1RetInterface)
	})
	s.PanicsWithValue(depAlreadyDefined,func() {
		s.Container.AddProvider(newTest1RetInterface)
	})
	s.Equal(1,len(s.Container.dependencyInitiator))
}

func(s *ProviderTestSuite) TestAddProviderSingleton() {
	s.NotPanics(func() {
		s.Container.AddProviderSingleton(newTest1RetInterface)
	})
	s.Equal(1,len(s.Container.dependencyInitiator))
}

func(s *ProviderTestSuite) TestAddProviderSingletonDepNotFound() {
	s.PanicsWithValue(errorDepNotFound(reflect.TypeOf(&test1{})).Error(),func() {
		s.Container.AddProviderSingleton(newTest3Ptr)
	})
	s.Equal(0,len(s.Container.dependencyInitiator))
}

func(s *ProviderTestSuite) TestAddSingletonObjectData() {
	s.NotPanics(func() {
		s.Container.AddObjectSingleton(&test2{})
	})
	s.Equal(1,len(s.Container.dependencyInitiator))
}

func(s *ProviderTestSuite) TestCleanUpFunc()  {
	count := 0
	s.Container.AddProviderSingleton(func() (test,CleanUpFunc) {
		return &test1{"asdasdasd"}, func() {
			count++
		}
	})
	s.Container.CleanUp()
	s.Equal(1,count)
}

func TestProviderTestSuite(t *testing.T) {
	suite.Run(t, new(ProviderTestSuite))
}
