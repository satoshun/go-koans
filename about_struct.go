package koans

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

	var pv *V
	// both raise nil pointer error?
	pv.value()
	pv.pointer()
}
