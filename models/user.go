package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(userObject User) (User, error) {
	if userObject.ID != 0 {
		return User{}, errors.New("user must not include ID")
	}
	userObject.ID = nextID
	nextID++
	users = append(users, &userObject)
	return userObject, nil
}

func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}

	return User{}, fmt.Errorf("User with id %v not found", id)
}

func UpdateUser(user User) (User, error) {
	for index, candidate := range users {
		if candidate.ID == user.ID {
			users[index] = &user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User with id %v not found", user.ID)
}

func DeleteUserByID(id int) error {
	for index, candidate := range users {
		if candidate.ID == id {
			users = append(users[:index], users[:index+1]...)
			return nil
		}
	}
	return fmt.Errorf("User with id %v not found", id)
}
