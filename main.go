package sting

import (
	"strings"
	"sort"
	"encoding/base64"
)

/**
*
* I don't know who you are. I don't know what you want.
* If you are looking for answers, I can tell you I don't have any.
* But what I do have are a very particular set of skills.
* Skills I have acquired over a very long career.
* Skills that make me a nightmare for people like you.
* If you leave my code alone, that'll be the end of it.
* I will not look for you, I will not pursue you.
* But if you don't, I will look for you, I will find you, and I will review your code.
*
* PRs are welcome :-)
*
* author: Younis Shah
* date: 22 September 2016
*
* sting - a String utility library.
*
* After all "String" is the universal data structure.
* Even aliens use it! All the world's knowledge is in the form of strings!
*/

// Appends strings to a single string
func Append(values ...string) string {
	var accumulator string
	for _, v := range values {
		accumulator += v
	}
	return accumulator
}

// Append an array to the string
func AppendArray(s string, strings []string) string {
	if len(strings) == 0 {
		return s
	} else {
		var accumulator string = s
		for _, v := range strings {
			accumulator += v
		}
		return accumulator
	}
}

// Gets the string at index value
func At(s string, index int) string {
	if index < 0 {
		panic("Cannot access an array with a negative index")
	}
	if index > len(s) {
		panic("Array index out of bounds")
	}
	return string(s[index])
}

// Gets the string as a character array
func Chars(s string) []string {
	chars := []string{}
	for _, v := range s {
		chars = append(chars, string(v))
	}
	return chars
}

// Removes all the whitespaces between string fields
func CollapseWhiteSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// Checks whether the haystack conatins all the needles
func ContainsAll(needles string, haystack []string) bool {
	_needles := strings.Fields(CollapseWhiteSpace(needles))
	if len(haystack) < len(_needles) {
		return false
	} else {
		sort.Strings(_needles)
		sort.Strings(haystack)
		for i := 0; i < len(haystack); i++ {
			if strings.Trim(haystack[i], " ") != _needles[i] {
				return false
			}
		}
		return true
	}
}

// Case-insensitive HasSuffix
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix))
}

// Encodes string with MIME base64.
// RFC 4648
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Decodes string with MIME base64.
// RFC 4648
func Base64Decode(s string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(decodedBytes)
}

// Returns the first n chars of the give string
func First(s string, n int) string {
	if s != "" && (n > 0 && n < len(s)) {
		return s[0:n]
	} else {
		return s
	}
}

// Returns the first character of a string
func Head(s string) string {
	if len(s) > 0 {
		return string(s[0])
	} else {
		return s
	}
}

// Insert a 'string' at 'index' into the given 'string'
// NOTE: indexing/counting starts at 0
func Insert(target, ins string, index int) string {
	if index <= len(target) {
		result := ""
		switch index {
		case 0 :
			result = ins + target
			break
		case len(target):
			result = target + ins
			break
		default:
			result = target[0:index + 1] + ins + target[index + 1:]
		}
		return result
	} else {
		return target
	}
}

// Returns a new string of a given length left padded with the given symbol
func LeftPad(target, symbol string, length int) string {
	if length == 0 || length < 0 {
		return target
	}
	accumulator := ""
	for i := 0; i < length - 1; i++ {
		accumulator += symbol
	}
	return accumulator + target
}

// Returns a new string of a given length right padded with the given symbol
func RightPad(target, symbol string, length int) string {
	if length == 0 || length < 0 {
		return target
	}
	accumulator := ""
	for i := 0; i < length - 1; i++ {
		accumulator += symbol
	}
	return target + accumulator
}

// Returns the string as "camelCase'd"
func CamelCase(s string) string {
	if strings.Contains(s, "-") {
		return strings.ToLower(string(s[0])) + strings.Replace(s[1:], "-", "", -1)
	} else {
		return strings.ToLower(string(s[0])) + s[1:]
	}
}

// Returns the string as "kebab-case'd"
func KebabCase(s string) string {
	fields := strings.Fields(s)
	lowerCase(&fields)
	return strings.Join(fields, "-")
}

// Returns the string as "snake_case'd"
func SnakeCase(s string) string {
	fields := strings.Fields(s)
	lowerCase(&fields)
	return strings.Join(fields, "_")
}

// Removes all invalid characters.
// ā => a, ب => b, etc
func Transliterate(s string) string {
	ascii := Ascii()
	for k, v := range ascii {
		for _, w := range v {
			s = strings.Replace(s, w, k, -1)
		}
	}
	return s
}

// Reverse a given string 'gnits' => 'sting'
func Reverse(s string) string {
	if len(s) == 0 {
		return s
	}
	bytes := []byte(s)
	length := len(s)
	halfLength := length / 2
	for i := 0; i < halfLength; i++ {
		temp := bytes[i]
		bytes[i] = bytes[length - 1 - i]
		bytes[length - 1 - i] = temp
	}
	return string(bytes)
}

func lowerCase(_fields *[]string) {
	for i := 0; i < len(*_fields); i++ {
		(*_fields)[i] = strings.ToLower((*_fields)[i])
	}
}
