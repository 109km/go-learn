package main

import "fmt"
import "time"

func main() {

	now := time.Now()
	hour := now.Hour()
	switch {
	case hour <= 3:
		fmt.Println(hour, "is in the night")
	case hour <= 12:
		fmt.Println(hour, "is in the morning")
	case hour <= 18:
		fmt.Println(hour, "is in the afternoon")
	default:
		fmt.Println(hour, "is in the evening")
	}

	printDataType := func(data interface{}) {
		switch t := data.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Printf("other,%T\n", t)
		}
	}
	printDataType(true)
}
