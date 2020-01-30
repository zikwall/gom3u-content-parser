package gom3u_content_parser

import (
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"
)

func Camelize(s string) string {
	words := strings.Split(strings.Replace(s, "-", "_", -1), "_")
	for i, word := range words {
		words[i] = ucFirst(strings.ToLower(word))
	}

	return lcFirst(strings.Join(words, ""))
}

func lcFirst(s string) string {
	if len(s) <= 1 {
		return strings.ToLower(s)
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func ucFirst(s string) string {
	if len(s) <= 1 {
		return strings.ToUpper(s)
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func ReadStringContentFromFile(source string) string {
	b, err := ioutil.ReadFile(source)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func ReadStringContentFromRemote(source string) string {
	res, err := http.Get(source)
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(contents)
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func ParseAttributes(str string) map[string]string {
	result := map[string]string{}

	lastQuote := rune(0)
	f := func(c rune) bool {
		switch {
		case c == lastQuote:
			lastQuote = rune(0)
			return false
		case lastQuote != rune(0):
			return false
		case unicode.In(c, unicode.Quotation_Mark):
			lastQuote = c
			return false
		default:
			return unicode.IsSpace(c)

		}
	}

	items := strings.FieldsFunc(str, f)
	for _, item := range items {
		x := strings.Split(item, "=")

		if _, exist := Find(availableAttributes, x[0]); !exist {
			continue
		}

		result[x[0]] = strings.Replace(x[1], `"`, "", -1)
	}

	return result
}
