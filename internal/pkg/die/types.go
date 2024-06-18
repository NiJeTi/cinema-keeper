package die

type Name string

func (n Name) String() string {
	return string(n)
}

type Size int

func (s Size) Int() int {
	return int(s)
}
