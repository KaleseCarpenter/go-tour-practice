// Run a function that takes two arguments
package main

import "fmt"

/*func add(happy string, coding string) string {
	return happy + coding
}
func main() {
	fmt.Println(add("I love", " elote"))
}*/

// SWAP Function

func swap(k, c string) (string, string) {
	return c, k
}
func main() {
	m, n := swap("I Love", "eating snacks")
	fmt.Println(m, n)
}
