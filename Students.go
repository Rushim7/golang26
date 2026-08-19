// Task: Creating a package using struct, array , forloop and switch on students data
package main

import "fmt"

type Student struct {
	Name  string
	Age   int64
	Marks [5]float64
}

func calAverage(student *Student) (float64, float64) {
	var total float64

	for _, mark := range student.Marks {
		total += mark
	}

	average := total / float64(len(student.Marks))

	return total, average
}

func Students() {

	fmt.Println("----------------")
	fmt.Println("Student Details")

	students := [3]Student{
		{"Snow", 14, [5]float64{80, 90, 70, 89, 87}},
		{"Arya", 13, [5]float64{90, 95, 85, 92, 86}},
		{"Danny", 15, [5]float64{70, 75, 80, 90, 94}},
	}

	for _, student := range students {

		total, average := calAverage(&student)

		fmt.Println("Name:", student.Name)
		fmt.Println("Total:", total)
		fmt.Println("Average:", average)

		switch {
		case average > 85:
			fmt.Println("Grade A")

		case average > 75:
			fmt.Println("Grade B")

		default:
			fmt.Println("Unknown")
		}

		fmt.Println("----------------")
	}
}
