package koans

type I interface {
	Get() int
}

type I1 struct {
	v int
}

func (self I1) Get() int {
	return self.v
}

type I2 struct {
	v float32
}

func (self I2) Get() float32 {
	return self.v
}

func checkI(v interface{}) bool {
	_, ok := v.(I)
	return ok
}

func aboutInterface() {
	i1, i2 := I1{1}, I2{1.0}

	// struct I1 impl interface I?
	assert(checkI(i1) == boolboolbool)
	// struct I2 impl interface I?
	assert(checkI(i2) == boolboolbool)

	// cast struct to interface
	i := I(i1)
	assert(i.Get() == intintint)
}
