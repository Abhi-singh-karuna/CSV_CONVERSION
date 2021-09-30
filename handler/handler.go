package handler

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func Conv() {
	f, err := os.Open("data.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	data, _ := reader.ReadAll()

	if len(data) < 1 {
		log.Fatal("unvalid length ")
	}

	//for contaning the first row of our csv file
	Array_Header := make([]string, 0)

	//to print header length
	for _, header := range data[0] {
		Array_Header = append(Array_Header, header)
	}

	//for removing the header
	data = data[1:]

	//use for the inbuilt struct
	var DATA bytes.Buffer

	DATA.WriteString("[  \n")

	for i, d := range data {
		DATA.WriteString("    { \n")
		for j, y := range d {
			DATA.WriteString("          " + `"` + Array_Header[j] + `": `)
			_, fErr := strconv.ParseFloat(y, 32)
			_, bErr := strconv.ParseBool(y)
			if fErr == nil {
				DATA.WriteString(y)
			} else if bErr == nil {
				DATA.WriteString(strings.ToLower(y))
			} else {
				DATA.WriteString((`"` + y + `"`))
			}
			//end of property
			if j < len(d)-1 {
				DATA.WriteString(", \n")
			}

		}
		//end of object of the array
		DATA.WriteString("\n    }")
		if i < len(Array_Header)-1 {
			DATA.WriteString(",\n")
		}

	}
	DATA.WriteString("]   \n")

	jsonFile, err := os.Create("./data.json")

	fmt.Println("csv file is converted in to:  data.Json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(DATA.Bytes())
	jsonFile.Close()

}
