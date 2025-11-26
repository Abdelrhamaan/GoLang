package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	msg1 := "Hello Go !"
	msg2 := "Hello\nGo !"
	msg3 := `Hello\nGo !` // bact-tick escape any seq-charachter

	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)

	// every char in string is a rune
	// rune is an int value that represents a char

	fmt.Println("Length of msg is: ", len(msg1))
	fmt.Println("Length of msg is: ", len(msg2))
	fmt.Println("Second char is: ", msg1[2])    // returns ASCII value of char
	fmt.Printf("Second char is: %c\n", msg1[2]) // prints the character

	greeting := "Hi"
	name := " Noor"
	fmt.Println(greeting + name)

	// comparing strings are by lexicographic (dictionary‑order)
	// he compares the ascii value of char

	// the Go compiler does lexicographic (dictionary‑order) comparison of the two strings.
	// The algorithm works like this:

	// 1-Treat each string as a slice of runes (Unicode code points).
	// 2-A rune is an int32 that represents a single Unicode character.
	// 3-Compare the runes one by one, from the first position onward.
	// 4-As soon as a pair of runes differs, the result of the comparison is the result of comparing those two rune values.
	// 5-If all runes are equal up to the length of the shorter string, the shorter string is considered less.

	// ex...
	// "cat" < "cater"   // true, because "cat" is a prefix of "cater"
	// "cater" < "cat"   // false
	// "cat" < "car"     // false, because at index 2 't' (0x74) > 'r' (0x72)
	value1 := "App"
	value2 := "ap"
	value3 := "Banana"
	value4 := "banana"
	value5 := "banan"

	fmt.Println(value1 < value2)
	fmt.Println(value3 < value4)
	fmt.Println(value5 < value4)

	// iteration
	// %d: for printing the integer value of a rune or number
	// %c: for printing the actual character represented by a rune
	// %x: for printing the hexadecimal representation of a rune or number
	// %v: for printing the default representation of a value (for a rune, this is its integer value)

	for i, v := range value3 {
		fmt.Println("Index is: ", i, "and value is: ", v)
		fmt.Printf("Index is %d and value is %c\n", i, v)
	}

	// cal number of alphapets
	// The `len()` function returns the number of bytes in a string.
	// The `utf8.RuneCountInString()` function returns the number of runes (Unicode code points) in a string.
	// For strings containing multi-byte characters (e.g., non-ASCII), these two values will differ.
	// For example, `len("Hello 👋")` would be 10 (6 bytes for "Hello " and 4 bytes for "👋"), while `utf8.RuneCountInString("Hello 👋")` would be 7.
	fmt.Println(utf8.RuneCountInString(name))
	fmt.Println(utf8.RuneCountInString("Hello 👋"))
	fmt.Println(len("Hello 👋"))

	s := "Hello 👋"
	b := []byte(s) // convert to a byte slice

	for i, v := range b {
		fmt.Printf("byte %d: 0x%02x (%d)\n", i, v, v)
	}

	for i := 0; i < len(s); i++ {
		// s[i] is a single byte (type uint8)
		fmt.Printf("byte %d: 0x%02x (%d)\n", i, s[i], s[i])
	}

	// rune can use single quote
	var r rune = 'a'
	var arabic rune = 'ب'
	fmt.Println("r", r)
	fmt.Println("arabic", arabic)
	fmt.Printf("r %c\n", r)
	fmt.Printf("arabic %c\n", arabic)
	fmt.Printf("r %d\n", r)
	fmt.Printf("arabic %d\n", arabic)
	fmt.Printf("r %x\n", r)
	fmt.Printf("arabic %x\n", arabic)
	fmt.Printf("r %v\n", r)
	fmt.Printf("arabic %v\n", arabic)

	fmt.Println("================================")
	converetRune := string(r)
	convetArabic := string(arabic)
	fmt.Println(converetRune)
	fmt.Println(convetArabic)
	fmt.Printf("type of r is  %T\n", converetRune)
	// rune type is int32
	fmt.Printf("type of ris  %T\n", r)
	fmt.Println("================================")
	// printing smily emoji
	emoji := "👋"
	arabicMsg := "سلام عليكم"
	fmt.Println(emoji)
	fmt.Printf("emoji: %s\n", emoji)
	fmt.Printf("emoji is %c\n", emoji)
	fmt.Println(len(emoji))
	fmt.Println(utf8.RuneCountInString(emoji))
	fmt.Println("================================")
	fmt.Println(arabicMsg)
	fmt.Printf("arabicMsg is  %s\n", arabicMsg)
	fmt.Println("================================")
}
