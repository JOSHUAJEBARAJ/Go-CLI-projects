package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", " a csv file")

	timeLimit := flag.Int("limit", 30, "time limit in seconds")
	_ = csvFilename
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {

		exit(fmt.Sprintf("Failed to open %s\n", *csvFilename))
	}
	// csv reader
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the file ")
	}
	// fmt.Println(lines)
	problems := parseLines(lines)
	// printing quiz
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem %d :%s", i+1, p.q)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer

		}()
		select {
		case <-timer.C:
			fmt.Println("Total correct", correct)
			return
		case answer := <-answerCh:

			if answer == p.a {
				correct++
				fmt.Println("Correct")
			}

		}
	}
}

// quiz type

type problem struct {
	q string
	a string
}

// parsing function
func parseLines(lines [][]string) []problem {

	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

// exit method
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)

}
