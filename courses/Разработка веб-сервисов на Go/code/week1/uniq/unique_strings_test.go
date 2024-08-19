package uniq

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
3
4
5
`

var testBad = `1
2
5
1
9
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := Unique(in, out)

	if err != nil {
		t.Error("ERROR FAILED")
	}
	res := out.String()
	if res != testOk {
		t.Errorf("got %v, expected %v", res, testOk)
	}
}

func TestErrorUniq(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testBad))
	out := new(bytes.Buffer)
	err := Unique(in, out)

	if err == nil {
		t.Error("expected error, but it not present")
	}
}
