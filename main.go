package main

import (
	"fmt"
	"strconv"
	"strings"
)

func counts(info string, m map[string]int) map[string]int {
	comma := strings.Index(info, ",")
	count, _ := strconv.Atoi(info[0:comma])

	domain := info[comma+1:]

	if val, found := m[domain]; found {
		m[domain] = val + count
	} else {
		m[domain] = count
	}

	dot := strings.Index(domain, ".")

	for dot > -1 {
		if val, found := m[domain]; found {
			m[domain] = val + count
		} else {
			m[domain] = count
		}

		domain = domain[dot+1:]
		dot = strings.Index(domain, ".")
	}

	return m
}

func spaces(number int) string {
	str := ""
	for i := 0; i < number; i++ {
		str += " "
	}
	return str
}

func main() {
	list := []string{
		"900,google.com",
		"333,mail.google.com",
		"22,yahoo.com",
		"180,example.yahoo.com",
		"18,mail.example.yahoo.com",
	}

	var m map[string]int = make(map[string]int)

	for _, info := range list {
		m = counts(info, m)
	}

	maxlen := 0

	for key := range m {
		if len(key) > maxlen {
			maxlen = len(key)
		}
	}

	for key, value := range m {
		fmt.Printf("%s : %d\n", key+spaces(maxlen-len(key)), value)
	}
}
