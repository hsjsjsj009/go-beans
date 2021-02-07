package bean_DI

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
	Test1 *test1
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

func newTest3Ptr(t *test1) *test3 {
	return &test3{
		t,
	}
}