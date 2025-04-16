package clauses

type Limit struct {
	Limit any
}

func NewLimit(limit any) *Limit {
	return &Limit{Limit: limit}
}
