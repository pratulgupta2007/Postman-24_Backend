package main

import (
	"fmt"
	"strings"
)

func IsVowel(s string) bool {
	return strings.ContainsAny(s, "aieou")
}

func IsConsonant(s string) bool {
	return strings.ContainsAny(s, "bcdfghjklmnpqrstvwxyz")
}

func m(s string) int {
	m := 0
	flag := 0
	for k := range s {
		curr := string(s[k])
		if IsVowel(curr) {
			flag = 1
		} else if IsConsonant(curr) {
			if flag == 1 {
				m += 1
				flag = 0
			}
		}
	}
	return m

}

func _s(s string) bool {
	return strings.HasSuffix(s, "s")
}

func _v_(s string) bool {
	return strings.ContainsAny(s, "aieou")
}

func _d(s string) bool {
	return s[len(s)-1] == s[len(s)-2] && IsConsonant(string(s[len(s)-1]))
}

func _o(s string) bool {
	return IsConsonant(string(s[len(s)-3])) &&
		IsVowel(string(s[len(s)-2])) &&
		strings.ContainsAny(string(s[len(s)-1]), "bcdfghjklmnpqrstvz")
}

func repsuff(s, s1, s2 string) string {
	return strings.TrimSuffix(s, s1) + s2
}

func remsuff(s, s1 string) string {
	return strings.TrimSuffix(s, s1)
}

func porter(s string) {
	//1a
	if strings.HasSuffix(s, "sess") {
		s = repsuff(s, "sess", "ss")
	} else if strings.HasSuffix(s, "ies") {
		s = repsuff(s, "ies", "i")
	} else if strings.HasSuffix(s, "s") {
		s = remsuff(s, "s")
	}

	//1b
	if strings.HasSuffix(s, "eed") {
		if m(remsuff(s, "eed")) > 0 {
			s = repsuff(s, "eed", "ee")
		}
	} else if strings.HasSuffix(s, "ed") {
		if _v_(remsuff(s, "ed")) {
			s = cleanup1(remsuff(s, "ed"))
		}
	} else if strings.HasSuffix(s, "ing") {
		if _v_(remsuff(s, "ing")) {
			s = cleanup1(remsuff(s, "ing"))
		}
	}

	//1c
	if strings.HasSuffix(s, "y") && _v_(s) {
		s = repsuff(s, "y", "i")
	}

	//2
	s = dermorphrep(s, "ational", "ate")
	s = dermorphrep(s, "ization", "ize")
	s = dermorphrep(s, "biliti", "ble")

	//3
	s = dermorphrep(s, "icate", "ic")
	s = dermorphrem(s, "ful")
	s = dermorphrem(s, "ness")

	//4
	s = dermorphrem(s, "ance")
	s = dermorphrem(s, "ent")
	s = dermorphrem(s, "ive")

	//5a
	if strings.HasSuffix(s, "e") {
		if m(remsuff(s, "e")) > 1 {
			s = remsuff(s, "e")
		}
	} else if strings.HasSuffix(s, "ness") {
		if m(remsuff(s, "ness")) == 1 && !_o(remsuff(s, "ness")) {
			s = remsuff(s, "ness")
		}
	}

	//5b
	if m(s) > 1 && _d(s) && s[len(s)-1] == 'l' {
		s = remsuff(s, "l")
	}

	fmt.Println(s)

}

func cleanup1(s string) string {
	if strings.HasSuffix(s, "at") || strings.HasSuffix(s, "bl") {
		return s + "e"

	} else if _d(s) && s[len(s)-1] != 'l' && s[len(s)-1] != 's' && s[len(s)-1] != 'z' {
		return remsuff(s, string(s[len(s)-1]))
	} else if _o(s) && m(s) == 1 {
		return s + "e"
	}
	return s
}

func dermorphrep(s, morph, fin string) string {
	if strings.HasSuffix(s, morph) {
		if m(remsuff(s, morph)) > 0 {
			s = repsuff(s, morph, fin)
		}
	}
	return s
}

func dermorphrem(s, morph string) string {
	if strings.HasSuffix(s, morph) {
		if m(remsuff(s, morph)) > 0 {
			s = remsuff(s, morph)
		}
	}
	return s
}
