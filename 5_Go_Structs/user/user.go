package user

import (
	"errors"
	"fmt"
	"time"
)

// structs that start with Uppercase can be exported outside of their packages, while lowercases can only be used locally.
// SAME THING FOR ITS FIELDS!!

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time // this variable's type is time coming from Time package.
}

type Admin struct {
	email    string
	password string
	User     // inherits User struct
}

// A METHOD = func that works only using data from a specific struct;
// The (...) after the 'func' is called a 'receiver' - in this case will access the targetted structure 'User'
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// methods that modify/alter struct values, should be passed as *ptr otherwise it will just create a copy of that struct everytime!
func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

// admin constructor with inheritance from User
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

// CONSTRUCTOR - output is a *User because when you create a new user it must point to the addr of the newly created user.
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
