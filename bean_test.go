package bean_DI

import (
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type BeanTestSuite struct {
	suite.Suite
	Container *ProviderContainer
}

func (s *BeanTestSuite) SetupTest()  {
	s.Container = NewContainer()
	s.Container.AddProvider(newTest1RetInterface)
}

func(s *BeanTestSuite) TestNewBeanNotFunc() {
	provider := "test"
	_,_,err := newBean(provider,s.Container)
	s.EqualError(err, mustBeFunc)
}

func(s *BeanTestSuite) TestNewBeanFuncWithNoOutput() {
	provider := func() {}
	_,_,err := newBean(provider,s.Container)
	s.EqualError(err, onlyOneOutput)
}

func(s *BeanTestSuite) TestNewBeanFuncWithTwoOutput() {
	provider := func() (error,struct{}) {
		return nil,struct{}{}
	}
	_,_,err := newBean(provider,s.Container)
	s.EqualError(err, onlyOneOutput)
}

func(s *BeanTestSuite) TestBeanCall() {
	_,bean,_ := newBean(newTest2Interface,s.Container)
	val, err := bean.call()
	s.Nil(err)
	s.Equal(val[0].Type(),reflect.TypeOf(&test2{}))
}

func(s *BeanTestSuite) TestBeanCallTypeNotFound() {
	provider := func(t *test1,_ test) *test2 {
		return &test2{
			Test1: t,
		}
	}
	_,bean,_ := newBean(provider,s.Container)
	_, err := bean.call()
	s.NotNil(err)
	s.EqualError(err, errorDepNotFound(reflect.TypeOf(&test1{})).Error())
}

func TestBeanTestSuite(t *testing.T) {
	suite.Run(t,new(BeanTestSuite))
}

