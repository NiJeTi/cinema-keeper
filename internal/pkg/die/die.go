package die

const (
	D2   Size = 2
	D4   Size = 4
	D6   Size = 6
	D8   Size = 8
	D10  Size = 10
	D12  Size = 12
	D20  Size = 20
	D100 Size = 100
)

const (
	D2Name   = "D2"
	D4Name   = "D4"
	D6Name   = "D6"
	D8Name   = "D8"
	D10Name  = "D10"
	D12Name  = "D12"
	D20Name  = "D20"
	D100Name = "D100"
)

func Sizes() map[Name]Size {
	return map[Name]Size{
		D2Name:   D2,
		D4Name:   D4,
		D6Name:   D6,
		D8Name:   D8,
		D10Name:  D10,
		D12Name:  D12,
		D20Name:  D20,
		D100Name: D100,
	}
}
