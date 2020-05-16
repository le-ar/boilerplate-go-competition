package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const isProd = false
const readFromFile = true
const fileIn = "input-201.txt"
const fileOut = "input-201.a.txt"

var fileInput *os.File
var fileInScanner *bufio.Scanner
var fileOutput *os.File
var fileOutWriter *bufio.Writer

func main() {
	start := time.Now()

	initReader()
	resolve()
	if fileInput != nil {
		fileInput.Close()
	}

	if fileOutWriter != nil {
		fileOutWriter.Flush()
	}
	if fileOutput != nil {
		fileInput.Close()
	}

	if !isProd {
		fmt.Printf("Executed in %v\n", time.Since(start))
	}
}

func initReader() {
	var err error

	fileInput, err = os.Open(fileIn)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	fileOutput, err = os.Create(fileOut)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	fileOutWriter = bufio.NewWriter(fileOutput)
}

var readedLine string = ""
var readWordScanner *bufio.Scanner

func getLine(v *string) bool {
	if fileInScanner == nil {
		fileInScanner = bufio.NewScanner(fileInput)
		fileInScanner.Split(bufio.ScanLines)
	}

	if fileInScanner.Scan() {
		line := fileInScanner.Text()
		*v = line
		return true
	}

	return false
}

func updateWordLineScanner() {
	reader := strings.NewReader(readedLine)
	readWordScanner = bufio.NewScanner(reader)
	readWordScanner.Split(bufio.ScanWords)
}

func getWord(v *string) bool {
	if readWordScanner == nil {
		updateWordLineScanner()
	}
	for !readWordScanner.Scan() {
		if !getLine(&readedLine) {
			return false
		}
		updateWordLineScanner()
	}

	word := readWordScanner.Text()
	*v = word
	return true
}

func getInt64(v *int64) bool {
	var word string
	if !getWord(&word) {
		return false
	}

	i, err := strconv.ParseInt(word, 10, 64)
	if err != nil {
		panic(err)
	}

	*v = i
	return true
}

func getString(v *string) bool {
	return getWord(v)
}

func writeLine(a interface{}) {
	if fileOutWriter != nil {
		write(a)
		fileOutWriter.WriteString("\n")
	}
}

func write(a interface{}) {
	switch v := a.(type) {
	case int64:
		write(strconv.FormatInt(v, 10))
	case string:
		fileOutWriter.WriteString(v)
	}
}

func resolve() {

}
