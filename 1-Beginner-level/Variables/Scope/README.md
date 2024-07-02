# Scope of Variables in GoLang

In GoLang, the scope of a variable determines where the variable can be accessed within the program. GoLang supports both lexical (static) scoping and block scoping.

## Global/Package Scope
Variables declared outside of any function have global/Package scope. They can be accessed from anywhere within the package.

Example:
```go
package main

import "fmt"

var globalVar = "I'm a global variable" // At global level we have to explitily declare the variable,Shortend method will not work. ( varName := value)

func main() {
    fmt.Println(globalVar)  // Output: I'm a global variable
}
```

## Function Scope
Variables declared inside a function are local to that function and cannot be accessed from outside of it.

Example:
```go
package main

import "fmt"

func main() {
    var localVar = "I'm a local variable"
    fmt.Println(localVar)  // Output: I'm a local variable
}

```

## Block Scope

Variables declared inside a block (e.g., if statements, for loops) are local to that block.

Example:
```go
package main

import "fmt"

func main() {
    if true {
        var blockVar = "I'm a variable inside a block"
        fmt.Println(blockVar)  // Output: I'm a variable inside a block
    }

    // fmt.Println(blockVar)  // Error: undefined: blockVar
}
```