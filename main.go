package main

//1. Init project by creating the module with the cmd "go mod init booking-app"
// This will generate the go.mod file

//2. Every file must belong to a package, which should be defined in the first line

//4. Packages are simply a collection of source files, we are importing one of those, fmt, here
//5. You can run the project using the cmd "go run main.go", this compiles and runs the code
import (
	"fmt"
)

// 3. This function is the entrypoint of the project
func main() {
    var conferenceName = "Go Conference";
    const conferenceTicketCount = 50;
    var remainingTickets = 50;

	fmt.Printf("Welcome to %s booking application\n", conferenceName);
    fmt.Println("You can book your tickets here!");
    fmt.Printf("We have %v out of %v tickets remaining\n", remainingTickets, conferenceTicketCount); 

}
