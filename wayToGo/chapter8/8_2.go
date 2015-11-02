package main
import "fmt"

func main() {
	week := map[string]int{
		"Monday":1,
		"Tuesday":2,
		"Wednesday":3,
		"Thrusday":4,
		"Friday":5,
		"Saturday":6,
		"Sunday":7,
	}

	for key, value := range week{
		fmt.Printf("key: %s, value: %d\n",key, value)
	}
	
	if _,isPresent:= week["Tuesday"]; isPresent{
		fmt.Printf("is Present Tuesday? %t\n",isPresent)
	}
	if _,isPresent:= week["Hollyday"]; isPresent{
		fmt.Printf("is Present Hollyday? %t\n",isPresent)
	}
}
