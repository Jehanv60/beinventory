package exception

type SameFound struct {
	Error string
}

func NewSameFound(error string) SameFound {
	return SameFound{
		Error: error,
	}
}
