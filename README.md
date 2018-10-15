# baseN
baseN lets you encode any number into any base of utf8 characters.

`go get github.com/jakobii/baseN`

## Examples
Base4 in emojis.
```go
package main
import (
    "github.com/jakobii/baseN"
    "fmt"
)
func main (){
    b4 := []rune("😁😡😪😋")
    encoding := baseN.Encode(255, b4)
    fmt.Println(encoding)
}
// prints: 😋😋😋😋
```

Base62 __[0-9a-zA-Z]__
```go
package main
import (
    "github.com/jakobii/baseN"
    "fmt"
)
func main (){
    b62 := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    encoding := Encode(65434198762432165, b62)
    fmt.Println(encoding)
}
// prints: 4PGJLkZKpD
```