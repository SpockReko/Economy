package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/**
go run step_one.go <path> <YYYY-MM>

1. get file-path from args
2. get year and month from args
3. open file for read and write
for each row {
	1. skip if the 1date is not in the given date from args
	2. replace comma with a punctioation
	3. replace semicolon with a comma
	4. split the text row by comma as a separator
	5. remove second date and total column
	   description, date1, date2, transaction, saldo -> description, date1, transaction
	6. present row info on the console
	7. lisen to input
	8. add the given label to a new column
		description, date1, transaction, label
	9. write in a new file named handled_<original file-name>
}
*/
func main() {

	filePath := os.Args[1]
	yearMonth := os.Args[2]

	file, err := os.Open(filePath)
	handleErr(err)
	defer func() {
		err = file.Close()
		handleErr(err)
	}()

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ";")
		if strings.Contains(row[2], yearMonth) {
			row[3] = strings.Replace(row[3], ",", ".", -1)
			row[3] = strings.Replace(row[3], " ", "", -1)
			row[3] = strings.Split(row[3], ".")[0]
			println(row[0], row[2], row[3])
			fmt.Print("c (clara), s (stefan), b (båda)>")
			label, _ := reader.ReadString('\n')
			label = strings.Replace(label, "\n", "", -1)
			switch label {
			case "c":
				newSlice := []string{row[0], row[2], row[3], label}
				newRow := []byte(strings.Join(newSlice, ","))
				writeToFile("clara.csv", fmt.Sprintln(string(newRow)))
				break
			case "s":
				newSlice := []string{row[0], row[2], row[3], label}
				newRow := []byte(strings.Join(newSlice, ","))
				writeToFile("stefan.csv", fmt.Sprintln(string(newRow)))
				break
			case "b":
				newSlice := []string{row[0], row[2], row[3]}
				newRow := []byte(strings.Join(newSlice, ","))
				writeToFile("gemensam.csv", fmt.Sprintln(string(newRow)))
				break
			default:
				newSlice := []string{row[0], row[2], row[3], "?"}
				newRow := []byte(strings.Join(newSlice, ","))
				writeToFile("undefined.csv", fmt.Sprintln(string(newRow)))
			}
		}
	}
	fmt.Println("Done")
}

func handleErr(err error) {
	if err != nil {
		panic("Coulen't open the file")
	}
}

func writeToFile(fileName string, text string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		ioutil.WriteFile(fileName, []byte(""), 0777)
	}
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 7777)
	file.WriteString(text)
	file.Close()
}
