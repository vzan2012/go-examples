package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}
type logWriter struct{}

type logWriter struct{}

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
	// experiment()

	// Call Colors
	// callColors()
	englishUser := englishBot{}
	spanishUser := spanishBot{}

	printGreeting(englishUser)
	printGreeting(spanishUser)

	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Errorf("Error : %v", err)
		os.Exit(1)
	}

	// fmt.Printf("Response: %v", resp)
	// fmt.Println(resp.Request)

	// In real time, this is not possible to have a fixed length of byte array
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}
	// io.Copy(os.Stdout, resp.Body)  Working one

	io.Copy(lw, resp.Body)

}

func (p person) print() {
	fmt.Println()
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Number of count: ", len(bs))
	return len(bs), nil
}
