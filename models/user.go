package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func getUsers() []*User {
	return users
}

func addUser(userObject User) (User, error) {
	userObject.ID = nextID
	nextID++
	users = append(users, &userObject)
	return userObject, nil
}
