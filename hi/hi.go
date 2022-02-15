package hi

import "fmt"

func SayHello(firstName string, lastName string) string {
	return fmt.Sprintf("Hello %s %s", firstName, lastName)
}
