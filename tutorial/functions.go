package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main1() {

	// =========================
	// STRING TRIMMING
	// =========================

	str := "   Hello World   "

	fmt.Println(strings.TrimSpace(str))
	fmt.Println(strings.Trim("///hello///", "/"))
	fmt.Println(strings.TrimLeft("   hello", " "))
	fmt.Println(strings.TrimRight("hello   ", " "))
	fmt.Println(strings.TrimPrefix("Mr John", "Mr "))
	fmt.Println(strings.TrimSuffix("file.txt", ".txt"))

	// =========================
	// STRING CASE CONVERSION
	// =========================

	fmt.Println(strings.ToLower("HELLO"))
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.EqualFold("Go", "go"))

	// =========================
	// CHARACTER CASE CONVERSION
	// =========================

	fmt.Printf("%c\n", unicode.ToLower('A'))
	fmt.Printf("%c\n", unicode.ToUpper('b'))

	// =========================
	// STRING INDEXING
	// =========================

	s := "A man, a plan, a canal: Panama"

	fmt.Println(len(s))
	fmt.Println(string(s[0]))
	fmt.Println(string(s[2]))

	// =========================
	// BYTE / RUNE CONVERSION
	// =========================

	fmt.Printf("%T\n", s[0]) // uint8 / byte

	bytes := []byte("hello")
	fmt.Println(bytes)

	runes := []rune("नमस्ते")
	fmt.Println(runes)

	// =========================
	// REVERSE STRING
	// =========================

	fmt.Println(reverseString("hello"))
	fmt.Println(reverseString("नमस्ते"))

	// =========================
	// MAPS
	// =========================

	mp := make(map[rune]int)

	for _, ch := range "hello" {
		mp[ch]++
	}

	fmt.Println(mp)

	// map lookup
	count, ok := mp['l']

	fmt.Println(count)
	fmt.Println(ok)

	// delete key
	delete(mp, 'l')

	fmt.Println(mp)

	// clear map
	clear(mp)

	fmt.Println(mp)

	// =========================
	// TYPE CONVERSION
	// =========================

	var b byte = 'A'

	fmt.Println(int(b))
	fmt.Println(string(b))

	var r rune = 'B'

	fmt.Println(string(r))

	// =========================
	// LOOPS
	// =========================

	word := "golang"

	for i := 0; i < len(word); i++ {
		fmt.Println(i, string(word[i]))
	}

	for i, ch := range word {
		fmt.Println(i, string(ch))
	}

	// =========================
	// SWAP
	// =========================

	a, c := 10, 20

	a, c = c, a

	fmt.Println(a, c)
}
