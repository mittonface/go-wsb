package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	tickerFile, err := os.Open("./nasdaqtraded.csv")

	if err != nil {
		log.Fatalln("Could not open ticker file for parsing ", err)
	}

	r := csv.NewReader(tickerFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record[1])
	}

}
