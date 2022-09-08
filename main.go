// Run a function that takes two arguments
package main

import "fmt"

func add(happy string, coding string) string {
	return happy + coding
}
func main() {
	fmt.Println(add("I love", " elote"))
}
