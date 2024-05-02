package mytypes

type MyInt int
type MyString string

// Twice multiplies its receiver by 2 and returns
// the result.
func (i MyInt) Twice() MyInt {
	return i * 2
}

func (s MyString) MyStringLen() int {
	return len(s)
}
