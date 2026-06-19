package vxlstruct

import "fmt"

type Human struct {
	Gender string
	Age    int
}

func (h Human) ToString() string {
	return fmt.Sprintf("gender: %s; age: %d;", h.Gender, h.Age)
}
