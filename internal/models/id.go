package models

type ID string

func (s ID) String() string {
	return string(s)
}
