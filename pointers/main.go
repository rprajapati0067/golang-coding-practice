package main

import (
	"errors"
	"fmt"
)

type User struct {
	Status string
}

func (u *User) UpdateStatus(status string) {
	u.Status = status
}
func updateUser(user *User, status string) {
	user.Status = status
}
func getUser(userId int64) (*User, error) {
	if userId <= 0 {
		return nil, errors.New("user not found")
	}
	user := User{}
	return &user, nil
}

func main() {
	user := User{Status: "active"}

	fmt.Println(user.Status)

	//user.UpdateStatus("inactive")
	updateUser(&user, "inactive")
	fmt.Println(user.Status)

	currentUser, userErrr := getUser(1)
	if userErrr != nil {
		return
	}
	fmt.Println(currentUser)

	/*var name = "ravi"
	var namePointer = &name // this will give memory address of name variable

	fmt.Println(*namePointer) // give me the value this pointer is pointing to(dereference operator gives actual value)
	fmt.Println(&name)*/

}
