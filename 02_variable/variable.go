package main
import "fmt"

func main(){


	//var name type

	var ageYear int = 45
	fmt.Println("ageYear:",ageYear)

	var rating float64 = 4.34
	fmt.Println("rating: ",rating)


	var city string;
	city = "landon"

	fmt.Println("city: ",city)


	var chennal = "prasant" // inferred to string
	fmt.Println(chennal)


	// :=
	subscribers := 500
	subscribers += 100
	fmt.Println(subscribers)


	likes , command := 100 , 30
	fmt.Println(likes,command)





}