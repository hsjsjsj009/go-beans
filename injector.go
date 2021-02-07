package bean_DI

import (
	"fmt"
	"reflect"
)

func(c *ProviderContainer) getDepValue(typ reflect.Type) ([]reflect.Value,error) {
	initFunc,ok := c.dependencyInitiator[typ]
	if !ok {
		return nil,ErrorDepNotFound(typ)
	}
	initValue,err := initFunc.call()
	if err != nil {
		return nil,err
	}
	return initValue,err
}

func(c *ProviderContainer) InjectStruct(data interface{}) (err error) {
	defer func() {
		if r:=recover();r != nil {
			err = fmt.Errorf("%s",r)
		}
	}()
	if data == nil {
		return fmt.Errorf(NotNil)
	}
	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Ptr {
		return fmt.Errorf(MustBePtr)
	}
	dataValueElem := reflect.ValueOf(data).Elem()
	dataTypeElem := dataType.Elem()
	numField := dataTypeElem.NumField()
	for i := 0;i<numField;i++ {
		dataField := dataTypeElem.Field(i)
		dataFieldType := dataField.Type
		dataFieldTag := dataField.Tag
		if dataFieldTag.Get("bean") != BeanAutoWired {
			continue
		}
		initValue,err := c.getDepValue(dataFieldType)
		if err != nil {
			return err
		}
		dataValueElem.Field(i).Set(initValue[0])
	}

	return nil
}

func(c *ProviderContainer) InjectVariable(vars ...interface{}) (err error) {
	defer func() {
		if r:=recover();r != nil {
			err = fmt.Errorf("%s",r)
		}
	}()
	for _,varData := range vars {
		varType := reflect.TypeOf(varData)
		if varType.Kind() != reflect.Ptr {
			return fmt.Errorf(MustBePtr)
		}
		varElemValue := reflect.ValueOf(varData).Elem()
		initValue,err := c.getDepValue(varType.Elem())
		if err != nil {
			return err
		}
		varElemValue.Set(initValue[0])
	}
	return nil
}
