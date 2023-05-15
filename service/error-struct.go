package service

type UuidError struct {
	message string
}

func (e UuidError) Error() string {
	return e.message
}
