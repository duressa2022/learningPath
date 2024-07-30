/*
  Go based grading system for calculating grades of the students
        uses maps and slices for storing stud name name grade information
		Deletes every information for the students when it is done
	Simple console based based application for grading system

*/

package main

import "fmt"

//main function for grading system

func main() {
	fmt.Println("WELCOME TO GRADING SYSTEM FOR STUDENT GRADE GRADE EVALUATION")
	//init map and string to store user information and permission
	var userInformation = make(map[string][]float64)
	var permission string
	for {
		//get permission from the usery
		fmt.Print("Do you want calculate your grade y or n: ")
		fmt.Scan(&permission)
		//if the permission is yes from the user
		if permission == "y" {
			//get name of the student of the user
			var studName string
			fmt.Print("Please enter your name: ")
			fmt.Scan(&studName)
			//stote the data inside the map
			userInformation[studName] = []float64{}
			//Get number of subjects from the user
			var numberSubject int
			fmt.Print("Enter number of your subjects: ")
			fmt.Scan(&numberSubject)
			//Get grade for each subjets from the user
			for index := 1; index <= numberSubject; index++ {
				var grade float64
				fmt.Printf("Enter subject %v: ", index)
				fmt.Scan(&grade)
				//store the grade inside the the map
				userInformation[studName] = append(userInformation[studName], grade)

			}
			//call the function to calculate grade for the student
			average := calculateGrade(userInformation[studName])
			//display the information of the user on the disply feild
			fmt.Println("============================================y")
			printInformation(userInformation, studName, average)
			fmt.Println("============================================")

		} else {
			//break if the condition is false
			fmt.Println("We are waiting for you to come back")
			break
		}

	}
}

//function to calculate the grade for the student
func calculateGrade(grades []float64) float64 {
	//compute the length of the grades slices
	length := len(grades)
	//compute the summation of the grades
	summation := 0.0
	for _, grade := range grades {
		summation += grade

	}
	//compute the average for the user
	return summation / float64(length)

}

//function to display information of the student
func printInformation(information map[string][]float64, name string, average float64) {
	//display the name to the console
	fmt.Printf("Student name is: %v", name)
	fmt.Println("")
	//display the grades to the console
	for index, grade := range information[name] {
		fmt.Printf("Grade for subject %v: %v\n", index+1, grade)
	}
	//display the averge to the console
	fmt.Printf("your average is: %v\n", average)

}
