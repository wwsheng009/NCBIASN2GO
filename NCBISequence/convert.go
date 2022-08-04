package NCBISequence

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type seqWrapper struct {
	Seq SeqEntry `json:"Seq_entry"`
}

func convert_json_data(fname string) {

	var file1 = fmt.Sprintf("example/%s.json", fname)
	jsonFile, err := os.Open(file1)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	var sample seqWrapper
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &sample); err != nil {
		log.Fatal(err)
	}

	output, err := json.MarshalIndent(sample, "", "    ")
	if err != nil {
		log.Fatal(err)
	} else {
		os.WriteFile(fmt.Sprintf("example/%s_out.json", fname), output, os.ModePerm)
	}
	// return &cdd
}
