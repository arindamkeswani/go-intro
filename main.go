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
    conferenceName := "Go Conference"; // the := syntactical sugar applies only to var, not const, and you cannot use it while defining a type
    const conferenceTicketCount = 50;
    var remainingTickets uint = 50;

	fmt.Printf("Welcome to %s booking application\n", conferenceName);
    fmt.Println("You can book your tickets here!");
    fmt.Printf("We have %v out of %v tickets remaining\n", remainingTickets, conferenceTicketCount); 

    // Data Types can either be inferred directly if they are assigned, or they can be explicitly defined
    var userName string;
    var userRequestedTicketCount int;
    // Ask user for their name

    userName = "Arindam"
    userRequestedTicketCount = 0;

    fmt.Printf("User %v booked %v tickets.\n",userName, userRequestedTicketCount);
    fmt.Printf("conferenceTicketCount is of the Type: %T, conferenceName is of the Type: %T",conferenceTicketCount, conferenceName);
}
