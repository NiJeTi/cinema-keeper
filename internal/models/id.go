package models

type (
	ID     string
	IMDBID string
)

func (id ID) String() string {
	return string(id)
}

func (id IMDBID) String() string {
	return string(id)
}
