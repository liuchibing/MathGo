package mathgo

import (
	"bufio"
	"io"
)

type tokenizer struct {
	input    *bufio.Reader
	readMode int
}

const (
	normal_readmode      = 0
	string_readmode      = 1
	escape_char_readmode = 2
)

func newTokenizer(reader *io.Reader) *tokenizer {
	t := new(tokenizer)
	t.input = bufio.NewReader(reader)

	return t
}

func (t *tokenizer) getNext() (string, error) {
	var buf string
	for {
		ch, _, err := t.input.ReadRune()
		if err != nil {
			return nil, err
		}
		switch ch {
		case '+', '-':
			if buf != "" {
				t.input.UnreadRune()
				return buf, nil
			} else {
				return string(ch), nil
			}
		case '"':
			switch t.readMode {
			case escape_char_readmode:
				buf += string(ch)
			case string_readmode:
				buf += string(ch)
				t.readMode = normal_readmode
				return buf, nil
			case normal_readmode:
				t.readMode = string_readmode
				buf += string(ch)
			}
		}
	}

}
