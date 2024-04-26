# Enumerations in Go

Imagine you’re building a Todo application using Go. Each task in your application has a status associated with it, indicating whether it’s **“completed”**, **“archived”** or **“deleted”**. 

As a developer, you want to ensure that only valid status values are accepted when creating or updating tasks via your API endpoints. Additionally, you want to provide clear error messages to users if they attempt to set an invalid status apart from the valid ones. 
This is where enumerations come in handy.

> An enumeration is a set of named integer constants that represent a set of related named constants.

In the context of the Todo application, you can define an enumeration called `TaskStatus` with three named constants: `Completed`, `Archived`, and `Deleted`.

## Why Enums?
- Enums make code more readable by providing descriptive names for values, rather than using raw integer or string constants. This enhances code clarity and makes it easier for other developers to understand and maintain the codebase.
  
- Enums provide compile-time checking, which helps catch errors early in the development process. This prevents invalid or unexpected values from being assigned to variables, reducing the likelihood of runtime errors.
  
- By restricting the possible values a variable can hold to a predefined set, enums help prevent logic errors caused by incorrect or unexpected values.

### But the bad news is that Go <span style="color: red">does not</span> support enumerations out of the box.

There are however hacks or workarounds to use enums in Go.
We use constants and custom types to achieve the same.

## Defining Enums with Constants
Constants are a straightforward way to define enums in Go. Let’s consider a basic example where we define enum constants representing different days of the week.

```go
package main

import "fmt"

const (
    SUNDAY    = "Sunday"
    MONDAY    = "Monday"
    TUESDAY   = "Tuesday"
    WEDNESDAY = "Wednesday"
    THURSDAY  = "Thursday"
    FRIDAY    = "Friday"
    SATURDAY  = "Saturday"
)

func main() {
    day := MONDAY
    fmt.Println("Today is ", day) // "Today is Monday"
}
```

## Simulating Enums with Custom Types and iota
Another common approach to defining enums in Go is to use custom types and the `iota` keyword. The `iota` identifier is used to create a sequence of related values without explicitly assigning them. It resets to zero whenever the `const` block is encountered.

### What is iota ?
iota is a special identifier in Go that is used with const declarations to generate a sequence of related values automatically. 

It simplifies the process of defining sequential values, particularly when defining enums or declaring sets of constants with incrementing values. Know more about iota [here](./iota.go).

When iota is used in a const declaration, it starts with the value 0 and increments by 1 for each subsequent occurrence within the same const block. If iota appears in multiple const declarations within the same block, its value is reset to 0 for each new const block. 

```Go
package main

import "fmt"

const (
    SUNDAY = iota // SUNDAY is assigned 0
    MONDAY        // MONDAY is assigned 1 (incremented from SUNDAY)
    TUESDAY       // TUESDAY is assigned 2 (incremented from MONDAY)
    WEDNESDAY     // WEDNESDAY is assigned 3 (incremented from TUESDAY) 
    THURSDAY      // THURSDAY is assigned 4 (incremented from WEDNESDAY)
    FRIDAY        // FRIDAY is assigned 5 (incremented from THURSDAY)
    SATURDAY      // SATURDAY is assigned 6 (incremented from FRIDAY)
)

func main() {
    fmt.Println(MONDAY, WEDNESDAY, FRIDAY) // Output: 1 3 5
}
```

We can add custom types to the mix to create enums in Go. Here’s how you can define an enum for the days of the week using a custom type and iota:

```Go
package main

import "fmt"

type Day int

const (
    Sunday Day = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func main() {
    day := Monday
    fmt.Println("Today is ", day) // Output: Today is 1
}
```

In this example, we define a custom-type Day and use iota to auto-increment the values of the constants. This approach is more concise and offers better type safety.

Optionally, we can also use like iota + 1 , if we want our constants to start from 1.

Another example:
```Go
package main

import "fmt"

type Size uint8

const (
	small      Size = 0
	medium     Size = 1
	large      Size = 2
	extraLarge Size = 3
)

func main() {
	fmt.Println(small)
	fmt.Println(medium)
	fmt.Println(large)
	fmt.Println(extraLarge)
}
```

We can also define a method toString on Size type to print the exact value of enum . See below program for the same.

```Go
package main
import "fmt"
type Size int
const (
    small = Size(iota)
    medium
    large
    extraLarge
)
func main() {
    var m Size = 1
    m.toString()
}
func (s Size) toString() {
    switch s {
    case small:
        fmt.Println("Small")
    case medium:
        fmt.Println("Medium")
    case large:
        fmt.Println("Large")
    case extraLarge:
        fmt.Println("Extra Large")
    default:
        fmt.Println("Invalid Size entry")
    }
}
```