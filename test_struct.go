package bean_DI

type Test interface {
	GetTest() string
}

type Test1 struct {
	Test1 string
}

type Test2 struct {
	Test1 Test
}

type Test3 struct {
	Test1 *Test1
}

func (t *Test1) GetTest() string {
	return t.Test1
}

func newTest1RetInterface() Test {
	return &Test1{Test1: "asda"}
}

func newTest2Interface(t Test) *Test2 {
	return &Test2{
		t,
	}
}

func newTest3Ptr(t *Test1) *Test3 {
	return &Test3{
		t,
	}
}
