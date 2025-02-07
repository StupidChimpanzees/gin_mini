package utils

import "strings"

func NameToLower(s string) string {
	var b [128]byte
	var l = 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
			b[l] = '_'
			l++
			b[l] = c
		} else {
			b[l] = c
		}
		l++
	}
	if b[0] == '_' {
		return strings.TrimSpace(string(b[1:]))
	} else {
		return strings.TrimSpace(string(b[:]))
	}
}

func NameToUpper(s string) string {
	var b [128]byte
	var l = 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '_' {
			c := s[i+1]
			c -= 'a' - 'A'
			b[l] = c
			i++
		} else {
			b[l] = c
		}
		l++
	}
	if 'a' <= b[0] && b[0] <= 'z' {
		b[0] -= 'a' - 'A'
		return strings.TrimSpace(string(b[:]))
	} else {
		return strings.TrimSpace(string(b[:]))
	}
}
