package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// 	user := person{firstName: "v", lastName: "zan2012"}
	// 	fmt.Println(user.f irstName, user.lastName)

	// var abcd person
	// abcd.firstName = "vzan2012"
	// abcd.lastName = "year"
	// fmt.Println(abcd)
	// fmt.Printf("%+v", abcd)

	user := person{
		firstName: "George",
		lastName:  "Washington",
		contactInfo: contactInfo{
			email:   "gwash@gmail.com",
			zipCode: 94400,
		},
	}

	// fmt.Println(user)
	// fmt.Printf("%+v", user)
	// Way 1:
	// userPointer := &user
	// userPointer.updateName("vzan2012")

	user.print()
	user.updateName("vzan2012")
	user.print()
	experiment()

}

func (p person) print() {
	fmt.Println()
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
