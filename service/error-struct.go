package service

type UuidError struct {
	Message string
}

func (e UuidError) Error() string {
	return e.Message
}
