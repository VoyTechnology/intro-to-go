package subpack

// User is a struct containing all information about the user
type User struct {
	Username string
	password string
}

// NewUser creates an instance of a new user
func NewUser(name string, pass string) User {
	user := User{
		Username: name,
		password: pass,
	}

	return user
}

// CheckPassword gets the user password and checks is it valid
func (u User) CheckPassword(pass string) bool {
	return pass == u.password
}
