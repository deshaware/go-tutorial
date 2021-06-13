package main

import "fmt"

// type person struct {
// 	firstName string
// 	lastName  string
// }

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// the below is correct syntax but it sucks big time
	// alex := person{"Alex", "Anderson"}
	// fmt.Println(alex)

	// not a bad syntax, better than the first one
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// lastly
	// var alex person
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	// fmt.Println()
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// we can default value here while declaring struct person

	// we create a new struct contactInfo

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@jhavada.com",
			zipCode: 13121,
		},
	}

	jim.print()

	
	jim.print()
	fmt.Println()

	// we need to pass POINTERS
	// WHY POINTERS NEEDED, cuz the fucking Go use pass by value, so it will create a copy of the varible, either
	// either used with receiver or passed in parameter, it will be a copy and it will work on that fucking copy
	// thus it wont modify here, thus we need pointers

	jimPointer := &jim 
	// fmt.Println(jimPointer)
	jimPointer.updateName("Jimmy")
	jim.print()
	fmt.Println()

	// & is the operator followed by a variable name
	// & it gives the memory address of the value of this variable is pointing at
	// jimPointer := &jim , jimPointer will contain RAM addess of the value at the location of the jim
	// jimPointer contains 0x1212312121AB12
	// which means jimPointer will contain address of the jim,

	// '*' gives the VALUE at that ADDRESS
	// means * will give the value at the address 0x1212312121AB12


	// let' look at the gotchas

	jim.updateName("Swap")
	jim.print()

	// Go allows to take shortcut, it says if you have a variable that is type of person, but you 
	// have a receiver is pointer to person, THAT IS TOTALLY FINE, we will glass over the fact for you
	// So we not have to cox the memory address, Go will take care of it


	// We don't need to creater a function with pointers if we work with built-in datatype, idk why
	
	// func main() {
	// 	myslice := []string{"Hi", "There", "How", "Are", "You"}
	// 	updateSlice(myslice)
	// 	fmt.Println(myslice)
	// }

	// func updateSlice(s []string){
	// 	s[0] = "Bye"
	// }

	// Output is [Bye There How Are You]

	name := "s"
	fmt.Println(*&name)


	// REFERENCE TYPE AND VALUE TYPE



}

// Below function is too important
// the * at *person shouldn't be actually thought as operator
// when we see a * in front of an actual TYPE, that is TYPE DESCRIPTION 
// which means this receiver written below can only be called with a pointer to a person

// So *person -> This is a type description, it means we are working with a pointer to a person ( which of type struct)
// , *pointerToPerson -> This is an operator - it means we want to manipulate the value the pointer is referencing

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
