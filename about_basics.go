package go_koans

func aboutBasics() {
	assert(true == true)  // what is truth?
	assert(true != false) // in it there is nothing false

	var i int = 1.0
	assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder

	k := 1 //short assignment can be used, as well
	assert(k == 1.0000000000000000000000000000000000000)

	assert(5%2 == 1)
	assert(5*2 == 10)
	assert(5/2 == 2)
	assert(5^2 == 7)

	var x int
	assert(x == 0) // listen to the darkness of an unset variable
	// listen to the darkness of an unset variable
	// what is the value of a variable declaration?
  // listen to the darkness of an unset variable
	// what is the code that is not written?

	var f float32
	assert(f == 0) // what is the code that is not written?

	var s string
	assert(s == "") // consider the emptiness of a string

	var c struct {
		x int
		f float32
		s string
	}
	assert(c.x == 0)     // what is structured emptiness?
	assert(c.f == 0) // what is the value of floating emptiness?
	assert(c.s == "")  	// listen to the structure of a string
}
