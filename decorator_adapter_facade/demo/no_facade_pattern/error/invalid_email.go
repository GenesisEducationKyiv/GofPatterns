package error

type ErrInvalidEmail struct {
}

func (ErrInvalidEmail) Error() string {
	return "Invalid email"
}
