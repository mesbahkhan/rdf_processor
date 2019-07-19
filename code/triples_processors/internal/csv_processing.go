package internal

import (
	"encoding/csv"
	"fmt"
	"github.com/TomOnTime/utfutil"
	"github.com/saintfish/chardet"
	"io"
	"os"
	"strings"
)

func Open_csv_file(csv_filename string) (*os.File, []byte) {

	csv_file, err :=
		os.OpenFile(
			csv_filename,
			os.O_CREATE|os.O_RDWR,
			0777)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	// charset detection
	buffer := make([]byte, 32<<10)
	size, _ := io.ReadFull(csv_file, buffer)
	input := buffer[:size]

	var detector = chardet.NewTextDetector()
	result, err := detector.DetectBest(input)

	fmt.Printf("File Ecoding : %s , Language: %s \n", result.Charset, result.Language)

	//#TODO - ADD Switch here based on the the encoding detected.

	var csv_data_utfutls []byte

	switch result.Charset {

	case "UTF-8":
		csv_data_utfutls, _ = utfutil.ReadFile(csv_filename, utfutil.UTF8)
	case "UTF-16LE":
		csv_data_utfutls, _ = utfutil.ReadFile(csv_filename, utfutil.UTF16LE)
	case "ISO-8859-1":
		csv_data_utfutls, _ = utfutil.ReadFile(csv_filename, utfutil.UTF8)
	}

	return csv_file, csv_data_utfutls

}

func Read_csv_data(csv_file_name string, delimiter string) [][]string {

	csv_file,
		csv_file_data :=
		Open_csv_file(
			csv_file_name)

	csv_dataset := Read_csv_to_slice(
		csv_file,
		csv_file_data,
		delimiter)

	return csv_dataset

}

func Read_csv_to_slice(csv_file *os.File, csv_data_utfutls []byte, delimiter string) [][]string {

	//--END OF UTFUTIL reader
	/***********easycsv reader
	easy_csv_reader := easycsv.NewReaderFile(file_name ,  easycsv.Option{
		Comma: '\t',
	})
	var rawCSVdata [][]string
	easy_csv_reader.ReadAll(&rawCSVdata)
	* END OF EASY CSV REader*/

	// csv reader standard

	csv_reader := csv.NewReader(strings.NewReader(string(csv_data_utfutls)))
	//csv_reader.TrimLeadingSpace = true
	csv_reader.FieldsPerRecord = -1 // see the Reader struct information below

	switch delimiter {
	case "tab":
		csv_reader.Comma = '\t'
	case "":

	}

	csv_reader.LazyQuotes = true
	csv_data, csv_reader_error := csv_reader.ReadAll()

	if csv_reader_error != nil {
		panic(csv_reader_error)
	}
	//--- END OF CSV READER
	fmt.Println(len(csv_data))

	//raw csv data generation  (NOT USED) -
	//#TODO - move this out.
	rawCSVdata_bytes := make([][]byte, len(csv_data)*len(csv_data[0]))
	for _, raw_csv_data_row := range csv_data {
		for _, raw_csv_data_column := range raw_csv_data_row {

			rawCSVdata_bytes = append(rawCSVdata_bytes, []byte(raw_csv_data_column))

		}
	}

	//  -- end of raw data bytes generation

	return csv_data
}
