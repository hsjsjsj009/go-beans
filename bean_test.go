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
	s.EqualError(err,MustBeFunc)
}

func(s *BeanTestSuite) TestNewBeanFuncWithNoOutput() {
	provider := func() {}
	_,_,err := newBean(provider,s.Container)
	s.EqualError(err,OnlyOneOutput)
}

func(s *BeanTestSuite) TestNewBeanFuncWithTwoOutput() {
	provider := func() (error,struct{}) {
		return nil,struct{}{}
	}
	_,_,err := newBean(provider,s.Container)
	s.EqualError(err,OnlyOneOutput)
}

func(s *BeanTestSuite) TestBeanCall() {
	_,bean,_ := newBean(newTest2Interface,s.Container)
	val, err := bean.call()
	s.Nil(err)
	s.Equal(val[0].Type(),reflect.TypeOf(&Test2{}))
}

func(s *BeanTestSuite) TestBeanCallTypeNotFound() {
	provider := func(t *Test1,_ Test) *Test2 {
		return &Test2{
			Test1: t,
		}
	}
	_,bean,_ := newBean(provider,s.Container)
	_, err := bean.call()
	s.NotNil(err)
	s.EqualError(err,ErrorDepNotFound(reflect.TypeOf(&Test1{})).Error())
}

func TestBeanTestSuite(t *testing.T) {
	suite.Run(t,new(BeanTestSuite))
}

