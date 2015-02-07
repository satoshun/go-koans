package koans

import (
	"encoding/json"
	"fmt"
	"strings"
)

// impl method
type V struct {
	i int
}

func (self V) incr() {
	self.i++
}

func (self *V) pincr() {
	if self != nil {
		self.i++
	}
}

func aboutStruct() {
	v := V{}

	// The diffrence between value receiver and pointer receiver
	assert(v.i == intintint)
	v.incr()
	assert(v.i == intintint)
	v.pincr()
	assert(v.i == intintint)

	var pv *V
	// both raise nil pointer error?
	pv.incr()
	pv.pincr()

	// composition
	type V2 struct {
		V
	}
	v2 := V2{}
	v2.pincr()
	assert(v2.i == intintint)

	// define same field
	type V3 struct {
		V
		i int
	}
	v3 := V3{i: 100}
	assert(v3.i == intintint)
	assert(v3.V.i == intintint)

	// translate json
	type J struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	j := J{FirstName: "sato", LastName: "tarou"}
	jByte, _ := json.Marshal(j)
	jString := string(jByte)
	assert(strings.Contains(jString, fmt.Sprintf(`"first_name":%q`, stringstring)))
	assert(strings.Contains(jString, fmt.Sprintf(`"last_name":%q`, stringstring)))

	var j2 J
	json.Unmarshal(jByte, &j2)
	assert(j2.FirstName == stringstring)
	assert(j2.LastName == stringstring)
}
