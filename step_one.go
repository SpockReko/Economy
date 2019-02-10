package main

import (
	"os"
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
	//YearMonth := os.Args[2]

	file, err := os.Open(filePath)
	handleErr(err)
	defer func() {
		err = file.Close()
		handleErr(err)
	}()

}

func handleErr(err error) {
	if err != nil {
		panic("Coulen't open the file")
	}
}
