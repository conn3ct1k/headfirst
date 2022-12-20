package main

import (
	"fmt"
)

func main() {
	// var status string

	// fmt.Print("Enter a grade: ")
	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// input = strings.TrimSpace(input)

	// grade, err := strconv.ParseFloat(input, 64)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if grade >= 60 {
	// 	status = "passing"
	// } else {
	// 	status = "failing"
	// }

	// fmt.Println("A grade of", grade, "is", status)
	for x := 1; x <= 3; x++ {
		fmt.Print(x)
	}
}
