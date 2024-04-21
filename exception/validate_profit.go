package exception

type NotEqual struct {
	Error string
}

func NewNotEqual(error string) NotEqual {
	return NotEqual{
		Error: error,
	}
}
