package clauses

type Values struct {
	Args []any
}

type MultiValues struct {
	Args [][]any
}

func NewValues(args []any) *Values {
	return &Values{Args: args}
}

func NewMultiValues(args [][]any) *MultiValues {
	return &MultiValues{Args: args}
}
