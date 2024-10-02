package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	csvFilename := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "time limit of quiz game")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open csv file %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provied csv file")
	}
	problemes := parseLines(lines) // parsing 2d slice to a slice

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) // this will return a channel C

	correct := 0
	for i, p := range problemes {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)

		answerCh := make(chan string) // a channel to store user answer

		//  anonymous function
		go func() { // goroutine to not block the code causing by scanf
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select { // which ever channel returns first that case will be selected and runs accordngly

		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problemes))
			return
		case answer := <-answerCh: // reading from answe channel and comparing with the correct answe , if correct will increment the correct count
			if answer == p.answer {
				correct++
			}

		}

	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problemes))

}

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)

}
