package domain

type (
	User struct {
		ID    string
		Name  string
		Email string
		Age   int
	}

	Opt func(*User)
)

func NewUser(opts ...Opt) *User {
	u := &User{}

	for _, opt := range opts {
		opt(u)
	}

	return u
}

func WithName(userName string) Opt {
	return func(u *User) {
		u.Name = userName
	}
}

func WithEmail(userEmail string) Opt {
	return func(u *User) {
		u.Email = userEmail
	}
}

func WithAge(userAge int) Opt {
	return func(u *User) {
		u.Age = userAge
	}
}
