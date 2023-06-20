package main

import "fmt"

func main() {
	
	subjectMarks := map[string]float32{"English": 89, "Physics": 98,"Chemistry": 72}
	fmt.Println(subjectMarks)

	fmt.Println(subjectMarks["English"])

	fmt.Println("Map: ",subjectMarks)

	subjectMarks["English"] = 90

	fmt.Println("Updated Map: ",subjectMarks )

	subjectMarks["French"] = 55

	fmt.Println(subjectMarks)

	delete(subjectMarks, "French")

	fmt.Println(subjectMarks)

	for subject, subjMarks := range subjectMarks {
		fmt.Println(subject, subjMarks)
	}

}