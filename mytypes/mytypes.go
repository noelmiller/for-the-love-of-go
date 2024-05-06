package mytypes

import (
	"errors"
	"strings"
)

type MyInt int
type MyString string
type MyBuilder struct {
	Contents strings.Builder
	hidden   bool
}

func NewMyBuilder() (*MyBuilder, error) {
	e := MyBuilder{hidden: true}
	return &e, nil
}

func Thing() (string, error) {
	nb, err := NewMyBuilder()
	nb.Contents.WriteString("Hello World!")
	if err != nil {
		return "", errors.New("bad...")
	}
	if nb.IsVisible() == true {
		return nb.Contents.String(), nil
	} else {
		return "", errors.New("this is also bad")
	}
}

func (mb *MyBuilder) IsVisible() bool {
	return mb.hidden
}

type StringUppercaser struct {
	Contents strings.Builder
}

// Twice multiplies its receiver by 2 and returns
// the result.
func (i MyInt) Twice() MyInt {
	return i * 2
}

func (s MyString) MyStringLen() int {
	return len(s)
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}
func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}

func Double(input *int) {
	*input *= 2
}

func (input *MyInt) Double() {
	*input *= 2
}
