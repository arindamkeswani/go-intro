package main

//1. Init project by creating the module with the cmd "go mod init booking-app"
// This will generate the go.mod file

//2. Every file must belong to a package, which should be defined in the first line

//4. Packages are simply a collection of source files, we are importing one of those, fmt, here
//5. You can run the project using the cmd "go run main.go", this compiles and runs the code
import (
	"fmt";
    "strings";
)

// 3. This function is the entrypoint of the project
func main() {
	conferenceName := "Go Conference" // the := syntactical sugar applies only to var, not const, and you cannot use it while defining a type
	const conferenceTicketCount = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Println("You can book your tickets here!")
	fmt.Printf("We have %v out of %v tickets remaining\n", remainingTickets, conferenceTicketCount)
    var bookings []string // This is a slice

	for remainingTickets > 0 && len(bookings) < 50{

		// var bookings [50]string; //This is an array declaration
		// Data Types can either be inferred directly if they are assigned, or they can be explicitly defined
		var firstName string
		var lastName string
		var email string
		var userRequestedTicketCount uint

		fmt.Println("Please enter your first name: ")
		// Ask user for their name
		// The following will take the user name as an input save it in userName variable
		// using the pointer, Scan can assign the value of userName as it now has the memory address
		// Can be visualised like:
		// Scan(userName) -> Scan("") -> Passing the value [NO]
		// Scan(&userName) -> Scan(&userName) -> Scan(0xc00124) -> Passing the reference [YES]
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email ID: ")
		fmt.Scan(&email)

		fmt.Println("Please enter the number of tickets that you would like to book: ")
		fmt.Scan(&userRequestedTicketCount)

        if userRequestedTicketCount > remainingTickets {
            fmt.Printf("We only have %v tickets remaining!\n", remainingTickets);
            continue;
        } else if userRequestedTicketCount == remainingTickets {
            fmt.Printf("Those were the last of our tickets, see you at the conference!");
        }

		remainingTickets = remainingTickets - userRequestedTicketCount
		// bookings[0] = firstName + " " + lastName; //Syntax for assigning an array index
		bookings = append(bookings, firstName+" "+lastName) // Syntax for appending an elem in the last place in a slice

		// fmt.Printf("The whole slice: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Slice type: %T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))

        firstNames := []string{};
    
        for _, booking := range bookings { //_ is used to define a variable that we wish to ignore
            var names = strings.Fields(booking); //splits the str passed in the param using a whitespace
            var firstName = names[0];
            firstNames = append(firstNames, firstName);
        }

		// fmt.Printf("These are our bookings: %v\n", bookings)
        fmt.Printf("Bookings (first names) are: %v\n",firstNames);

		// fmt.Printf("Thank you for booking %v ticket(s), %v %v. You will receive a confirmation mail at %v\n", userRequestedTicketCount, firstName, lastName, email)
		// fmt.Printf("conferenceTicketCount is of the Type: %T, conferenceName is of the Type: %T",conferenceTicketCount, conferenceName);
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
       
        var noTicketsRemaining bool = remainingTickets <= 0;
        if  noTicketsRemaining {
            //end program
            fmt.Println("Sold Out! Come back next year!\n");
            break;
        }
        

	}
}
