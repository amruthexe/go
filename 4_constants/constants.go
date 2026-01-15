package main

import "fmt"

const age =30
// name:="golang" not possible outside the function

const (
	host=3000 
	request="https"
)

func main() {

	const name string = "go lang"
	// name="js" ❌
	fmt.Println(name)
	fmt.Println(host,request)
}