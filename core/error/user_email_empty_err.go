package error

type UserEmailEmptyErr struct{}

func NewUserEmailEmptyErr() error {
	return UserEmailEmptyErr{}
}

func (e UserEmailEmptyErr) Error() string {
	return "email cannot be empty"
}
