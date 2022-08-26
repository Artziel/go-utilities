package main

import (
	"encoding/json"
	"fmt"
	"strings"

	Util "github.com/artziel/go-utilities"
)

func printSystemInfo() {
	sys, err := Util.GetSystemInfo()
	if err != nil {
		fmt.Printf("Error getting System Info: %s", err)
		return
	}

	strinfo, _ := json.MarshalIndent(sys, "", "    ")

	fmt.Printf("System Info\n%s\n%s\n", strings.Repeat("-", 60), strinfo)
}

func main() {

	printSystemInfo()
}
