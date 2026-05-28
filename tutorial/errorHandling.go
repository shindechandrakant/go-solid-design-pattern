package main

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}

	//mp := make(map[string]weak.Pointer[string])
	return x
}

// 65 'A'
// 90
// 122
// 97 -> 'a'

// 65 'A'
// 90
// 122
// 97 -> 'a'
func isChar(char uint8) bool {
	return (char >= 65 && char <= 90) ||
		(char >= 97 && char <= 122) ||
		(char <= 49 && char >= 40)
}

func lower(char uint8) uint8 {
	if char <= 49 && char >= 40 {
		return char
	}
	if char >= 97 && char <= 122 {
		return char
	}
	return 32 + char
}

func isPalindrome(s string) bool {

	start, end := 0, len(s)-1
	for start <= end {

		if !isChar(s[start]) {
			start++
			continue
		} else if !isChar(s[end]) {
			end--
			continue
		} else if lower(s[start]) != lower(s[end]) {
			// fmt.Println(start, end, s[start], s[end], lower(s[start]), lower(s[end]))
			return false
		} else {
			start++
			end--
		}
	}
	return true
}
func main() {
	//r := Must(regexp.Compile("()"))

	isPalindrome("1A21")
	//fmt.Println(r)
}
