package variables

import "io"

// go does not encourage verbose code
// we gophers are lazy, I bet you are too
// thus instead of coming up with big ass names of variables
// we use short abbreviated names
// below are some examples used in go libraries

var (
	s      string    // string
	i      int       // int
	num    int       // number
	msg    string    // message
	v      string    // value
	val    string    // value
	fv     string    // flag value
	err    string    // error value
	args   []string  // agruments
	seen   bool      // has seen?
	parsed bool      // parsing ok?
	buf    []byte    // buffer
	off    int       // offeset
	op     int       // operation
	opRead int       // read operation
	l      int       // length
	n      int       // number or number of
	m      int       // another number
	c      int       // capacity
	ch     rune      // character
	a      []int     // array
	r      rune      // rune
	sep    string    // separator
	src    int       // source
	dst    int       // destination
	b      byte      // byte
	w      io.Writer // writer
	r      io.Reader // reader
	pos    int       // position
	//... the list goes on and on
	// you are not expected to remember all of them
	// but you can always refer to the documentation
	// or the source code of the library
	// this is just to give you an idea of how common abbreviations are used
	// this is just an example, you can use any name you want
)
