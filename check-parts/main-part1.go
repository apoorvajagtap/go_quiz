package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
	"flag"
	"strings"
	//"time"
)

func main(){
	var ans, correctAns, incorrectAns int
	csvFileName := flag.String("csv", "problems.csv", "a CSV file in the format of Q,A")
	flag.Parse()


	//fmt.Printf("**** %T", timeLimit)
	file, err := os.Open(*csvFileName)
	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for idx, item := range records{
		fmt.Printf("Problem no.%d: %s = ", idx+1, item[0])
		fmt.Scanln(&ans)

		checkAns, _ := strconv.Atoi(strings.TrimSpace(item[1]))

		if ans == checkAns{
			correctAns++
		} else {
			incorrectAns++
		}
	}

	fmt.Printf("You have answered %d out of %d questions correct\n", correctAns, correctAns+incorrectAns)
}