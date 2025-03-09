package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact contactInfo
	contactInfo
}

func main() {
	// 1st approach
	// alex := person{"Alex", "Anderson"}
	
	// 2nd approach
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// 3rd Approach

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contact.email = "xyz@gmail.com"
	// alex.contact.zipCode = 599595
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// jim := person{
	// 	firstName: "Jim",
	// 	lastName: "Party",
	// 	contact: contactInfo{
	// 		email : "xyz@gmail.com",
	// 		zipCode : 599595,
	// 	},
	// }
	// fmt.Printf("%+v", jim)

	joe:= person{
		firstName: "Joe",
		lastName: "Biden",
		contactInfo: contactInfo{
			email : "xyz@gmail.com",
			zipCode : 599595,
		},
	}
	// joePointer := &joe
	// joePointer.updateName("Josh")
	// joePointer.print()
	joe.updateName("Josh")
	joe.print()
}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
} 

func (p person) print()  {
	fmt.Printf("%+v", p)
}