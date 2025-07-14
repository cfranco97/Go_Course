package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birdate (DD/MM/YYYY):")

	// you're not forced to fill all values, if you decide not to fill a certain variable, it will simply come out as null.

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	// check for if errors and stop the app
	if err != nil {
		fmt.Println(err)
		return
	}

	// can also do it like this, but you'll need to make sure first that its in the same order as the struct otherwise it will asign wrong key:value.
	/* appUser = user{
		userFirstName,
		userLastName,
		userBirthDate,
		time.Now(),
	} */

	// null value example
	/* appUser =user{} */

	admin := user.NewAdmin("admin@mail.com", "password")

	admin.OutputUserDetails()
	admin.ClearUsername()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUsername()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
