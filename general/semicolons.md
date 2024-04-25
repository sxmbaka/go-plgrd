# Semicolons in Go

In Go, semicolons are used to terminate statements, but unlike in C, they are not explicitly written in the source code. Instead, the lexer automatically inserts semicolons as it scans the code, resulting in source code that is mostly free of semicolons.

The Go Programming Language Specification states:<br>
***The formal grammar uses semicolons ";" as terminators in a number of productions.***

### Go programs may omit most of these semicolons using the following two rules:
1. When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a line's final token if that token is an `identifier`, a `literal`, a `closing parenthesis` or `bracket`, or one of the following tokens:

	```Go
	break continue fallthrough return ++ -- ) }
	```

2. A semicolon can also be omitted immediately before a closing brace ( `)` or `}` ), so a statement such as
   
   ```Go
   go func() { for { dst <- <-src } } ()
   ```

   needs no semicolons. Idiomatic Go programs have semicolons only in places such as for loop clauses, to separate the initializer, condition, and continuation elements. They are also necessary to separate multiple statements on a line, should you write code that way.

3. One consequence of the semicolon insertion rules is that you cannot put the opening brace of a control structure (`if`, `for`, `switch`, or `select`) on the next line. If you do, a semicolon will be inserted before the brace, which could cause unwanted effects. Write them like this
	```Go
	if i < f() {
    	g()
	}
	```
	not like this
	```Go
	if i < f()  // wrong!
	{           // wrong!
    	g()
	}
	```
