package main
import "fmt"
func main() {
	var ConferenceName =" go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf( "welcom to %v booking app \n", ConferenceName)
	fmt.Printf( " we have total of %v tickets and %v are still available " , conferenceTickets, remainingTickets)

}
