package mathgo

import (
	"errors"
	"io"
	"bufio"
)

type tokenizer struct {
	input *bufio.Reader
}

func newTokenizer(reader *io.Reader) *tokenizer {
	t := new(tokenizer)
	t.input = bufio.NewReader(reader)

	return t
}

func (t *tokenizer) getNext() (string, error) {
	var breakchar rune = nil
	for {
		ch, err := t.input.ReadRune()
}
