package main

import (
	"fmt"
	"os"
	"encoding/csv"
	//"strconv"
	"flag"
	//"strings"
	"time"
)

func main(){
	var correctAns, incorrectAns int
	csvFileName := flag.String("csv", "problems.csv", "a CSV file in the format of Q,A")
	timeLimit := flag.Int("limit", 30, "time limit(in seconds) for the quiz to complete")
	flag.Parse()


	//fmt.Printf("**** %T", timeLimit)
	file, err := os.Open(*csvFileName)
	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	//fmt.Println("Press enter for the timer to start \n")
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemLoop:
	for idx, item := range records{
		// ans := ""
		answerCh := make(chan string)

		go func(){
			ans := ""
			fmt.Printf("Problem no.%d: %s = ", idx+1, item[0])
			fmt.Scanln(&ans)
			answerCh <-ans
		}()

		select{
		case <-timer.C:
			fmt.Println()
			//fmt.Printf("\nYou have answered %d out of %d questions correct\n", correctAns, len(records))
			break problemLoop
		case ans := <-answerCh:
			//checkAns, _ := strconv.Atoi(strings.TrimSpace(item[1]))
			if ans == item[1]{
				correctAns++
			} else {
				incorrectAns++
			}
		}
	
	}
	fmt.Printf("\nYou have answered %d out of %d questions correct\n", correctAns, len(records))
}