package bean_DI

import (
	"github.com/stretchr/testify/suite"
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
	s.PanicsWithValue(onlyOneOutput, func() {
		s.Container.AddProvider(provider)
	})
}

func(s *ProviderTestSuite) TestAddProviderFuncWithTwoOutput() {
	provider := func() (error,struct{}) {
		return nil,struct{}{}
	}
	s.PanicsWithValue(onlyOneOutput, func() {
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

func TestProviderTestSuite(t *testing.T) {
	suite.Run(t, new(ProviderTestSuite))
}
