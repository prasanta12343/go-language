package main

import(
	"fmt"
)

func main(){

	// score := 74

	// if score >= 90 {
	// 	fmt.Println("student get A")
	// }else if score >= 50 {
	// 	fmt.Println("student Get B")
	// }else{
	// 	fmt.Println("Student Get D")
	// }


	item := 3
	pricePerItem := 4
	if total := item * pricePerItem ; total >= 100 {
		fmt.Println("Eligible For Shipping")
	} else {
		fmt.Println("Not Eligible For Shipping")
	}

}