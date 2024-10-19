package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Rankings struct {
	Keyword  string
	GetCount uint32
	Engine   string
	Locale   string
	Mobile   bool
}

func main() {
	var jsonBlob = []byte(`
        {"keyword":"hipaa compliance form", "get_count":157, "engine":"google", "locale":"en-us", "mobile":false}
    `)
	rankings := Rankings{}
	err := json.Unmarshal(jsonBlob, &rankings)
	if err != nil {
		// nozzle.printError("opening config file", err.Error())
	}

	rankingsJson, _ := json.Marshal(rankings)
	err = ioutil.WriteFile("output.json", rankingsJson, 0644)
	fmt.Printf("%+v", rankings)
}
