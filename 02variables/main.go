package main

import "fmt"

const LoginToken string = "ghbhjhsb" // Public

func main() {
	var username string = "Pranav"
	fmt.Println("Pranav")
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.455544555151525525
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default vakues and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var abc string
	fmt.Println(abc)
	fmt.Printf("Variable is of type: %T \n", abc)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	numberOfUsers := 20.90
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
