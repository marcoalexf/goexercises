package main

import (
	csv "encoding/csv"
	"fmt"
	"os"
	"time"
)

type question struct {
	question         string
	expectedResponse string
}

var messages chan string = make(chan string)

func readCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file.", filename)
		return [][]string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		fmt.Println("Error parsing csv file.", filename)
		return [][]string{}, err
	}

	return lines, nil
}

func mapLinesToStructList(filename string) ([]question, error) {
	var resultList []question

	lines, err := readCSV(filename)

	if err != nil {
		fmt.Println("Error parsing csv file.", filename, err)
		return resultList, nil
	}

	for _, line := range lines {
		resultList = append(resultList, question{
			question:         line[0],
			expectedResponse: line[1],
		})
	}

	return resultList, nil
}

func timer() {
	timer := time.NewTimer(time.Second * 5)
	defer timer.Stop()
	messages <- "Tick"
	//stopped := timer.Stop()

	/*if stopped {
		messages <- "ended"
	}*/
}

func runGame(questions []question) (bool, int) {
	alive := true
	score := 0

	for _, question := range questions {
		go timer()
		fmt.Println(question.question)
		var answer string
		fmt.Scan(&answer)

		if answer != question.expectedResponse {
			alive = false
			break
		} else {
			score += 1
		}
	}

	return alive, score
}

func main() {
	fmt.Println("Starting to parse..")
	questions, err := mapLinesToStructList("problems.csv")

	if err != nil {
		fmt.Println("Ups, somethings went wrong.", err)
	}

	go runGame(questions)
	for {
		msg := <-messages
		fmt.Println(msg)
		/*if msg == "stop" {
			break
		} */
	}

	// gameResult, score := runGame(questions)

	/*if gameResult {
		fmt.Println("Congratulations, you won!")
	} else {
		fmt.Println("You lost!")
	}*/

	//fmt.Printf("Final score of %d \n", score)
	fmt.Println("Thank you for playing!")
}
