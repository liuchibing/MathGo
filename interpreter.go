package mathgo

import (
	"bufio"
	"io"
	"strings"
	"errors"
)

type Interpreter struct {
	output *io.Writer
	vars map[string]variable
}

func NewInterpreter(outStream *io.Writer) *Interpreter {
	i := new(Interpreter)
	i.vars = make(map[string]variable)

	return i
}

func (i *Interpreter) Run(input *io.Reader) (obj, error) {
	
}
