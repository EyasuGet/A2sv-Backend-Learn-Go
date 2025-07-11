package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter Your name: ")
	fmt.Scan(&name)

	var numOfSub int
	fmt.Print("Enter the number of subjects you take? ")
	fmt.Scan(&numOfSub)

	report := make(map[string]int)
	total := 0


	for i := 0; i < numOfSub; i++ {

		var subject string
		var grade int

		fmt.Printf("Enter subject %d name: ", i+1)
		fmt.Scan(&subject)

		for {
			fmt.Printf("Enter grade for %s (0-100): ", subject)
			fmt.Scan(&grade)

			if grade >= 0 && grade <= 100{
				break
			}

			fmt.Println("Invalid grade")
		}

		report[subject] = grade
		total += grade
	}

	average := calculateAverage(total, numOfSub)

	fmt.Printf("Name: %v\n", name)
	fmt.Printf("You are learning %v\n subjects", numOfSub)

	for subject, grade := range report{
		fmt.Printf(" %s: %d\n", subject, grade)
	}

	fmt.Printf("Your average is %v\n", average)

}

func calculateAverage(total int, number int) float64{
	return float64(total) / float64(number)
}
