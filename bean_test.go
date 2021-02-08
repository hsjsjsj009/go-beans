package beans

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
	_,_,err := newBean(provider,s.Container,false)
	s.EqualError(err, mustBeFunc)
}

func(s *BeanTestSuite) TestNewBeanFuncWithNoOutput() {
	provider := func() {}
	_,_,err := newBean(provider,s.Container,false)
	s.EqualError(err, outputRestriction)
}

func(s *BeanTestSuite) TestNewBeanFuncWithTwoOutput() {
	provider := func() (error,struct{}) {
		return nil,struct{}{}
	}
	_,_,err := newBean(provider,s.Container,false)
	s.EqualError(err, outputRestriction)
}

func(s *BeanTestSuite) TestBeanCall() {
	_,bean,_ := newBean(newTest2Interface,s.Container,false)
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
	_,bean,_ := newBean(provider,s.Container,false)
	_, err := bean.call()
	s.NotNil(err)
	s.EqualError(err, errorDepNotFound(reflect.TypeOf(&test1{})).Error())
}

func(s *BeanTestSuite) TestBeanSingleton() {
	_,bean,_ := newBean(newTest2Interface,s.Container,true)
	val1,err1 := bean.call()
	val2,err2 := bean.call()
	s.Nil(err1)
	s.Nil(err2)
	s.Equal(val1,val2)
	s.Equal(val1[0],val2[0])
	s.Equal(1,len(val1))
	s.Equal(1,len(val2))
}

func TestBeanTestSuite(t *testing.T) {
	suite.Run(t,new(BeanTestSuite))
}

