package koans

type V struct {
	i int
}

func (self V) value() {
	self.i++
}

func (self *V) pointer() {
	if self != nil {
		self.i++
	}
}

func aboutStruct() {
	v := V{}

	// The diffrence between value receiver and pointer receiver
	assert(v.i == intintint)
	v.value()
	assert(v.i == intintint)
	v.pointer()
	assert(v.i == intintint)

	var pv *V
	// both raise nil pointer error?
	pv.value()
	pv.pointer()
}
