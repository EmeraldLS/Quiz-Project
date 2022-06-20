package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to Emerald Quizgame")
	var first_name string
	var last_name string
	var full_name string
	var age uint
	var score float64 = 0
	var total_questions float64 = 5
	failed_questions := []string{}
	// Taking input for name
	fmt.Printf("Kindly provide your full name: ")
	fmt.Scanf("%v %v\n", &first_name, &last_name)
	full_name = first_name + " " + last_name
	fmt.Printf("Hi %v , You are honored.\n", full_name)

	// Taking input for age and it conditions
	fmt.Print("How old are you: ")
	fmt.Scanf("%v\n", &age)
	if age < 15 {
		fmt.Printf("Hi %v, you cannot sit for the quiz\n", full_name)
		return
	} else {
		fmt.Printf("\nREAD THIS!!!\nYou have a total of 10 questions and you are expected to answer all. answer questions by inputing the option(e.g A)\n\n")
	}

	// Questions with answers

	// Question 1
	question1 := "1.) Today is khadija’s birthday. After one year, she will become two times as she was ten years ago. current age of Khadija?"
	fmt.Println(question1)
	question1_options := [4]string{"A. 19", "B. 20", "C. 21", "D. 22"}
	var question1_input string
	question1_answer := "c"
	for i := 0; i < 4; i++ {
		fmt.Println(question1_options[i])
	}
	fmt.Print("Choose an option: ")
	fmt.Scan(&question1_input)
	if strings.ToLower(question1_input) == question1_answer {
		score++
	} else {
		failed_questions = append(failed_questions, "Question 1")
	}

	// Question 2
	question2 := "\n\n2.) If the average marks of three batches of 55, 60 and 45 students respectively is 50, 55, 60, then the average marks of all the students is__________?"
	fmt.Println(question2)
	question2_options := [4]string{"A. 53.33", "B. 54.68", "C. 55", "D. None of there"}
	var question2_input string
	question2_answer := "b"
	for i := 0; i < 4; i++ {
		fmt.Println(question2_options[i])
	}
	fmt.Print("Choose an option: ")
	fmt.Scan(&question2_input)
	if strings.ToLower(question2_input) == question2_answer {
		score++
	} else {
		failed_questions = append(failed_questions, "Question 2")
	}

	// Question 3
	question3 := "\n\n3.) The average weight of 16 boys in a class is 50.25 kg and that of the remaining 8 boys is 45.15 kg. Find the average weights of all the boys in the class."
	fmt.Println(question3)
	question3_options := [4]string{"A. 47.55kg", "B. 48.kg", "48.55kg", "49.25kg"}
	var question3_input string
	question3_answer := "c"
	for i := 0; i < 4; i++ {
		fmt.Println(question3_options[i])
	}

	fmt.Print("Choose an option: ")
	fmt.Scan(&question3_input)
	if strings.ToLower(question3_input) == question3_answer {
		score++
	} else {
		failed_questions = append(failed_questions, "Question 3")
	}

	// Question 4
	question4 := "\n\n4.) The average weight of A, B and C is 45 kg. If the average weight of A and B be 40 kg and that of B and C be 43 kg, then the weight of B is___________?"
	fmt.Println(question4)
	question4_options := [4]string{"A. 17kg", "B. 20kg", "C. 26kg", "D. 31kg"}
	var question4_input string
	question4_answer := "d"
	for i := 0; i < 4; i++ {
		fmt.Println(question4_options[i])
	}
	fmt.Print("Choose an option: ")
	fmt.Scan(&question4_input)
	if strings.ToLower(question4_input) == question4_answer {
		score++
	} else {
		failed_questions = append(failed_questions, "Question 4")
	}

	// Question 5
	question5 := "\n\n5.) The average of 18 observations was calculated and it was 124. Later on it was discovered that two observations 46 and 82 were incorrect. The correct values are 64 and 28. The correct average of 18 observations is___________?"
	fmt.Println(question5)
	question5_options := [4]string{"A. 123", "B. 137", "C. 121", "D. 122"}
	var question5_input string
	question5_answer := "d"
	for i := 0; i < 4; i++ {
		fmt.Println(question5_options[i])
	}
	fmt.Print("Choose an option: ")
	fmt.Scan(&question5_input)
	if strings.ToLower(question5_input) == question5_answer {
		score++
	} else {
		failed_questions = append(failed_questions, "Question 5")
	}

	percentage_score := (score / total_questions) * 100
	// fmt.Println("percentage score is ",percentage_score)
	// fmt.Printf("\nYou scored %d out of %v\n", int(score), total_questions)
	fmt.Printf("\n\nYou scored %d%%. Good work %v\n\n", int(percentage_score), full_name)
	fmt.Println("Below is/are the list of question you scored wrong.")
	for i := 0; i < len(failed_questions); i++ {
		fmt.Println(failed_questions[i])
	}
	if percentage_score < 50 {
		fmt.Printf("You can try again, %v\n", first_name)
	} else {
		fmt.Println("\nWow! You are incredible")
	}
	fmt.Printf("\nHey %v don't for get to follow me on twitter @EmeraldLS. ", full_name)

}
