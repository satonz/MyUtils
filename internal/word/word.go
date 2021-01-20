package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func UnderScoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)
	return s
}

func UnderscoreToLowerCamelCase(s string) string {
	s = UnderScoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func CamelcaseToUnderscore(s string) string {
	var outPut []rune
	for i, r := range s {
		if i == 0 {
			outPut = append(outPut, unicode.ToLower(r))
			continue
		}

		if unicode.IsUpper(r) {
			outPut = append(outPut, '_')
		}

		outPut = append(outPut, unicode.ToLower(r))
	}

	return string(outPut)
}
