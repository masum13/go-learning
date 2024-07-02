# GoLang Print Format Specifiers

GoLang offers a variety of format specifiers for use in print statements. These specifiers allow you to format your output in different ways. Here are some commonly used format specifiers:

| Specifier | Description                                           | Example Output              |
|-----------|-------------------------------------------------------|-----------------------------|
| `%v`      | Default format                                        | `fmt.Printf("%v", 123)`     -> `123` |
| `%+v`     | Default format with field names (for structs)         | `fmt.Printf("%+v", person)` -> `{Name: John Age: 30}` |
| `%#v`     | Go-syntax representation of the value                 | `fmt.Printf("%#v", 123)`    -> `123` |
| `%T`      | Type of the value                                     | `fmt.Printf("%T", 123)`     -> `int` |
| `%%`      | Literal percent sign                                  | `fmt.Printf("%%")`          -> `%` |
| `%t`      | Boolean                                               | `fmt.Printf("%t", true)`    -> `true` |
| `%b`      | Base 2 (binary)                                       | `fmt.Printf("%b", 5)`       -> `101` |
| `%c`      | Character                                             | `fmt.Printf("%c", 97)`      -> `a` |
| `%d`      | Base 10 (decimal)                                     | `fmt.Printf("%d", 123)`     -> `123` |
| `%o`      | Base 8 (octal)                                        | `fmt.Printf("%o", 8)`       -> `10` |
| `%O`      | Base 8 (octal) with leading 0o                       | `fmt.Printf("%#o", 8)`      -> `0o10` |
| `%x`      | Base 16 (hexadecimal, lowercase)                      | `fmt.Printf("%x", 255)`     -> `ff` |
| `%X`      | Base 16 (hexadecimal, uppercase)                      | `fmt.Printf("%X", 255)`     -> `FF` |
| `%U`      | Unicode format: U+1234                                | `fmt.Printf("%U", 'A')`     -> `U+0041` |
| `%e`      | Scientific notation (lowercase)                       | `fmt.Printf("%e", 1234.56)` -> `1.234560e+03` |
| `%E`      | Scientific notation (uppercase)                       | `fmt.Printf("%E", 1234.56)` -> `1.234560E+03` |
| `%f`      | Decimal point but no exponent                         | `fmt.Printf("%f", 123.456)` -> `123.456000` |
| `%F`      | Same as %f                                            | `fmt.Printf("%F", 123.456)` -> `123.456000` |
| `%g`      | Uses %e for large exponents, %f otherwise             | `fmt.Printf("%g", 12345678)`-> `1.2345678e+07` |
| `%G`      | Uses %E for large exponents, %F otherwise             | `fmt.Printf("%G", 12345678)`-> `1.2345678E+07` |
| `%s`      | String                                                | `fmt.Printf("%s", "hello")` -> `hello` |
| `%q`      | Double-quoted string                                  | `fmt.Printf("%q", "hello")` -> `"hello"` |
| `%p`      | Pointer (hexadecimal)                                 | `fmt.Printf("%p", &a)`      -> `0x10a47fc0` |

These specifiers can be used with `fmt.Printf`, `fmt.Sprintf`, and similar functions to format output in Go.

# Difference Between `Printf` and `Sprintf` in Go

Go provides several functions for formatted I/O, with `Printf` and `Sprintf` being two commonly used functions. Here is a comparison of their key differences:

- **`Printf`**: Formats according to a format specifier and writes to the standard output (typically the console). It does not return the formatted string.

  ```go
  fmt.Printf("Hello, %s!\n", "World")  // Output: Hello, World!
  ```

- **`Sprintf`**: Formats according to a format specifier and returns the resulting string. It does not write to the standard output. This is useful when you need to store the formatted string in a variable or use it for further processing.
 ```go
 formattedString := fmt.Sprintf("Hello, %s!", "World")
fmt.Println(formattedString)  // Output: Hello, World!
  ```
### Summary
- Use `Printf` when you want to print directly to the console or standard output.
- Use `Sprintf` when you need the formatted string for further use in your program.

