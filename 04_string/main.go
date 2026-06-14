package main
import (
	"fmt"
	"strings"
)

func main(){

	firstName := "prasanta"
	lastName := "maity"

	fullName := firstName + " " + lastName

	fmt.Println(fullName)

	fmt.Println(strings.ToUpper(fullName))
}