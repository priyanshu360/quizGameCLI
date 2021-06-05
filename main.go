package main

import(
	"encoding/csv"
	"fmt"
	"flag"
	"os"
	"strings"
	"time"
)


func main(){
	quizFile := flag.String("csv", "problems.csv", "a csv file of format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*quizFile)

	if(err != nil){
		errorExit(fmt.Sprintf("Failed to open the CSV file: %s\n", *quizFile))
	} 

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		errorExit("Failed to parse the provided CSV file.")
	}

	questions := Lines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	score := 0

	loop:
	for i, task := range questions{
		fmt.Printf("Task #%d: %s = ", i + 1, task.q)

		answerCh := make(chan string)
		go func(){
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}()

		
		select{
			case <-timer.C:
				fmt.Println()
				break loop
			case ans := <-answerCh:
				if(ans == task.a){
					score++
				}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", score, len(questions))

}


func Lines(lines [][]string) []question{
	ret := make([]question, len(lines))

	for i, line := range lines{
		ret[i] = question{
			q : line[0],
			a : strings.TrimSpace(line[1]),
		}
	}

	return ret
}


type question struct{
	q string
	a string
}

func errorExit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}