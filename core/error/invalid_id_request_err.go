package error

import "fmt"

type InvalidIdRequestErr struct {
	id int
}

func NewInvalidIdRequestErr(id int) error {
	return InvalidIdRequestErr{id: id}
}

func (e InvalidIdRequestErr) Error() string {
	return fmt.Sprintf("invalid or empty id parameter, id = %v", e.id)
}
