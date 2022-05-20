package main

import (
	"strings"
)

func contains(s []string, str []string) bool {
	for _, v := range s {
		if v == str[0] {
			return true
		}
	}

	return false
}

func isNotAscii(s string) bool {
	res := strings.ReplaceAll(s, "\n", "\\n")
	for _, val := range res {
		if val < ' ' || val > 126 {
			return true
		}
	}
	return false
}

func isEmpty(s string) bool {
	return strings.ReplaceAll(s, "\\n", "") == ""
}

func getStr(s string, x map[rune]string) string {
	if s == "" {
		return "\n"
	} else {
		res := ""
		temp := make([]string, 11)
		for _, val := range s {
			for n, r := range strings.Split(x[val], "\n") {
				temp[n] += r
			}
		}
		for _, val := range temp {
			res += val + "\n"
		}
		return res[1 : len(res)-2]
	}
}

func getMap(s string) map[rune]string {
	symbol := make(map[rune]string)
	str := ""
	j := rune(32)
	count := 0
	for _, v := range s {
		str += string(v)
		if string(v) == "\n" {
			count++
		}
		if count == 9 {
			symbol[j] = str
			str = ""
			j++
			count = 0
		}
	}
	return symbol
}
