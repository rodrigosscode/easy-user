package error

type UserNameEmptyErr struct{}

func NewUserNameEmptyErr() error {
	return UserNameEmptyErr{}
}

func (e UserNameEmptyErr) Error() string {
	return "name cannot be empty"
}
