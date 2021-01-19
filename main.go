package main

import (
	"fmt"
	"strings"
)

func appendZeros(s string, width, parts int) string {
	splitString := strings.Split(s, ".")
	appendZeroString := ""
	for _, value := range splitString {
		appendZeroString += fmt.Sprintf("%0*s", width, value)
	}
	for i := len(splitString); i < parts; i++ {
		appendZeroString += fmt.Sprintf("%0*s", width, "")
	}
	fmt.Printf("SuffixZeroVersion: [%s] => [%s]\n", s, appendZeroString)
	return appendZeroString
}

func main() {
	// it will be  combination of minor version major version and actual version
	for _, v := range [][3]string{
		{"1.39.52.70", "2.12.43.90", "1.52.29.759"},
		{"1.0", "2.0", "1.5.2"},
		{"1.02.03", "2.1.3", "1.2.3"},
	} {
		a, b, c := v[0], v[1], v[2]
		minorVersion := appendZeros(a, 5, 4)
		majorVersion := appendZeros(b, 5, 4)
		hostVersion := appendZeros(c, 5, 4)

		if minorVersion <= hostVersion && majorVersion >= hostVersion {
			fmt.Println("Successful")
		}
	}
}
