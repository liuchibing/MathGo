package mathgo

import (
	"strconv"
)

type expr interface {
	eval() (variable, err)
}

type numExpr struct {
	value string
}

func (e *numExpr) eval() (variable, error) {

}
