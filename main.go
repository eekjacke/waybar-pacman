package main

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
)

type Output struct {
	Text       string `json:"text"`
	Alt        string `json:"alt"`
	Tooltip    string `json:"tooltip"`
	Class      string `json:"class"`
	Percentage string `json:"percentage"`
}

func main() {
	cmd1 := exec.Command("checkupdates")
	var out1 bytes.Buffer
	cmd1.Stdout = &out1
	if err := cmd1.Run(); err != nil {
		//fmt.Println("Error running checkupdates:", err)
		os.Exit(1)
	}
	updates := out1.String()
	//fmt.Println(updates)
	out1Reader := bytes.NewReader([]byte(updates))
	cmd2 := exec.Command("wc", "-l")
	cmd2.Stdin = out1Reader
	var out2 bytes.Buffer
	cmd2.Stdout = &out2
	if err := cmd2.Run(); err != nil {
		//fmt.Println("Error running wc -l:", err)
		os.Exit(2)
	}
	//fmt.Println(out2.String())
	numUpdates := out2.String()
	numUpdates = numUpdates[:len(numUpdates)-1]
	output := Output{
		Text:       numUpdates,
		Alt:        "",
		Tooltip:    updates,
		Class:      "pacman",
		Percentage: "",
	}
	//fmt.Println(output)
	jsonString, err := json.Marshal(output)
	if err != nil {
		//fmt.Println("error marshaling json: ", err)
		os.Exit(3)
	}
	var buffer bytes.Buffer
	buffer.Write(jsonString)
	buffer.WriteString("\n")
	buffer.WriteTo(os.Stdout)
}
