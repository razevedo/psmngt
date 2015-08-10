package utils

import "bytes"

func ConcateStrings(input ...string) string {
	var buffer bytes.Buffer
	for _, s := range input {
		buffer.WriteString(s)
	}
	return buffer.String()
}
