package go_koans

import "fmt"

func aboutStrings() {
	assert("a"+"bc" == "abc") // concatenation is an operation on a string type
	assert(len("abc") == 3)   // the len built-in function returns the length of v, according to its type

	assert("走" == "走") // walk past the land of 256 characters
	assert(len("走") == 3) // what is the length of a string?

	assert("abc"[0] == 97) // reflect on the ascii table, which can encode simple scripts like English
	assert("走"[1] == 181) // the expressiveness of human written language cannot be encoded in a byte

	assert("smith"[2:] == "ith")   // slicing may omit the end point
	assert("smith"[:4] == "smit")  // or the beginning
	assert("smith"[2:4] == "it")   // or neither
	assert("smith"[:] == "smith")  // or both

	var s = "smith"
	//assert.Equal(s, __string__, "The two words should be the same.")
	assert(s ==  "smith") // they can be compared directly
	assert("smith" <  "zzz")  // what is the order of a string?

	// TODO: weird to introduce this before [] syntax
	bytes := []byte{'a', 'b', 'c'}
	assert(string(bytes) == "abc") // strings can be created from byte-slices

	bytes[0] = 'z'
	assert(string(bytes) == "zbc") // byte-slices can be mutated, although strings cannot

	assert(fmt.Sprintf("hello %s", "world") == "hello world") // our old friend sprintf returns
	assert(fmt.Sprintf("hello \"%s\"", "world") == `hello "world"`)   // quoting is familiar
	assert(fmt.Sprintf("hello %q", "world") == `hello "world"`)       // although it can be done more easily

	assert(fmt.Sprintf("your balance: %d and %0.2f", 3, 4.5589) == "your balance: 3 and 4.56") // "the root of all evil" is actually a misquotation, by the way
}
