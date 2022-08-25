package main

import (
	"encoding/json"
	"fmt"
	"math"
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

	dec := 3.0
	frac := math.Mod(dec, 1)

	fmt.Printf("%v\n", frac)

	// printSystemInfo()
}
