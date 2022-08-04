package NCBICdd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type cddWrapper struct {
	Cdd Cdd `json:"Cdd"`
}

func Convert_cddfile_json() {

	var file1 = "example/cd00180_notree_4.json"
	jsonFile, err := os.Open(file1)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	var cdd cddWrapper
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &cdd); err != nil {
		log.Fatal(err)
	}

	output, err := json.MarshalIndent(cdd,"","    ")
	if err != nil {
		log.Fatal(err)
	} else {
		os.WriteFile("example/output.json", output, os.ModePerm)
	}
	// return &cdd
}

func Convert_cddfile() *Cdd {
	var cdd Cdd
	var file1 = "example/test.xml"
	xmlFile, err := os.Open(file1)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	if err := xml.Unmarshal(byteValue, &cdd); err != nil {
		log.Fatal(err)
	}
	return &cdd
}
