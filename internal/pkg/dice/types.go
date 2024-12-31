package dice

type Name string

type Size int

func (n Name) String() string {
	return string(n)
}

func (s Size) Int() int {
	return int(s)
}
