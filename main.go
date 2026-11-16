package main

import "fmt"

// calculateStats calculates the total and average from a slice of marks
func calculateStats(marks []int) (int, float64) {
	total := 0
	for _, mark := range marks {
		total += mark
	}
	average := float64(total) / float64(len(marks))
	return total, average
}

// calculateGrade determines the letter grade based on the average
func calculateGrade(average float64) string {
	if average >= 80 {
		return "A"
	} else if average >= 60 {
		return "B"
	} else if average >= 40 {
		return "C"
	} else {
		return "F"
	}
}

func main() {
	var name string
	fmt.Print("Enter student name: ")
	fmt.Scanln(&name)

	var marks []int
	for i := 1; i <= 3; i++ {
		var mark int
		for {
			fmt.Printf("Enter mark %d: ", i)
			fmt.Scanln(&mark)
			if mark >= 0 && mark <= 100 {
				break
			}
			fmt.Println("Invalid mark. Please enter a value between 0 and 100.")
		}
		marks = append(marks, mark)
	}

	total, average := calculateStats(marks)
	grade := calculateGrade(average)

	result := "Fail"
	if average >= 40 {
		result = "Pass"
	}

	fmt.Printf("Student: %s\n", name)
	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Average: %.2f\n", average)
	fmt.Printf("Grade: %s\n", grade)
	fmt.Printf("Result: %s\n", result)
}
