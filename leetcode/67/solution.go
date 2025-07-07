package _67

import (
	"strconv"
)

func addBinary(a string, b string) string {
	res := ""
	rem := 0
	indexA := len(a) - 1
	indexB := len(b) - 1
	for indexA >= 0 || indexB >= 0 || rem > 0 {
		var aVal int
		if indexA >= 0 {
			aVal = int(a[indexA] - '0')
			indexA--
		}

		var bVal int
		if indexB >= 0 {
			bVal = int(b[indexB] - '0')
			indexB--
		}

		sum := aVal + bVal + rem
		res = strconv.Itoa(sum%2) + res
		rem = sum / 2
	}
	return res
}
