package error

import "fmt"

type InvalidPageRequestErr struct {
	page  string
	limit string
}

func NewInvalidPageRequestErr(page string, limit string) error {
	return InvalidPageRequestErr{page: page, limit: limit}
}

func (e InvalidPageRequestErr) Error() string {
	return fmt.Sprintf("invalid pagination parameters page=%v and limit=%v", e.page, e.limit)
}
