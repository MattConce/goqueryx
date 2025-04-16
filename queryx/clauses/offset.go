package clauses

type Offset struct {
	Offset any
}

func NewOffset(offset any) *Offset {
	return &Offset{Offset: offset}
}
