package main

import (
	"fmt"
	"strconv"
	"strings"
)

func appendZeros(s string, zeros int) string {
	splitString := strings.Split(s, ".")
	appendZeroString := ""

	for _, value := range splitString {
		appendZeroString += fmt.Sprintf("%0*s", zeros, value)
	}

	fmt.Printf("SuffixZeroVersion: [%s] => [%s]\n", s, appendZeroString)
	return appendZeroString
}

func convertToString(alphaNumeric string) string {

	newString := ""
	for _, r := range alphaNumeric {
		if int(r) >= 65 && int(r) <= 90 {
			newString += string(strconv.Itoa(int(r)))
		} else {
			newString += string(int(r))
		}
	}

	return newString
}

func main() {
	// it will be  combination of minor version major version and actual version
	for _, v := range [][3]string{
		{"1.39.52.70", "1.52.300.90", "1.52.29.759"},
		{"1.4.6", "1.5.001", "1.4.9"},
		{"1.02.03", "2.1.003", "1.2.3"},
		{"1.02.03", "2.1.03", "1.2.3"},
		{"1.0.0.007", "6.0.0.0", "2.0.01.05"},
		{"1.003.0003.0", "6.001.05.0", "2.0002.01.05"},
		{"5.06.05.0090", "6.02.80.020", "5.0006.05.05"},
		{"2.600.500.90", "6.70.8.02", "5.0800.09.2"},
		{"1.03.c9.1", "2.006.0d4.02", "2.009.09.2"},
		{"1.03.00.00", "2.003.00.00", "1.003.01.00"},
		{"1.03.00.00", "2.003.00.00", "2.03.00.00"},
	} {
		a, b, c := v[0], v[1], v[2]
		minorVersion := appendZeros(convertToString(strings.ToUpper(a)), 8)
		majorVersion := appendZeros(convertToString(strings.ToUpper(b)), 8)
		hostVersion := appendZeros(convertToString(strings.ToUpper(c)), 8)

		if minorVersion <= hostVersion && majorVersion >= hostVersion {
			fmt.Println("Successful")
		} else {
			fmt.Println("Failed")
		}
	}
}
