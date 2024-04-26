# Short Declarations in Go

Go provides another way of declaring variables which is using the `:=` operator. When `:=` operator is used both var keyword and type info can be omitted. Below is the format for such declaration:
```Go
{var_name} := {value}
```
Type inference will happen based upon the type of value. Also please refer to [this](./../variables/type-inference.go) article for details on type inference.

## Points about := operator

1. `:=` operator is available within a function. It cannot be used outside a function.
2. Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.
3. A variable once declared using `:=` cannot be redeclared using the `:=` operator. So the below statement will raise a compiler error  “no new variables in the left side of `:=`”.
   
   ```Go
    a := 8
    a := 16
    ```
4. `:=` operator can also be used to declare multiple variables in a single line. See below example
   ```Go 
   a,b := 1, 2
   ```
5. In the case of multiple declarations, `:=` can also be used again for a particular variable if at least one of the variables on the left-hand side is new. See the below example. Notice that `b` is again declared using `:=` This is only possible if at least one of the variables is new which is variable `c` here. In this case it acts as a assignment for variable `b`
   ```Go
   package main

   import "fmt"

   func main() {
       a, b := 1, 2
       b, c := 3, 4 // this works
       a, b := 5, 6 // this doesn't work
       fmt.Println(a, b, c)
   }
   ```

## Examples
In below code;
- Type of `m2` is correctly inferred as `int` as value assigned to it is `123` which is `int`.
- Similarly, type of `n2` is also correctly inferred as `string` as the value assigned to it is `“circle”` which is a `string` and so on.
- Also notice that the type of variable `t2` is inferred correctly as a struct `main.sample`
- Type of `u2` is also correctly inferred as `[]string` as this is type returned by the `get()` function call.
  
```Go
package main

import "fmt"

func main() {
    m2 := 123                   //Type Inferred will be int
    n2 := "circle"              //Type Inferred will be string
    o2 := 5.6                   //Type Inferred will be float64
    p2 := true                  //Type Inferred will be bool
    q2 := 'a'                   //Type Inferred will be rune
    r2 := 3 + 5i                //Type Inferred will be complex128
    s2 := &sample{name: "test"} //Type Inferred will be pointer to main.sample
    t2 := sample{name: "test"}  //Type Inferred will be main.sample
    u2 := get()                 //Type inferred will []string as that is the return value of function get()

    fmt.Println("")
    fmt.Printf("m2: Type: %T Value: %v\n", m2, m2)
    fmt.Printf("n2: Type: %T Value: %v\n", n2, n2)
    fmt.Printf("o2: Type: %T Value: %v\n", o2, o2)
    fmt.Printf("p2: Type: %T Value: %v\n", p2, p2)
    fmt.Printf("q2: Type: %T Value: %v\n", q2, q2)
    fmt.Printf("r2: Type: %T Value: %v\n", r2, r2)
    fmt.Printf("s2: Type: %T Value: %v\n", s2, s2)
    fmt.Printf("t2: Type: %T Value: %v\n", t2, t2)
    fmt.Printf("u2: Type: %T Value: %v\n", u2, u2)
}

type sample struct {
    name string
}

func get() []string {
    return []string{"a", "b"}
}
```

### Output:
```Go
m2: Type: int Value: 123
n2: Type: string Value: circle
o2: Type: float64 Value: 5.6
p2: Type: bool Value: true
q2: Type: int32 Value: 97
r2: Type: complex128 Value: (3+5i)
s2: Type: *main.sample Value: &{test}
t2: Type: main.sample Value: {test}
u2: Type: []string Value: [a b]
```